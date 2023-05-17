package repository

import (
	"context"
	"testing"

	"gin-chat/internal/model"
)

func TestRepo_GetUserByID(t *testing.T) {
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
			name: "GetUserByID",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := r.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotUser: %v", gotUser)
		})
	}
}

func TestRepo_GetUserByPhone(t *testing.T) {
	type args struct {
		ctx   context.Context
		phone int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetUserByPhone",
			args: args{
				ctx:   context.Background(),
				phone: 13333333333,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := r.GetUserByPhone(tt.args.ctx, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotUser: %v", gotUser)
		})
	}
}

func TestRepo_GetUserByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetUserByUsername",
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUser, err := r.GetUserByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotUser: %v", gotUser)
		})
	}
}

func TestRepo_GetUsersByIds(t *testing.T) {
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
			name: "GetUsersByIds",
			args: args{
				ctx: context.Background(),
				ids: []int{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsers, err := r.GetUsersByIds(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersByIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("gotUsers len: %v", len(gotUsers))
		})
	}
}

func TestRepo_UserCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.UserModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "UserCreate",
			args: args{
				ctx: context.Background(),
				user: &model.UserModel{
					Username: "test",
					Password: "123456",
					Phone:    13333333333,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := r.UserCreate(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_UserExist(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		phone    int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "UserExist",
			args: args{
				ctx:      context.Background(),
				username: "test",
				phone:    13333333333,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.UserExist(tt.args.ctx, tt.args.username, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_UserUpdate(t *testing.T) {
	type args struct {
		ctx     context.Context
		id      int
		userMap map[string]any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "UserUpdate",
			args: args{
				ctx: context.Background(),
				id:  1,
				userMap: map[string]any{
					"nickname": "aaa",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.UserUpdate(tt.args.ctx, tt.args.id, tt.args.userMap); (err != nil) != tt.wantErr {
				t.Errorf("UserUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepo_UserUpdatePwd(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.UserModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "UserUpdatePwd",
			args: args{
				ctx: context.Background(),
				user: &model.UserModel{
					PriID:    model.PriID{ID: 1},
					Password: "123123",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := r.UserUpdatePwd(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserUpdatePwd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
