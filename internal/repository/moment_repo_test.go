package repository

import (
	"context"
	"testing"

	"gorm.io/gorm"

	"gin-chat/internal/model"
	"gin-chat/pkg/dbs"
)

func TestRepo_GetMomentByID(t *testing.T) {
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
			name: "GetMomentByID",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMoment, err := r.GetMomentByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMomentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotMoment: %v", gotMoment)
		})
	}
}

func TestRepo_GetMomentsByIds(t *testing.T) {
	type args struct {
		ctx context.Context
		ids []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetMomentsByIds",
			args: args{
				ctx: context.Background(),
				ids: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMoments, err := r.GetMomentsByIds(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMomentsByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotMoments len: %v", len(gotMoments))
		})
	}
}

func TestRepo_GetMomentsByUserID(t *testing.T) {
	type args struct {
		ctx    context.Context
		myID   int
		userID int
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetMomentsByUserID",
			args: args{
				ctx:    context.Background(),
				myID:   1,
				userID: 2,
				offset: 0,
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetMomentsByUserID(tt.args.ctx, tt.args.myID, tt.args.userID, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMomentsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_GetMyMoments(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetMyMoments",
			args: args{
				ctx:    context.Background(),
				userID: 1,
				offset: 0,
				limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetMyMoments(tt.args.ctx, tt.args.userID, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMyMoments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_MomentCreate(t *testing.T) {
	type args struct {
		ctx    context.Context
		tx     *gorm.DB
		moment *model.MomentModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "MomentCreate",
			args: args{
				ctx: context.Background(),
				tx:  dbs.DB,
				moment: &model.MomentModel{
					UID:      model.UID{UserID: 1},
					Content:  "test",
					Image:    "",
					Video:    "",
					Location: "",
					Remind:   "",
					Type:     0,
					SeeType:  0,
					See:      "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.MomentCreate(tt.args.ctx, tt.args.tx, tt.args.moment)
			if (err != nil) != tt.wantErr {
				t.Errorf("MomentCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId: %v", gotId)
		})
	}
}
