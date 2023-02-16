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
)
