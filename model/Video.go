package model

import (
	"time"
)

type Video struct {
	Id            int64     `db:"id"`
	Author        int64     `db:"author"`                          // The author
	PlayUrl       string    `db:"play_url"`                        // The play_url
	CoverUrl      string    `db:"cover_url"`                       // The cover_url
	Title         string    `db:"title"`                           // The title
	FavoriteCount int64     `db:"favorite_count" gorm:"default:0"` // The favorite_count
	CommentCount  int64     `db:"comment_count" gorm:"default:0"`  // The comment_count
	CreateAt      time.Time `db:"create_at" gorm:"autoCreateTime"`
}

type VideoResp struct {
	Id            int64    `json:"id"`
	Author        UserResp `json:"author"`          // The author
	PlayUrl       string   `json:"play_url"`        // The play_url
	CoverUrl      string   `json:"cover_url"`       // The cover_url
	Title         string   `json:"title"`           // The title
	FavoriteCount int64    `json:"favorite_count" ` // The favorite_count
	CommentCount  int64    `json:"comment_count" `  // The comment_count
	IsFavorite    bool     `json:"is_favorite"`
}
