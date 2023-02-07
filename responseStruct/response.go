package responseStruct

// /douyin/feed/  和/douyin/publish/list/ 和/douyin/favorite/list
type DouYinResponse struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg,omitempty"`
	VideoList  []*Video `json:"video_list"`
	NextTime   int64    `json:"next_time,omitempty"`
}

type Video struct {
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count "`
	ID            int64  `json:"id"`
	IsFavorite    bool   `json:"is_favorite"`
}

type User struct {
	ID            int64  `json:"id" db:"user_id"`
	FollowCount   int64  `json:"follow_count" db:"follow_count"`
	FollowerCount int64  `json:"follower_count" db:"follower_count"`
	Name          string `json:"name" db:"name" db:"name"`
	IsFollow      bool   `json:"is_follow" db:"is_follow"`
}

// /douyin/user/register 和  /douyin/user/login
type DouYinRegisterAndLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

// /douyin/user/
type DouYinUserResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	User       User   `json:"user"`
}

// /douyin/publish/action/   /douyin/favorite/action   /douyin/relation/action/
type DouYinPublishActionResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// /douyin/comment/action
type DouYinCommentActionResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg,omitempty"`
	Comment    Comment `json:"comment,omitempty"`
}
type Comment struct {
	ID         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

// /douyin/comment/list/
type DouYinCommentListResponse struct {
	StatusCode  int32     `json:"status_code"`
	StatusMsg   string    `json:"status_msg,omitempty"`
	CommentList []Comment `json:"comment_list"`
}

// /douyin/relatioin/follow/list/ 和/douyin/relation/follower/list/ -
type DouYinRelationFollowList struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserList   []User `json:"user_list"`
}
