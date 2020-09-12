package handler

import (
	"context"

	"mall/proto/user"
	"mall/user/dao"

	"github.com/micro/go-micro/v2/util/log"
)

type User struct{}

func (u *User) CreateUser(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	data, err := dao.GetDao().CreateUser(req.GetName())
	if err != nil {
		log.Error("[CreateUser]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	rsp.Name = data.Name
	return nil
}

func (u *User) QueryUser(ctx context.Context, req *user.UserRequest, rsp *user.UserResponse) error {
	data, err := dao.GetDao().QueryUser(req.GetName())
	if err != nil {
		log.Error("[QueryUser]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.Name = data.Name
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	return nil
}
