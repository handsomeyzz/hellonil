package models

import "time"

type Accounts struct {
	Username string `db:"username"`
	Password string `db:"password"`
	ID       int64  `db:"id"`
}

type Comments struct {
	Status    uint8     `db:"status"`
	ID        int64     `db:"id"`
	VideoID   int64     `db:"video_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatDate time.Time `db:"creat_date"`
}
type Follows struct {
	ID       int64 `db:"id"`
	UserID   int64 `db:"user_id"`
	TargetID int64 `db:"target_id"`
}
type Likes struct {
	ID      int64 `db:"id"`
	VideoID int64 `db:"video_id"`
	UserID  int64 `db:"user_id"`
	Status  uint8 `db:"status"`
}
type Message struct {
	ID         int64     `db:"id"`
	FromUserID int64     `db:"from_user_id"`
	ToUserID   int64     `db:"to_user_id"`
	Content    int64     `db:"content"`
	CreatDate  time.Time `db:"creat_date"`
}
type Users struct {
	ID            int64  `db:"id"`
	UserID        int64  `db:"user_id"`
	FollowCount   int64  `db:"follow_count"`
	FollowerCount int64  `db:"follower_count"`
	Avatar        string `db:"avatar"`
}

type videos struct {
	ID            int64  `db:"id"`
	AuthorID      int64  `db:"author_id"`
	PlayUrl       int64  `db:"play_url"`
	CoverUrl      int64  `db:"cover_url"`
	FavoriteCount int64  `db:"favorite_count"`
	CommentCount  int64  `db:"comment_count"`
	title         string `db:"title"`
}
