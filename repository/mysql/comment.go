package mysql

import (
	"github.com/RaymondCode/simple-demo/model"
	"log"
)

func (u *UserRepository) CommentInsert(data *model.Comment) error {
	//TODO implement me
	err := u.db.Create(data).Error
	return err
}

func (u *UserRepository) CommentAdd(data *model.Comment) error {
	//TODO implement me
	var video model.Video
	u.db.Where("id = ?", data.VideoId).Find(&video)
	tx := u.db.Begin()
	if err := tx.Create(data).Error; err != nil {
		log.Println("插入comment表出错了", err)
		tx.Rollback()
		return err
	}
	video.CommentCount++
	if err := tx.Save(&video).Error; err != nil {
		log.Println("更新video表出错了", err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *UserRepository) CommentFindOne(id int64) (*model.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CommentFindByVideo(videoId int64) ([]model.Comment, error) {
	//TODO implement me
	var commentList []model.Comment
	err := u.db.Where("video_id=?", videoId).Find(&commentList).Error
	return commentList, err
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
	var comment model.Comment
	err := u.db.Where("id=?", id).Delete(&comment).Error
	return err
}

func (u *UserRepository) CommentRemove(id, videoId int64) error {
	//TODO implement me
	tx := u.db.Begin()
	var comment model.Comment
	if err := tx.Where("id=?", id).Delete(&comment).Error; err != nil {
		tx.Rollback()
		log.Println("从comment表删除出错了", err)
		return err
	}
	var video model.Video
	//fmt.Println(comment)
	tx.Where("id=?", videoId).Find(&video)
	video.CommentCount--
	if err := tx.Save(&video).Error; err != nil {
		tx.Rollback()
		log.Println("更新video表出错了", err)
		return err
	}
	return tx.Commit().Error
}
