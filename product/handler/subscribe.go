package handler

import (
	"encoding/json"
	"fmt"

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
	fmt.Println(data)
	return nil
}
