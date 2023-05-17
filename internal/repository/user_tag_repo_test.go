package repository

import (
	"context"
	"testing"

	"gin-chat/internal/model"
)

func TestRepo_GetTagsByUserID(t *testing.T) {
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
			name: "GetTagsByUserID",
			args: args{
				ctx:    context.Background(),
				userID: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetTagsByUserID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_TagBatchCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		tags []*model.UserTagModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TagBatchCreate",
			args: args{
				ctx: context.Background(),
				tags: []*model.UserTagModel{
					{
						UID:  model.UID{UserID: 1},
						Name: "aaa",
					},
					{
						UID:  model.UID{UserID: 2},
						Name: "bb",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := r.TagBatchCreate(tt.args.ctx, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagBatchCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotIds: %v", gotIds)
		})
	}
}
