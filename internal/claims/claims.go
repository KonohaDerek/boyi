package claims

import (
	"boyi/pkg/infra/errors"
	"boyi/pkg/model/enums/types"
	"context"

	"gopkg.in/vmihailenco/msgpack.v2"
)

// Claims
type Claims struct {
	Id          uint64            `json:"id,omitempty"`
	AccountType uint64            `json:"account_type,omitempty"`
	Competences map[string]bool   `json:"competences,omitempty"`
	Username    string            `json:"username,omitempty"`
	AliasName   string            `json:"alias_name,omitempty"`
	Token       string            `json:"token,omitempty"`
	DeviceUid   string            `json:"device_uid,omitempty"`
	ExpiredAt   int64             `json:"expired_at,omitempty"`
	Extra       map[string]string `json:"extra,omitempty"`
}

// claimsKey ...
type claimsKey struct{}

// GetClaims ...
func GetClaims(ctx context.Context) (Claims, error) {
	c := ctx.Value(claimsKey{})
	if c == nil {
		return Claims{}, errors.Wrap(errors.ErrInvalidAuthenticationInfo, "claims is empty")
	}

	claims, ok := c.(Claims)
	if !ok {
		return Claims{}, errors.Wrap(errors.ErrInvalidAuthenticationInfo, "claims struct is invalid")
	}
	return claims, nil
}

// SetClaims ...
func SetClaimsToContext(ctx context.Context, c Claims) context.Context {
	return context.WithValue(ctx, claimsKey{}, c)
}

func (c Claims) VerifyRole(keys ...string) error {
	if c.AccountType == uint64(types.AccountType__System) {
		return nil
	}

	for i := range keys {
		if _, ok := c.Competences[keys[i]]; ok {
			return nil
		}
	}

	return errors.NewWithMessagef(errors.ErrNotAllowed, "role not allowed, role: %+v", keys)
}

func (c Claims) Marshal() []byte {
	b, _ := msgpack.Marshal(c)
	return b
}

func (c *Claims) Unmarshal(s string) error {
	if err := msgpack.Unmarshal([]byte(s), &c); err != nil {
		return err
	}
	return nil
}
