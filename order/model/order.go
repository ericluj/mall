package model

import "time"

const (
	OrderStatusCreate = "create"
)

// GormModel ...
type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// Order ...
type Order struct {
	*GormModel
	UserID    int64  `json:"userID" gorm:"index:name;type:int;not null;column:user_id"`
	ProductID int64  `json:"productID" gorm:"index:name;type:int;not null;column:product_id"`
	Status    string `json:"status" gorm:"type:varchar(50);not null;column:status"`
}
