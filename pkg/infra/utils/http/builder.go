package http

import (
	"fmt"
	"net/url"
)

// BuildURLWithQueryString 組合 query string 回傳完整 url 字串
func BuildURLWithQueryString(oriURL string, query map[string]interface{}) (retURL string) {
	u, _ := url.Parse(oriURL)
	q := u.Query()
	for k, v := range query {
		q.Set(k, fmt.Sprintf("%v", v))
	}
	u.RawQuery = q.Encode()
	return u.String()
}
