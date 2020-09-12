package dao

import (
	"mall/product/model"
)

// CreateProduct ...
func (d *Dao) CreateProduct(name string, num int64) (*model.Product, error) {
	data := &model.Product{
		Name: name,
		Num:  num,
	}
	err := d.db.Create(data).Error

	return data, err
}

// QueryProduct ...
func (d *Dao) QueryProduct(name string) (user model.Product, err error) {
	if err = d.db.Where(&model.Product{Name: name}).Order("id desc").First(&user).Error; err != nil {
		return
	}
	return
}

// UpdateProduct ...
func (d *Dao) UpdateProduct(name string, num int64) (user model.Product, err error) {
	if err = d.db.Model(&user).Where(&model.Product{Name: name}).Update("num", num).Error; err != nil {
		return
	}
	return
}
