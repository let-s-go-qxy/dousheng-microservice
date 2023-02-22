package like

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/henrylee2cn/goutil/calendar/cron"
)

type LikeCacheToDBJob struct {
	Name string
}

func CronTaskSetUp() {
	c := cron.New()
	job := LikeCacheToDBJob{Name: "LikeRefresh"}
	err := c.AddJob("0/30 * * * * ?", job)
	if err != nil {
		hlog.Error(err.Error())
		return
	}
	go c.Start()
	defer c.Stop()
}

// Run 将点赞记录从redis定时刷新到MySQL中
func (likeCacheToDBJob LikeCacheToDBJob) Run() {
	like.RefreshLikeCache()
}
