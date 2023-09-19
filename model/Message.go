package model

import "time"

type Message struct {
	Id         int64 `gorm:"primaryKey autoIncrement"`
	ToUserId   int64
	FromUserId int64
	Content    string
	CreateAt   time.Time `gorm:"autoCreateTime"`
}

type MessageResp struct {
	Id         int64  `json:"id,omitempty" `
	ToUserId   int64  `json:"to_user_id" `
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content,omitempty" `
	CreateTime int64  `json:"create_time,omitempty"`
}
