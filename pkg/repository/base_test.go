package repository

import (
	"boyi/pkg/iface"
	"boyi/pkg/model/dto"
	"boyi/pkg/model/option"
	"context"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func Test_repository_List(t *testing.T) {
	type args struct {
		ctx  context.Context
		tx   *gorm.DB
		data interface{}
		opt  iface.WhereOption
	}

	var data []dto.User
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx:  suite.ctx,
				tx:   nil,
				data: &data,
			},
		},
	}

	repo := suite.repo
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.List(tt.args.ctx, tt.args.tx, &data, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("%+v %+v\n", data, got)
		})
	}
}

func Test_repository_Create(t *testing.T) {
	type args struct {
		ctx  context.Context
		tx   *gorm.DB
		data iface.Model
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test create role",
			args: args{
				ctx: suite.ctx,
				tx:  nil,
				data: &dto.Role{
					Name: "Test role",
					Authority: dto.Authority{
						dto.Chatroom_System:       struct{}{},
						dto.MyChatroom_Management: struct{}{},
					},
				},
			},
		},
	}

	repo := suite.repo
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Create(tt.args.ctx, tt.args.tx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_Get(t *testing.T) {
	type args struct {
		ctx   context.Context
		tx    *gorm.DB
		model iface.Model
		opt   iface.WhereOption
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test: role get",
			args: args{
				ctx:   suite.ctx,
				tx:    nil,
				model: &dto.Role{},
				opt:   &option.RoleWhereOption{},
			},
		},
	}

	repo := suite.repo
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Get(tt.args.ctx, tt.args.tx, tt.args.model, tt.args.opt); (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("%+v\n", tt.args.model)
		})
	}
}

func Test_repository_BatchInsert(t *testing.T) {
	type args struct {
		ctx  context.Context
		tx   *gorm.DB
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				ctx: suite.ctx,
				tx:  suite.repo.GetDB().Begin(),
				data: []dto.User{
					{
						Username: "test1",
					},
					{
						Username: "test2",
					},
				},
			},
		},
	}
	repo := suite.repo
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.BatchInsert(tt.args.ctx, tt.args.tx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("repository.BatchInsert() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.args.tx.Commit()
		})
	}
}
