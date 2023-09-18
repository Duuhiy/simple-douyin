package model

import (
	"time"
)

type Relation struct {
	Id       int64     `db:"id"`
	UserId   int64     `db:"user_id"`
	ToUserId int64     `db:"to_user_id" gorm:"to_user_id"`
	CreateAt time.Time `db:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `db:"update_at" gorm:"autoCreateTime"`
}
