package status

const (
	SUCCESS        = 200
	UserADDSUCCESS = "注册成功,请记住唯一UID"
)

var (
	OK  = response(200, "OK")    // 通用成功
	Err = response(500, "ERROR") // 通用错误

	ParamErr      = response(1001, "获取参数失败")
	CreateErr     = response(1002, "创建礼包码失败")
	CodeLenErr    = response(1003, "礼包码输入错误")
	FindCodeErr   = response(1004, "查询礼品码失败")
	VerifyCodeErr = response(1005, "礼品码验证失败")
	CodeTypeErr   = response(1006, "礼品码类型错误")
	CodeUIDErr    = response(1007, "请输用户Uid")
	RedisErr      = response(1008, "redis储存异常")
	MarshalErr    = response(1009, "序列化异常")
	StringErr     = response(1010, "字符串不能为空")
	LoginUserErr  = response(1011, "用户登陆失败")
	RegUserErr    = response(1012, "用户创建失败")
	MongoDBErr    = response(1013, "MongoDB链接异常")
	RedisConErr   = response(1014, "redis链接异常")
	DBUpdateErr   = response(1015, "用户数据更新异常")
	DBInsertErr   = response(1016, "用户新建数据异常")
	CodeErr       = response(1017, "礼包码无效")
	UserADD       = response(1018, "注册成功")
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 错误描述
	Data interface{} `json:"data"` // 返回数据
}

// 自定义响应信息

func (res *Response) WithMsg(message string) Response {
	return Response{
		Code: res.Code,
		Msg:  message,
		Data: res.Data,
	}
}

// 追加响应数据

func (res *Response) WithData(data interface{}) Response {
	return Response{
		Code: res.Code,
		Msg:  res.Msg,
		Data: data,
	}
}

// 构造函数
func response(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
