package model

import (
	"time"
)

type User struct {
	Id              int64     `db:"id" gorm:"primaryKey autoIncrement"`
	Name            string    `db:"name"`                            // The username
	Password        string    `db:"password"`                        // The password
	FollowCount     int64     `db:"follow_count" gorm:"default:0"`   // The follow_count
	FollowerCount   int64     `db:"follower_count" gorm:"default:0"` // The follower_count
	Avatar          string    `db:"avatar"`
	BackgroundImage string    `db:"background_image"`
	Signature       string    `db:"signature"`
	TotalFavorited  int64     `db:"total_favorited" gorm:"default:0"` // The total_favorited
	WorkCount       int64     `db:"work_count" gorm:"default:0"`      // The work_count
	FavoriteCount   int64     `db:"favorite_count" gorm:"default:0"`  // The favorite_count
	CreateAt        time.Time `db:"create_at" gorm:"autoCreateTime"`
	UpdateAt        time.Time `db:"update_at" gorm:"autoCreateTime"`
}
