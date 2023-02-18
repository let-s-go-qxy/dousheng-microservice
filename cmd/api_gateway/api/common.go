package api

// Response 基础返回模型
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// User 用户返回模型
type User struct {
	Id            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
	Avatar        string `json:"avatar"`
}

// Video 视频返回模型
type Video struct {
	Id            int    `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int    `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type Message struct {
	Id         int    `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int    `json:"user_id,omitempty"`
	ToUserId   int    `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int    `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type FavoriteListResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	VideoList  []Video `json:"video_list"`
}
type CommentListResponse struct {
	StatusCode  int32     `json:"status_code"`
	StatusMsg   string    `json:"status_msg,omitempty"`
	CommentList []Comment `json:"comment_list,,omitempty"`
}
type CommentResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	Comment    Comment `json:"comment,,omitempty"`
}

type MessageResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	Message    Message `json:"comment,,omitempty"`
}
