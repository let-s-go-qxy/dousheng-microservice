package msg

// video
const (
	HasNoVideoMsg            = "已经没有视频了"
	GetVideoInfoSuccessMsg   = "获取视频信息成功"
	GetVideoInfoErrorMsg     = "获取视频信息失败，出错了"
	PublishVideoFailedMsg    = "上传视频失败"
	PublishVideoSuccessMsg   = "上传视频成功"
	GetPublishListFailedMsg  = "获取已发布视频失败"
	GetPublishListSuccessMsg = "获取已发布视频成功"
)

const (
	LikeFavoriteActionFail    = "点赞失败"
	LikeFavoriteActionSuccess = "点赞成功"
)

// JWT
const (
	TokenValidationErrorMalformed   = "token 格式错误"
	TokenValidationErrorExpired     = "登录状态已失效，请重新登录"
	TokenValidationErrorNotValidYet = "token 尚未激活"
	TokenValid                      = "无效 token"
	TokenHandleFailed               = "无法处理此Token"
	TokenParameterAcquisitionError  = "token参数获取错误"
)
