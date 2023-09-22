package model

import (
	"time"
)

type Comment struct {
	Id       int64     `db:"id"`
	UserId   int64     `db:"user_id"`
	VideoId  int64     `db:"video_id"`
	Contents string    `db:"contents"`
	CreateAt time.Time `db:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `db:"update_at" gorm:"autoCreateTime"`
}

type CommentResp struct {
	Id         int64    `json:"id"`
	User       UserResp `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}
