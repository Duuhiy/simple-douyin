package model

import (
	"time"
)

type User struct {
	Id              int64     `db:"id" gorm:"primaryKey autoIncrement" redis:"id"`
	Name            string    `db:"name" redis:"name"`                                      // The username
	Password        string    `db:"password"`                                               // The password
	FollowCount     int64     `db:"follow_count" gorm:"default:0" redis:"follow_count"`     // The follow_count
	FollowerCount   int64     `db:"follower_count" gorm:"default:0" redis:"follower_count"` // The follower_count
	Avatar          string    `db:"avatar" redis:"avatar"`
	BackgroundImage string    `db:"background_image" redis:"background_image"`
	Signature       string    `db:"signature" redis:"signature"`
	TotalFavorited  int64     `db:"total_favorited" gorm:"default:0" redis:"total_favorited"` // The total_favorited
	WorkCount       int64     `db:"work_count" gorm:"default:0" redis:"work_count"`           // The work_count
	FavoriteCount   int64     `db:"favorite_count" gorm:"default:0" redis:"favorite_count"`   // The favorite_count
	CreateAt        time.Time `db:"create_at" gorm:"autoCreateTime"`
	UpdateAt        time.Time `db:"update_at" gorm:"autoCreateTime"`
}

type UserResp struct {
	Id              int64  `json:"id" db:"id" gorm:"primaryKey autoIncrement"`
	Name            string `json:"name" db:"name"`                                      // The username
	FollowCount     int64  `json:"follow_count" db:"follow_count" gorm:"default:0"`     // The follow_count
	FollowerCount   int64  `json:"follower_count" db:"follower_count" gorm:"default:0"` // The follower_count
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar" db:"avatar"`
	BackgroundImage string `json:"background_image" db:"background_image"`
	Signature       string `json:"signature" db:"signature"`
	TotalFavorited  int64  `json:"total_favorited" db:"total_favorited" gorm:"default:0"` // The total_favorited
	WorkCount       int64  `json:"work_count" db:"work_count" gorm:"default:0"`           // The work_count
	FavoriteCount   int64  `json:"favorite_count" db:"favorite_count" gorm:"default:0"`   // The favorite_count
}
