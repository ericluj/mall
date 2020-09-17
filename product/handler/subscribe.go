package handler

import (
	"encoding/json"
	"mall/product/dao"

	"mall/proto/order"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/util/log"
)

func Order(e broker.Event) error {
	var data *order.OrderResponse
	if err := json.Unmarshal(e.Message().Body, &data); err != nil {
		log.Error(err)
		return err
	}
	if err := dao.GetDao().ProductNumDecrby(data.ProductID); err != nil {
		log.Error("[ProductNumDecrby]", err)
		return err

	}
	return nil
}
