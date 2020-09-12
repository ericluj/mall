package dao

import (
	"mall/order/model"
)

// CreateOrder ...
func (d *Dao) CreateOrder(userID, productID int64) (*model.Order, error) {
	data := &model.Order{
		UserID:    userID,
		ProductID: productID,
		Status:    model.OrderStatusCreate,
	}
	err := d.db.Create(data).Error

	return data, err
}
