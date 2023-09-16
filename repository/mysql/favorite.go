package mysql

import (
	"database/sql"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"log"
)

func (u *UserRepository) FavoriteInsert(data *model.Favorite) (sql.Result, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FavoriteFindOne(id int64) (*model.Favorite, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FavoriteFindOneByUserVideo(userId int64, videoId int64) (*model.Favorite, error) {
	//TODO implement me
	var favorite model.Favorite
	err := u.db.Where("user_id = ? and video_id = ?", userId, videoId).Find(&favorite).Error
	return &favorite, err
}

func (u *UserRepository) FavoriteFindByUser(userId int64) ([]model.Favorite, error) {
	//TODO implement me
	var favorite []model.Favorite
	err := u.db.Where("user_id = ?", userId).Find(&favorite).Error
	fmt.Println(favorite)
	return favorite, err
}

func (u *UserRepository) FavoriteUpdate(data *model.Favorite) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FavoriteDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) FavoriteAction(username, password string, videoId, actionType int64) error {
	user, _ := u.FindOneByToken(username, password)
	video, _ := u.VideoFindOne(videoId)
	var favorite model.Favorite
	switch actionType {
	case 1:
		// 点赞
		// 1.判断是否点过赞了
		res, err := u.FavoriteFindOneByUserVideo(user.Id, videoId)
		log.Println(favorite, err)
		if *res != favorite {
			// 点过赞了
			return nil
		} else {
			tx := u.db.Begin()
			// 1.视频点赞+1
			video.FavoriteCount++
			if err = tx.Save(&video).Error; err != nil {
				log.Println("视频点赞+1 出错")
				tx.Rollback()
			}
			// 2.用户点赞+1
			user.FavoriteCount++
			if err = tx.Save(&user).Error; err != nil {
				log.Println("用户点赞+1 出错")
				tx.Rollback()
			}
			// 3.加入favorite表中
			favorite := model.Favorite{
				UserId:  user.Id,
				VideoId: videoId,
			}
			if err = tx.Create(&favorite).Error; err != nil {
				log.Println("加入favorite表中 出错")
				tx.Rollback()
			}
			return tx.Commit().Error
		}
	case 2:
		// 取消赞
		// 1.判断是否点过赞了
		res, err := u.FavoriteFindOneByUserVideo(user.Id, videoId)
		if *res == favorite {
			// 没点过赞
			return nil
		} else {
			tx := u.db.Begin()
			// 1.视频点赞-1
			video.FavoriteCount--
			if err = tx.Save(&video).Error; err != nil {
				log.Println("视频点赞-1 出错")
				tx.Rollback()
			}
			// 2.用户点赞-1
			user.FavoriteCount--
			if err = tx.Save(&user).Error; err != nil {
				log.Println("用户点赞-1 出错")
				tx.Rollback()
			}
			// 3.删除favorite表中
			var favorite model.Favorite
			if err = tx.Where("user_id=? and video_id=?", user.Id, videoId).Delete(&favorite).Error; err != nil {
				log.Println("删除favorite表中 出错")
				tx.Rollback()
			}
			return tx.Commit().Error
		}
	}
	return nil
}
