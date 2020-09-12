package handler

import (
	"context"
	"mall/proto/order"

	"github.com/micro/go-micro/v2/util/log"
	"mall/order/dao"
)

type Order struct{}

func (u *Order) CreateOrder(ctx context.Context, req *order.OrderRequest, rsp *order.OrderResponse) error {
	data, err := dao.GetDao().CreateOrder(req.GetUserID(), req.GetProductID())
	if err != nil {
		log.Error("[CreateOrder]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	rsp.UserID = data.UserID
	rsp.ProductID = data.ProductID
	rsp.Status = data.Status
	return nil
}
