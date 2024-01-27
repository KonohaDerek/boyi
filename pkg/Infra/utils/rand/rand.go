package rand

import (
	"math/rand"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	rand_gen    = rand.New(rand.NewSource(time.Now().UnixNano()))
	rand_locker sync.Mutex
)

/**
 * 獲取n亂數字元 （ABCDEFGHJKLMNPQRSTUVWXYZ23456789）
 */
func GetRandomString(n int) string {
	letterRunes := "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterRunes[rand.Int63()%int64(len(letterRunes))]
	}
	return string(b)
}

/**
 * 生成n个数字字母 (abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789)
 */
func GenHashWithLength(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return randStringBytesMaskImprSrc(n, letterBytes)
}

/*
 * n 个数
 * t 1=>字母 ，2 数字，3数字+字母
 */
func GenHashWithLengthType(n, t int) string {
	var letter = ""

	if t&1 == 1 {
		letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if t&2 == 2 {
		letter += "0123456789"
	}
	return randStringBytesMaskImprSrc(n, letter)
}

func randStringBytesMaskImprSrc(n int, letterBytes string) string {
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// 獲取n內隨機數
func RandIntn(n int) int {
	rand_locker.Lock()
	nRet := rand_gen.Intn(n)
	rand_locker.Unlock()
	return nRet
}

// 獲取UUID
func RandomUUID() string {
	u1 := uuid.Must(uuid.NewV4(), nil)
	return string(u1.String())
}
