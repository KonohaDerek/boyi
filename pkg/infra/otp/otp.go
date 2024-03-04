package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"math"
	"net/url"
	"strings"
	"time"
)

// Google One-Time Password
type GoogleOTP struct {
	Secret     string
	OtpAuthURL string
	Digits     int
	Algorithm  Algorithm
	QrCodeURL  string
}

type Algorithm int

const (
	// SHA1 is the default hash algorithm used by Google Authenticator
	SHA1 Algorithm = iota
	// SHA256 is the hash algorithm used by Google Authenticator
	SHA256
	// SHA512 is the hash algorithm used by Google Authenticator
	SHA512
)

// New One-time Password For Google Authenticator
// key: the seed of base32 encoded secret key
// issuer: the issuer of the account
// accountName: the name of the account
// digits: number of digits , default is 6
func NewGoogleOTP(key, issuer, accountName string, digits int) (GoogleOTP, error) {
	var (
		result GoogleOTP
	)

	// key is required
	if key == "" {
		return result, errors.New("key is required")
	}

	// issuer is required
	if issuer == "" {
		return result, errors.New("issuer is required")
	}

	// accountName is required
	if accountName == "" {
		return result, errors.New("accountName is required")
	}

	// set default digits
	if digits == 0 {
		digits = 6
	}

	secret := base32.StdEncoding.EncodeToString([]byte(key))

	result = GoogleOTP{
		Secret:     strings.ToUpper(secret),
		OtpAuthURL: fmt.Sprintf("otpauth://totp/%s:%s?digits=%d&issuer=%s&secret=%s", issuer, accountName, digits, issuer, strings.ToUpper(secret)),
		Digits:     digits,
		Algorithm:  SHA1,
	}
	// generate qrcode url for google authenticator
	result.QrCodeURL = fmt.Sprintf("https://chart.googleapis.com/chart?chs=200x200&chld=M%%7C0&cht=qr&chl=%s", url.QueryEscape(result.OtpAuthURL))
	return result, nil
}

// OTP Code Valid
// secret: base32 encoded secret key
// code: OTP code
// period: time period
// digits: number of digits
// algorithm: hash algorithm
func IsValid(secret, code string, period uint, digits int, algorithm Algorithm) (bool, error) {
	expectedCode, err := NewTOTP(strings.ToUpper(secret), period, digits, algorithm)
	if err != nil {
		return false, err
	}
	return code == expectedCode, nil
}

// TOTP: Time-Based One-Time Password Algorithm
// https://datatracker.ietf.org/doc/html/rfc6238
// TOTP(K,T) = HOTP(K,C)
func NewTOTP(secret string, period uint, digits int, algorithm Algorithm) (string, error) {
	if digits == 0 {
		digits = 6
	}

	// Time step default is 30 seconds
	if period == 0 {
		period = 30
	}

	// base32 decode secret
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// TC =  最低值((unixtime(當前時間)−unixtime(T0))/TS)，其中TS為時間間隔，即時間間隔為30秒。
	tc := time.Now().Unix() / int64(period)
	return NewHOTP(key, tc, digits, algorithm), nil
}

// HOTP: An HMAC-Based One-Time Password Algorithm
// https://datatracker.ietf.org/doc/html/rfc4226
// HOTP(K,C) = Truncate(HMAC-SHA-1(K,C))
func NewHOTP(key []byte, counter int64, digits int, algorithm Algorithm) string {
	var (
		h func() hash.Hash
	)

	// set hash algorithm
	switch algorithm {
	case SHA1:
		h = sha1.New
	case SHA256:
		h = sha256.New
	case SHA512:
		h = sha512.New
	default:
		h = sha1.New
	}
	mac := hmac.New(h, key)
	binary.Write(mac, binary.BigEndian, counter)
	sum := mac.Sum(nil)

	// "Dynamic truncation" in RFC 4226
	// offset is the offset of truncated hash value
	offset := sum[len(sum)-1] & 0xf

	// calculate truncated hash value
	truncatedHash := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	// truncated hash value is a digits number
	code := truncatedHash % int64(math.Pow10(digits))
	otp := fmt.Sprintf("%0*d", digits, code)
	return otp
}
