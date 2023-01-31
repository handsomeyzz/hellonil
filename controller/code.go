package controller

type ResCode uint64

const (
	CodeSuccess      = "success"
	CodeInvalidParam = "请求参数错误"

	CodeUserToLength = "用户名或密码过长"

	CodeUserExist       = "用户名已存在"
	CodeUserNotExist    = "用户名不存在"
	CodeInvalidPassword = "用户名或密码错误"

	CodeServerBusy = "服务繁忙"

	CodeNeedLogin    = "需要登录"
	CodeInvalidToken = "无效的token"
	//上传列表
	CodeUploadSuccess = ""
	CodeUploadFail    = ""

	//评论
	CodeCommentSuccess = ""
	CodeCommentFail    = ""
	//关注
	CodeFocusSuccess = ""
	CodeFocusFail    = ""

	CodeSendSuccess = ""
	CodeSendFail    = ""
)
