package dao

import (
	"github.com/jinzhu/gorm"
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
func (d *Dao) QueryProduct(name string, id int64) (product model.Product, err error) {
	if id > 0 {
		if err = d.db.Where("id=?", id).Order("id desc").First(&product).Error; err != nil {
			return
		}
	} else {
		if err = d.db.Where("name=?", name).Order("id desc").First(&product).Error; err != nil {
			return
		}
	}

	return
}

// UpdateProduct ...
func (d *Dao) UpdateProduct(name string, num int64) (product model.Product, err error) {
	if err = d.db.Model(&product).Where(&model.Product{Name: name}).Update("num", num).Error; err != nil {
		return
	}
	return
}

// UpdateProduct ...
func (d *Dao) ProductNumDecrby(id int64) (err error) {
	if err = d.db.Model(&model.Product{}).Where("id=?", id).Update("num", gorm.Expr("num - ?", 1)).Error; err != nil {
		return
	}
	return
}
