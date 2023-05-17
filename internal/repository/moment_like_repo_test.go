package repository

import (
	"context"
	"testing"

	"gin-chat/internal/model"
)

func TestRepo_GetLikeUserIdsByMomentID(t *testing.T) {
	type args struct {
		ctx      context.Context
		momentID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetLikeUserIdsByMomentID",
			args: args{
				ctx:      context.Background(),
				momentID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserIds, err := r.GetLikeUserIdsByMomentID(tt.args.ctx, tt.args.momentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLikeUserIdsByMomentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotUserIds: %v", gotUserIds)
		})
	}
}

func TestRepo_GetLikesByMomentIds(t *testing.T) {
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
			name: "GetLikesByMomentIds",
			args: args{
				ctx: context.Background(),
				ids: []int{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLikes, err := r.GetLikesByMomentIds(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLikesByMomentIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotLikes len: %v", len(gotLikes))
		})
	}
}

func TestRepo_LikeCreate(t *testing.T) {
	type args struct {
		ctx   context.Context
		model *model.MomentLikeModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "LikeCreate",
			args: args{
				ctx: context.Background(),
				model: &model.MomentLikeModel{
					UID:      model.UID{UserID: 1},
					MomentID: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.LikeCreate(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("LikeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId: %v", gotId)
		})
	}
}

func TestRepo_LikeDelete(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   int
		momentID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "LikeDelete",
			args: args{
				ctx:      context.Background(),
				userID:   1,
				momentID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.LikeDelete(tt.args.ctx, tt.args.userID, tt.args.momentID); (err != nil) != tt.wantErr {
				t.Errorf("LikeDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
