package dto

import (
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRole_ConvertToSingleLayer(t *testing.T) {
	type fields struct {
		ID                 uint64
		Name               string
		IsEnable           common.YesNo
		Authority          Authority
		SupportAccountType types.AccountType
		CreatedAt          time.Time
		CreateUserID       uint64
		UpdatedAt          time.Time
		UpdateUserID       uint64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "normal test",
			fields: fields{
				Authority: Authority{
					Chatroom_System:       {},
					MyChatroom_Management: {},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Role{
				ID:                 tt.fields.ID,
				Name:               tt.fields.Name,
				IsEnable:           tt.fields.IsEnable,
				Authority:          tt.fields.Authority,
				SupportAccountType: tt.fields.SupportAccountType,
				CreatedAt:          tt.fields.CreatedAt,
				CreateUserID:       tt.fields.CreateUserID,
				UpdatedAt:          tt.fields.UpdatedAt,
				UpdateUserID:       tt.fields.UpdateUserID,
			}
			assert.Equal(t, 2, len(r.Authority), "authority is not equal")
		})
	}
}
