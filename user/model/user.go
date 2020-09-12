package model

import "time"

// GormModel ...
type GormModel struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// User ...
type User struct {
	GormModel
	Name string `json:"name" gorm:"unique_index:name;type:varchar(100);not null;column:name"`
}
