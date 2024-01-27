package dto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthority_FromMenuKeys(t *testing.T) {
	_menuMap = map[ManagerMenuKey]Menu{
		Manager_System: {
			Key: Manager_System,
		},
		Manager_Management: {
			SuperKey: Manager_System,
			Key:      Manager_Management,
		},
	}
	tests := []struct {
		name string
		m    Authority
	}{
		{
			name: "Normal test",
			m: Authority{
				Manager_System:     struct{}{},
				Manager_Management: struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.GetMenus()
			fmt.Printf("%+v\n", got)
			assert.Equal(t, 1, len(got), "len need 1")
			assert.NotEqual(t, 0, len(got[0].Next))
		})
	}
}
