package handler

import (
	"context"
	"net/http"

	"mall/api/srv"
	"mall/proto/order"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var req = struct {
		UserID    int64 `json:"userID" binding:"required"`
		ProductID int64 `json:"productID" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.OrderSrv.CreateOrder(context.Background(), &order.OrderRequest{UserID: req.UserID, ProductID: req.ProductID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
