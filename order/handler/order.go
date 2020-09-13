package handler

import (
	"context"
	"encoding/json"
	"mall/order/dao"
	"mall/order/model"
	"mall/proto/order"
	"strconv"

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

	rsp.Id = int64(data.ID)
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	rsp.UserID = data.UserID
	rsp.ProductID = data.ProductID
	rsp.Status = data.Status

	// Create a broker message
	msgBody, err := json.Marshal(rsp)
	if err != nil {
		log.Fatal(err.Error())
	}
	msg := &broker.Message{
		Header: map[string]string{
			"id": strconv.FormatInt(data.UserID, 10),
		},
		Body: msgBody,
	}
	if err := u.Broker.Publish(model.TopicOrder, msg); err != nil {
		return err
	}

	return nil
}
