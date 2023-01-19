package msg

import (
	"github.com/China-Chris/go_errors/errorss/errors_const"
)

// ErrorCodeTextMapZh 定义中文的业务报错信息
var ErrorCodeTextMapZh = map[int]string{
	errors_const.ErrInternalServer: "内部服务发生错误",
	errors_const.ErrShouldBind:     "参数解析错误请确认参数正确",
	errors_const.ErrSignUp:         "注册时发生错误",
	errors_const.ErrCheckMobile:    "手机号校验不匹配",
	errors_const.ErrGenerateToken:  "生成令牌发生错误",
}
