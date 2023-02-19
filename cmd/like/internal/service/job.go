package like

type LikeCacheToDBJob struct {
	Name string
}

// Run 将点赞记录从redis定时刷新到MySQL中
func (likeCacheToDBJob LikeCacheToDBJob) Run() {
	like.RefreshLikeCache()
}
