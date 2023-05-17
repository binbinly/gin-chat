package repository

import (
	"context"
	"testing"

	"gin-chat/internal/model"
)

func TestRepo_CollectCreate(t *testing.T) {
	type args struct {
		ctx     context.Context
		collect *model.CollectModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CollectCreate",
			args: args{
				ctx: context.Background(),
				collect: &model.CollectModel{
					UID:     model.UID{UserID: 1},
					Content: "test",
					Type:    0,
					Options: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, err := r.CollectCreate(tt.args.ctx, tt.args.collect)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollectCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotId:%v", gotId)
		})
	}
}

func TestRepo_CollectDelete(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID int
		id     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CollectDelete",
			args: args{
				ctx:    context.Background(),
				userID: 1,
				id:     1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.CollectDelete(tt.args.ctx, tt.args.userID, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CollectDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_GetCollectsByUserID(t *testing.T) {
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
			name: "GetCollectsByUserID",
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
			gotList, err := r.GetCollectsByUserID(tt.args.ctx, tt.args.userID, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCollectsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}
