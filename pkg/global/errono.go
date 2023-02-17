package global

import "errors"

// 错误大全
var (
	// 响应错误码
	StatusCodeFail int32 = 1

	// 数据库错误
	ErrDbCreateUniqueKeyRepeatedly error = errors.New("ErrDbCreateUniqueKeyRepeatedly") // 重复创建按了应该唯一的key的一条记录
)

var (
	StatusOk int32 = 0
)
