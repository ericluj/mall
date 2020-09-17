package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"mall/order/dao"
	"mall/order/model"
	"mall/order/srv"
	"mall/proto/order"
	"mall/proto/product"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/util/log"
)

type Order struct {
	Broker broker.Broker
}

func (u *Order) CreateOrder(ctx context.Context, req *order.OrderRequest, rsp *order.OrderResponse) error {
	//判断是否有库存
	res, err := srv.ProductSrv.QueryProduct(context.Background(), &product.ProductRequest{Id: req.ProductID})
	if err != nil {
		log.Error("[CreateOrder]", err)
		return err
	}
	if res.Num <= 0 {
		err = errors.New("库存不足")
		log.Error("[CreateOrder]", err)
		return err
	}

	//创建订单
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

	//消息队列
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
