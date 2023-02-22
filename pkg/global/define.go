package global

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 常量
const (
	FeedStatusOK    = 1
	FeedStatusError = -1
	FeedStatusNull  = 0
	//CancelFavoriteAction        = 0 //取消点赞
	FavoriteAction              = 1 //点赞
	RequestCancelFavoriteAction = 2 //取消点赞
)

// 状态码
var (
	StatusCodeOk int32 = 0 // 响应状态码 - 成功
	//StatusCodeFail int32 = 1 // 响应状态码 - 一般失败
)

const (
	MessageSendEvent = 1 //发送消息
)

// 数据库、OSS相关
var (
	OssBucket    *oss.Bucket
	MysqlDB      *gorm.DB
	WriteMysqlDB *gorm.DB
	ReadMysqlDB  *gorm.DB
	RedisContext = context.Background()
	DbVerify     *redis.Client
	DbUserLike   *redis.Client
	DbVideoLike  *redis.Client
	Config       = &Conf{
		JWT{
			JwtSecretKey:   "MSR2pH^N6dqqQ5Ns5x!eD2YWVpwzmb3@8RzphRFbEkRwLEra86v3LCB%PvGx$a$L",
			JwtExpiresTime: 604800,
		},
	}
)

type Conf struct {
	JWT
}

// JWT TODO 放到 config.yaml 中，在需要的服务中再读取
type JWT struct {
	JwtSecretKey   string
	JwtExpiresTime int64
}

// Response 响应共有响应头
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

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

// JWT
const (
	TokenValidationErrorMalformed   = "token 格式错误"
	TokenValidationErrorExpired     = "登录状态已失效，请重新登录"
	TokenValidationErrorNotValidYet = "token 尚未激活"
	TokenValid                      = "无效 token"
	TokenHandleFailed               = "无法处理此Token"
	TokenParameterAcquisitionError  = "token参数获取错误"
)

// SaoHua 骚话
var SaoHua = []string{
	"习惯了，是个很强大的短句，它能替代所有一言难尽",
	"你是来和我告别的吗。那就隆重一点，等我眼里装满泪水。",
	"你一定不懂那种看清了还想继续喜欢的感觉",
	"她已经跟我讲清了 再喜欢就真的不礼貌了",
	"能让海王收心的 绝对是个毫无防备心的笨女孩",
	"你说放下了其实还是偷偷哭了好久对吧",
	"偷偷相爱吧，别让世俗知道",
	"凡是和你有点像的人 我总会多看两眼",
	"属你最得我意，也是你最不识抬举",
	"你就是仗着我喜欢你,才这么字字珠玑的伤害我",
}
