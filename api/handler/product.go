package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"mall/api/srv"
	"mall/proto/product"
	"net/http"
)

func QueryProduct(c *gin.Context) {
	var req = struct {
		Name string `form:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.ProductSrv.QueryProduct(context.Background(), &product.ProductRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateProduct(c *gin.Context) {
	var req = struct {
		Name string `json:"name" binding:"required"`
		Num  int64  `json:"num"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.ProductSrv.CreateProduct(context.Background(), &product.ProductRequest{Name: req.Name, Num: req.Num})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func UpdateProduct(c *gin.Context) {
	var req = struct {
		Name string `json:"name" binding:"required"`
		Num  int64  `json:"num"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.ProductSrv.UpdateProduct(context.Background(), &product.ProductRequest{Name: req.Name, Num: req.Num})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
