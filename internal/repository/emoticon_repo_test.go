package repository

import (
	"context"
	"testing"
)

func TestRepo_GetEmoticonCatAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetEmoticonCatAll",
			args: args{ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetEmoticonCatAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmoticonCatAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}

func TestRepo_GetEmoticonListByCat(t *testing.T) {
	type args struct {
		ctx context.Context
		cat string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetEmoticonListByCat",
			args: args{
				ctx: context.Background(),
				cat: "贡献🇨🇳",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetEmoticonListByCat(tt.args.ctx, tt.args.cat)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmoticonListByCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotList len: %v", len(gotList))
		})
	}
}
