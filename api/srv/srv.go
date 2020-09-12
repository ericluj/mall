package srv

import (
	"mall/proto/user"
)

var UserSrv user.UserService

func Init(userSrv user.UserService) {
	UserSrv = userSrv
}
