package msg

import (
	"github.com/China-Chris/go_errors/errorss/errors_const"
)

//ErrorCodeTextMapEn 定义英文的业务报错信息
var ErrorCodeTextMapEn = map[int]string{
	errors_const.ErrInternalServer: "Internal server error",
	errors_const.ErrShouldBind:     "Parameter parsing error, please check the parameters",
	errors_const.ErrSignUp:         "Error occurred during registration",
	errors_const.ErrCheckMobile:    "Mobile number verification does not match",
}
