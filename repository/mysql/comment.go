package mysql

import (
	"database/sql"
	"github.com/RaymondCode/simple-demo/model"
)

func (u *UserRepository) CommentInsert(data *model.Comment) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentFindOne(id int64) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentFindByVideo(videoId int64) ([]model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentFindByUserVideo(userId int64, videoId int64) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentUpdate(data *model.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}
