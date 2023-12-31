package mysql

import (
	"bytes"
	"github.com/RaymondCode/simple-demo/model"
	oss "github.com/RaymondCode/simple-demo/utils"
	"log"
)

func (u *UserRepository) Publish(video *model.Video, author *model.User, data []byte) error {
	//TODO implement me
	tx := u.db.Begin()
	uploadPath := video.Title + ".mp4"
	video.PlayUrl = "https://douyin-duu.oss-cn-beijing.aliyuncs.com/" + uploadPath
	video.CoverUrl = "https://douyin-duu.oss-cn-beijing.aliyuncs.com/" + uploadPath + "?x-oss-process=video/snapshot,t_0,f_jpg,w_800,h_600"
	// 插入video
	if err := tx.Create(video).Error; err != nil {
		log.Println("插入video表出错了", err)
		tx.Rollback()
		return err
	}
	// 修改作者的work_count
	if err := tx.Save(author).Error; err != nil {
		log.Println("修改作者的work_count出错了", err)
		tx.Rollback()
		return err
	}
	// 上传到oss
	if err := oss.Bucket.PutObject(uploadPath, bytes.NewReader(data)); err != nil {
		log.Println("上传到oss出错了", err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *UserRepository) VideoInsert(data *model.Video) error {
	//TODO implement me
	err := u.db.Create(data).Error
	return err
}

func (u *UserRepository) VideoFindOne(id int64) (*model.Video, error) {
	//TODO implement me
	var video model.Video
	err := u.db.Where("id = ?", id).Find(&video).Error
	return &video, err
}

func (u *UserRepository) VideoFindOneByUser(name string, password string) (*model.Video, error) {
	//TODO implement me
	var video model.Video
	err := u.db.Where("name = ? and password = ?", name, password).Find(&video).Error
	return &video, err
}

func (u *UserRepository) VideoFindAll() ([]model.Video, error) {
	//TODO implement me
	var videos []model.Video
	err := u.db.Order("create_at desc").Find(&videos).Error
	return videos, err
}

func (u *UserRepository) FindAllByAuthor(userId int64) ([]model.Video, error) {
	//TODO implement me
	var videos []model.Video
	//log.Println(u.db.Where("author=?", userId).Order("create_at desc").Find(&videos).Debug())
	err := u.db.Order("create_at desc").Where("author=?", userId).Find(&videos).Error
	//fmt.Println(videos)
	return videos, err
}

func (u *UserRepository) VideoFindByIdList(list []int64) ([]model.Video, error) {
	//TODO implement me
	var videos []model.Video
	err := u.db.Find(&videos, list).Error
	return videos, err
}

func (u *UserRepository) VideoFindOneByVideo(id int64) (*model.Video, error) {
	//TODO implement me
	var video model.Video
	err := u.db.Where("id = ?", id).Order("create_at desc").Find(&video).Error
	return &video, err
}

func (u *UserRepository) VideoUpdate(data *model.Video) error {
	//TODO implement me
	err := u.db.Save(data).Error
	return err
}

func (u *UserRepository) VideoDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}
