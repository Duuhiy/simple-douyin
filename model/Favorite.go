package model

import (
	"time"
)

type Favorite struct {
	Id       int64     `db:"id"`
	UserId   int64     `db:"user_id"`
	VideoId  int64     `db:"video_id"`
	CreateAt time.Time `db:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `db:"update_at" gorm:"autoCreateTime"`
}
