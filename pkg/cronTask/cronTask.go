package cronTask

import (
	"dousheng/cmd/api_gateway/api"
	"github.com/henrylee2cn/goutil/calendar/cron"
)

func CronTaskSetUp() {
	c := cron.New()

	c.AddFunc("0/30 * * * * ?", api.RefreshLikeCache)
	go c.Start()
	defer c.Stop()
}
