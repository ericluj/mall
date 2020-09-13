package model

import "time"

const (
	TopicOrder = "order"
)

// GormModel ...
type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// Product ...
type Product struct {
	GormModel
	Name string `json:"name" gorm:"unique_index:name;type:varchar(100);not null;column:name"`
	Num  int64  `json:"num" gorm:"type:int;not null;column:num"`
}
