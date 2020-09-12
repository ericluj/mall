package handler

import (
	"context"
	"mall/proto/user"
)

type User struct{}

func (u *User) CreateUser(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	return nil
}

func (u *User) QueryUser(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	return nil
}
