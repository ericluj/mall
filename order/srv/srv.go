package srv

import (
	"mall/proto/product"
	"mall/proto/user"
)

var (
	UserSrv    user.UserService
	ProductSrv product.ProductService
)

func Init(userSrv user.UserService, productSrv product.ProductService) {
	UserSrv = userSrv
	ProductSrv = productSrv
}
