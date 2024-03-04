package value

import (
	"encoding/json"
	"io"
	"strconv"

	"boyi/pkg/infra/errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

func MarshalUint64(i uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatUint(i, 10))
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case int64:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case string:
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, errors.NewWithMessagef(errors.ErrInvalidInput, "%s is not number", v)
		}
		return n, nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case json.Number:
		n, err := strconv.ParseUint(v.String(), 10, 64)
		if err != nil {
			return 0, errors.NewWithMessagef(errors.ErrInvalidInput, "%s is not number", v)
		}
		return n, nil
	default:
		return 0, errors.NewWithMessagef(errors.ErrInvalidInput, "%T is not an uint", v)
	}
}

func MarshalDecimal(i decimal.Decimal) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, i.String())
	})
}

func UnmarshalDecimal(v interface{}) (decimal.Decimal, error) {
	switch v := v.(type) {
	case int64:
		return decimal.NewFromInt(v), nil
	case int:
		return decimal.NewFromInt(int64(v)), nil
	case string:
		return decimal.NewFromString(v)
	case uint:
		return decimal.NewFromInt(int64(v)), nil
	case uint64:
		return decimal.NewFromInt(int64(v)), nil
	case json.Number:
		return decimal.NewFromString(v.String())
	case float32:
		return decimal.NewFromFloat(float64(v)), nil
	case float64:
		return decimal.NewFromFloat(v), nil
	default:
		return decimal.Zero, errors.NewWithMessagef(errors.ErrInvalidInput, "%T is not an decimal", v)
	}
}
