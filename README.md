# go_errors
统一的自定义错误处理程序

install
```go
go get -u github.com/China-Chris/go_errors
```
使用在具体的业务中调用HandleError函数
```text 例如
accessToken, refreshToken, err := jwt.GenerateToken(that.Phone, that.Password)
if err != nil {
	return nil, errorss.HandleError(errors_const.ErrGenerateToken, "zn", err)
}
```
函数签名
```text
//HandleError 统一错误处理程序 错误码，中zh，英en，日ja，韩ko(默认中文)，透传的错误内容
func HandleError(code int, language string, err error) error  
```

自定义错误:我想要增加错误怎么办
找到errors_const.go文件  (您可以自己添加常量).并在zn_error_map.go或其他语言的error_map.go中添加您定义的错误常量并定义错误信息
```text
errors_const.go
// 用于存放各模块定义的错误常量，如过业务错误太多可以单独将各模块分成多个文件存放 例：user_const.go 存放 用户模块定义的错误

const ( //定义常规错误
	ErrInternalServer = iota + 500 //Post 解析失败 参数错误
)

const ( //  用户模块
	ErrShouldBind = iota + 1000 //Post 解析失败 参数错误
	ErrSignUp
	ErrCheckMobile
	ErrGenerateToken
)

zn_error_map.go
// ErrorCodeTextMapZh 定义中文的业务报错信息
var ErrorCodeTextMapZh = map[int]string{
	errors_const.ErrInternalServer: "内部服务发生错误",
	errors_const.ErrShouldBind:     "参数解析错误请确认参数正确",
	errors_const.ErrSignUp:         "注册时发生错误",
	errors_const.ErrCheckMobile:    "手机号校验不匹配",
	errors_const.ErrGenerateToken:  "生成令牌发生错误",
}


```

