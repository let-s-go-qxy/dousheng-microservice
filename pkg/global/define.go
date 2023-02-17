package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// 常量
const (
	CancelFavoriteAction        = 0 //取消点赞
	FavoriteAction              = 1 //点赞
	RequestCancelFavoriteAction = 2
)

const (
	MessageSendEvent = 1 //发送消息
)

// 数据库相关
var (
	MysqlDB      *gorm.DB
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
