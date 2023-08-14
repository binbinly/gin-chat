package repository

import (
	"context"
	"gin-chat/pkg/app"
	"testing"

	"gorm.io/gorm"

	"gin-chat/internal/model"
)

func TestRepo_GetGroupUserByID(t *testing.T) {
	type args struct {
		ctx     context.Context
		userID  int
		groupID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetGroupUserByID",
			args: args{
				ctx:     context.Background(),
				userID:  1,
				groupID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := r.GetGroupUserByID(tt.args.ctx, tt.args.userID, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotInfo: %v", gotInfo)
		})
	}
}

func TestRepo_GroupUserAll(t *testing.T) {
	type args struct {
		ctx     context.Context
		groupID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserAll",
			args: args{
				ctx:     context.Background(),
				groupID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GroupUserAll(tt.args.ctx, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupUserAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_GroupUserBatchCreate(t *testing.T) {
	type args struct {
		ctx   context.Context
		tx    *gorm.DB
		users []*model.GroupUserModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserBatchCreate",
			args: args{
				ctx: context.Background(),
				tx:  app.DB,
				users: []*model.GroupUserModel{
					{
						UID:      model.UID{UserID: 1},
						GroupID:  2,
						Nickname: "test",
					},
					{
						UID:      model.UID{UserID: 2},
						GroupID:  2,
						Nickname: "test2",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupUserBatchCreate(tt.args.ctx, tt.args.tx, tt.args.users); (err != nil) != tt.wantErr {
				t.Errorf("GroupUserBatchCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GroupUserCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.GroupUserModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserCreate",
			args: args{
				ctx: context.Background(),
				user: &model.GroupUserModel{
					UID:      model.UID{UserID: 1},
					GroupID:  2,
					Nickname: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupUserCreate(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("GroupUserCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GroupUserDelete(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.GroupUserModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserDelete",
			args: args{
				ctx: context.Background(),
				user: &model.GroupUserModel{
					PriID: model.PriID{ID: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupUserDelete(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("GroupUserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GroupUserDeleteByGroupID(t *testing.T) {
	type args struct {
		ctx     context.Context
		tx      *gorm.DB
		groupID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserDeleteByGroupID",
			args: args{
				ctx:     context.Background(),
				tx:      app.DB,
				groupID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupUserDeleteByGroupID(tt.args.ctx, tt.args.tx, tt.args.groupID); (err != nil) != tt.wantErr {
				t.Errorf("GroupUserDeleteByGroupID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GroupUserUpdateNickname(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   int
		groupID  int
		nickname string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupUserUpdateNickname",
			args: args{
				ctx:      context.Background(),
				userID:   1,
				groupID:  2,
				nickname: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupUserUpdateNickname(tt.args.ctx, tt.args.userID, tt.args.groupID, tt.args.nickname); (err != nil) != tt.wantErr {
				t.Errorf("GroupUserUpdateNickname() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
