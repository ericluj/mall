package srv

import (
	"mall/proto/order"
	"mall/proto/product"
	"mall/proto/user"
)

var (
	UserSrv    user.UserService
	ProductSrv product.ProductService
	OrderSrv   order.OrderService
)

func Init(userSrv user.UserService, productSrv product.ProductService, orderSrv order.OrderService) {
	UserSrv = userSrv
	ProductSrv = productSrv
	OrderSrv = orderSrv
}
