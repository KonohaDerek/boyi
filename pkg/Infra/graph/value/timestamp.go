package value

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"boyi/pkg/Infra/errors"

	"github.com/99designs/gqlgen/graphql"
)

// if the type referenced in .gqlgen.yml is a function that returns a marshaller we can use it to encode and decode
// onto any existing go type.
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		t = time.Unix(0, 0)
	}
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
	})
}

// Unmarshal{Typename} is only required if the scalar appears as an input. The raw values have already been decoded
// from json into int/float64/bool/nil/map[string]interface/[]interface
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	switch t := v.(type) {
	case int:
		return time.Unix(int64(t), 0), nil
	case int64:
		return time.Unix(t, 0), nil
	case json.Number:
		unixT, err := t.Int64()
		if err != nil {
			return time.Time{}, errors.NewWithMessage(errors.ErrInvalidInput, "time should be a unix timestamp")
		}
		return time.Unix(unixT, 0), nil
	default:
		return time.Time{}, errors.NewWithMessage(errors.ErrInvalidInput, "time should be a unix timestamp")
	}
}

func GetSmallerTime(x, y time.Time) time.Time {
	if x.IsZero() {
		return y
	}
	if y.IsZero() {
		return x
	}

	if x.Before(y) {
		return x
	}
	return y
}

func GetSmallerUnixTime(x, y int) int {
	a, b := time.Unix(int64(x), 0), time.Unix(int64(y), 0)

	if a.IsZero() {
		return int(b.Unix())
	}
	if b.IsZero() {
		return int(a.Unix())
	}

	if a.Before(b) {
		return x
	}
	return y
}
