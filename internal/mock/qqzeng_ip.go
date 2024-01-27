package mock

import (
	"testing"

	"boyi/pkg/Infra/qqzeng_ip"

	gomock "github.com/golang/mock/gomock"
)

// NewQQZengIP ...
func NewQQZengIP(t *testing.T) qqzeng_ip.IPSearch {
	m := gomock.NewController(t)
	mock := NewMockIPSearch(m)

	mock.EXPECT().Get("1.1.1.1").AnyTimes().Return("台灣|台北|||||")
	mock.EXPECT().Get(gomock.Any()).AnyTimes().Return("||||||")

	return mock
}
