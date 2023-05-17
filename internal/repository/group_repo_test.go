package repository

import (
	"context"
	"testing"

	"gorm.io/gorm"

	"gin-chat/internal/model"
	"gin-chat/pkg/mysql"
)

func TestRepo_GetGroupByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetGroupByID",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := r.GetGroupByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotInfo: %v", gotInfo)
		})
	}
}

func TestRepo_GetGroupsByUserID(t *testing.T) {
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
			name: "GetGroupsByUserID",
			args: args{
				ctx:    context.Background(),
				userID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetGroupsByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len:%v", len(gotList))
		})
	}
}

func TestRepo_GroupCreate(t *testing.T) {
	type args struct {
		ctx   context.Context
		tx    *gorm.DB
		group *model.GroupModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupCreate",
			args: args{
				ctx: context.Background(),
				tx:  mysql.DB,
				group: &model.GroupModel{
					UID:           model.UID{UserID: 1},
					Name:          "test",
					Avatar:        "",
					Remark:        "",
					InviteConfirm: 0,
					Status:        0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.GroupCreate(tt.args.ctx, tt.args.tx, tt.args.group)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId: %v", gotId)
		})
	}
}

func TestRepo_GroupDelete(t *testing.T) {
	type args struct {
		ctx   context.Context
		tx    *gorm.DB
		group *model.GroupModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupDelete",
			args: args{
				ctx: context.Background(),
				tx:  mysql.DB,
				group: &model.GroupModel{
					PriID: model.PriID{ID: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupDelete(tt.args.ctx, tt.args.tx, tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("GroupDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GroupSave(t *testing.T) {
	type args struct {
		ctx   context.Context
		group *model.GroupModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GroupSave",
			args: args{
				ctx: context.Background(),
				group: &model.GroupModel{
					PriID:         model.PriID{ID: 1},
					UID:           model.UID{UserID: 1},
					Name:          "test123",
					Avatar:        "",
					Remark:        "",
					InviteConfirm: 0,
					Status:        0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.GroupSave(tt.args.ctx, tt.args.group); (err != nil) != tt.wantErr {
				t.Errorf("GroupSave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
