package repository

import (
	"context"
	"testing"

	"gin-chat/internal/model"
)

func TestRepo_CommentCreate(t *testing.T) {
	type args struct {
		ctx   context.Context
		model *model.MomentCommentModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CommentCreate",
			args: args{
				ctx: context.Background(),
				model: &model.MomentCommentModel{
					UID:      model.UID{UserID: 1},
					ReplyID:  2,
					MomentID: 1,
					Content:  "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.CommentCreate(tt.args.ctx, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommentCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId: %v", gotId)
		})
	}
}

func TestRepo_GetCommentsByMomentID(t *testing.T) {
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
			name: "GetCommentsByMomentID",
			args: args{
				ctx:      context.Background(),
				momentID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetCommentsByMomentID(tt.args.ctx, tt.args.momentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentsByMomentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_GetCommentsByMomentIds(t *testing.T) {
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
			name: "GetCommentsByMomentIds",
			args: args{
				ctx: context.Background(),
				ids: []int{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMComments, err := r.GetCommentsByMomentIds(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommentsByMomentIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotMComments len: %v", len(gotMComments))
		})
	}
}
