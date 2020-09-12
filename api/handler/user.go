package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"mall/api/srv"
	"mall/proto/user"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func QueryUser(c *gin.Context) {
	var req = struct {
		Name string `form:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.UserSrv.QueryUser(context.Background(), &user.UserRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUser(c *gin.Context) {
	var req = struct {
		Name string `json:"name" binding:"required"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	res, err := srv.UserSrv.CreateUser(context.Background(), &user.UserRequest{Name: req.Name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
