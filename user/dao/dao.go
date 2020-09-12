package dao

import (
	"fmt"

	"mall/user/conf"
	"mall/user/model"

	"github.com/jinzhu/gorm"
)

// Dao ...
type Dao struct {
	db *gorm.DB
}

var defaultDao *Dao

// GetDao ...
func GetDao() *Dao {
	return defaultDao
}

// Init 初始化数据库连接
func Init(c *conf.Config) (err error) {
	defaultDao, err = newDao(c)

	return
}

func newDao(c *conf.Config) (d *Dao, err error) {
	d = &Dao{}
	if d.db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DB.User, c.DB.PassWord, c.DB.Host, c.DB.Port, c.DB.DBName)); err != nil {
		panic(err)
	}
	d.db.SingularTable(true)       //表名采用单数形式
	d.db.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	d.db.DB().SetMaxIdleConns(10)  //SetMaxIdleConns用于设置闲置的连接数
	d.db.LogMode(true)

	d.db.AutoMigrate(&model.User{})

	return
}

// Ping ...
func (d *Dao) Ping() (err error) {
	if err = d.db.DB().Ping(); err != nil {
		return
	}
	return
}

// Disconnect ...
func (d *Dao) Disconnect() error {
	return d.db.DB().Close()
}
