package repository

import (
	"context"
	"testing"

	"gorm.io/gorm"

	"gin-chat/internal/model"
	"gin-chat/pkg/mysql"
)

func TestRepo_FriendCreate(t *testing.T) {
	type args struct {
		ctx    context.Context
		tx     *gorm.DB
		friend *model.FriendModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "FriendCreate",
			args: args{
				ctx: context.Background(),
				tx:  mysql.DB,
				friend: &model.FriendModel{
					UserID:   1,
					FriendID: 2,
					Nickname: "test",
					LookMe:   0,
					LookHim:  0,
					IsStar:   0,
					IsBlack:  0,
					Tags:     "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.FriendCreate(tt.args.ctx, tt.args.tx, tt.args.friend); (err != nil) != tt.wantErr {
				t.Errorf("FriendCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_FriendDelete(t *testing.T) {
	type args struct {
		ctx    context.Context
		friend *model.FriendModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "FriendDelete",
			args: args{
				ctx: context.Background(),
				friend: &model.FriendModel{
					PriID: model.PriID{ID: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.FriendDelete(tt.args.ctx, tt.args.friend); (err != nil) != tt.wantErr {
				t.Errorf("FriendDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_FriendSave(t *testing.T) {
	type args struct {
		ctx    context.Context
		friend *model.FriendModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "FriendSave",
			args: args{
				ctx: context.Background(),
				friend: &model.FriendModel{
					PriID:    model.PriID{ID: 2},
					Nickname: "test123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.FriendSave(tt.args.ctx, tt.args.friend); (err != nil) != tt.wantErr {
				t.Errorf("FriendSave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetFriendAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetFriendAll",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetFriendAll(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len:%v", len(gotList))
		})
	}
}

func TestRepo_GetFriendInfo(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   int
		friendID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetFriendInfo",
			args: args{
				ctx:      context.Background(),
				userID:   1,
				friendID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFriend, err := r.GetFriendInfo(tt.args.ctx, tt.args.userID, tt.args.friendID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFriendInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotFirend: %v", gotFriend)
		})
	}
}
