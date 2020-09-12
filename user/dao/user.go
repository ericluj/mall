package dao

import (
	"mall/user/model"
)

func (d *Dao) CreateUser(name string) (*model.User, error) {
	data := &model.User{
		Name: name,
	}
	err := d.db.Create(data).Error

	return data, err
}

// User ...
func (d *Dao) QueryUser(name string) (user model.User, err error) {
	if err = d.db.Where(&model.User{Name: name}).Order("id desc").First(&user).Error; err != nil {
		return
	}
	return
}
