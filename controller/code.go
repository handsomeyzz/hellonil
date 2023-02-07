package controller

const (
	CodeStatusOK = iota
	CodeStatusFail
	CodeServerFail

	CodeSuccess
	CodeInvalidParam
	CodeLoginOk
	CodeUserToLength
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeRegisterOk
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
	//上传列表
	CodeUploadSuccess
	CodeUploadFail
	CodeUpParamFail

	//评论
	CodeCommentSuccess
	CodeCommentFail
	//关注
	CodeFocusSuccess
	CodeFocusFail

	CodeSendSuccess
	CodeSendFail
)

var codeString = map[int]string{
	CodeSuccess:      "success",
	CodeInvalidParam: "请求参数错误",
	CodeServerFail:   "服务器内部出错",

	CodeRegisterOk:      "注册成功",
	CodeUserToLength:    "用户名或密码过长",
	CodeUserExist:       "用户名已存在!请登录",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeLoginOk:         "登录成功",

	CodeServerBusy: "服务繁忙",

	CodeNeedLogin:    "需要登录",
	CodeInvalidToken: "无效的token",
	//上传列表
	CodeUploadSuccess: "文件上传成功",
	CodeUploadFail:    "文件上传失败",
	CodeUpParamFail:   "文件上传失败!参数不符合规定",

	//评论
	CodeCommentSuccess: "",
	CodeCommentFail:    "",
	//关注
	CodeFocusSuccess: "",
	CodeFocusFail:    "",

	CodeSendSuccess: "",
	CodeSendFail:    "",
}
