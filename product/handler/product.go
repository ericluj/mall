package handler

import (
	"context"
	"mall/product/dao"
	"mall/proto/product"

	"github.com/micro/go-micro/v2/util/log"
)

type Product struct{}

func (u *Product) CreateProduct(ctx context.Context, req *product.ProductRequest, rsp *product.ProductResponse) error {
	data, err := dao.GetDao().CreateProduct(req.GetName(), req.GetNum())
	if err != nil {
		log.Error("[CreateProduct]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	rsp.Name = data.Name
	rsp.Num = data.Num
	return nil
}

func (u *Product) QueryProduct(ctx context.Context, req *product.ProductRequest, rsp *product.ProductResponse) error {
	data, err := dao.GetDao().QueryProduct(req.GetName(), req.GetId())
	if err != nil {
		log.Error("[QueryProduct]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.Name = data.Name
	rsp.Num = data.Num
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	return nil
}

func (u *Product) UpdateProduct(ctx context.Context, req *product.ProductRequest, rsp *product.ProductResponse) error {
	data, err := dao.GetDao().UpdateProduct(req.GetName(), req.GetNum())
	if err != nil {
		log.Error("[UpdateProduct]", err)
		return err
	}

	rsp.Id = int64(data.ID)
	rsp.Name = data.Name
	rsp.Num = data.Num
	rsp.CreatedAt = data.CreatedAt.Unix()
	rsp.UpdatedAt = data.UpdatedAt.Unix()
	return nil
}
