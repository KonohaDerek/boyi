package user

import (
	"context"
	"fmt"
	"testing"
)

func Test_service_CreateTouristUser(t *testing.T) {
	type args struct {
		ctx       context.Context
		deviceUID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx:       context.Background(),
				deviceUID: "create tourist user test device",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := suite.svc
			got, err := s.CreateTouristUser(tt.args.ctx, tt.args.deviceUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateTouristUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%+v\n", got)
		})
	}
}
