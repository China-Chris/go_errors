package errorss

import (
	"fmt"
	"github.com/China-Chris/go_errors/errorss/msg"
	"github.com/beego/beego/v2/adapter/logs"
)

//HandleCustomError 统一自定义错误处理
func HandleCustomError(code int, msg string, err error) error {
	err = fmt.Errorf("code:%d, %s，err:%s", code, msg, err)
	logs.Error(err)
	return err
}

//HandleError 统一错误处理程序 中zh，英en，日ja，韩ko(默认中文)
func HandleError(code int, language string, err error) error {
	var ErrorCodeTextMap map[int]string
	switch language {
	case "zn":
		ErrorCodeTextMap = msg.ErrorCodeTextMapZh
	case "en":
		ErrorCodeTextMap = msg.ErrorCodeTextMapEn
	default:
		ErrorCodeTextMap = msg.ErrorCodeTextMapZh
	}
	return HandleCustomError(code, ErrorCodeTextMap[code], err)
}
