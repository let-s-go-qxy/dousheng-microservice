package main

import (
	"context"
	feed "dousheng/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeedList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeedList(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
