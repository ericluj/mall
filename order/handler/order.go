package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"mall/order/dao"
	"mall/proto/order"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/util/log"
)

type Order struct {
	Broker broker.Broker
}

func (u *Order) CreateOrder(ctx context.Context, req *order.OrderRequest, rsp *order.OrderResponse) error {
	data, err := dao.GetDao().CreateOrder(req.GetUserID(), req.GetProductID())
	if err != nil {
		log.Error("[CreateOrder]", err)
		return err
	}

	// Create a broker message
	msgBody, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{
			"id": strconv.FormatInt(data.UserID, 10),
		},
		Body: msgBody,
	}
	if err := u.Broker.Publish("order", msg); err != nil {
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
