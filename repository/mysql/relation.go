package mysql

import (
	"github.com/RaymondCode/simple-demo/model"
	"log"
)

func (u *UserRepository) RelationInsert(data *model.Relation) (*model.Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationFindOne(id int64) (*model.Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationFindOneByUserToUser(userId int64, toUserId int64) (*model.Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationFindByUser(userId int64) ([]model.Relation, error) {
	//TODO implement me
	var userList []model.Relation
	err := u.db.Where("user_id=?", userId).Find(&userList).Error
	return userList, err
}

func (u *UserRepository) RelationFindByToUser(toUserId int64) ([]model.Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationUpdate(data *model.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationDeleteByUser(userId int64, toUserId int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) RelationAdd(user *model.User, user2 *model.User, follow *model.Relation) error {
	// 1.用户follow_count++
	// 2.toUserId follower_count++
	// 3.插入relation表中
	var relation, none model.Relation
	// 已经关注过了，不能再关注
	u.db.Where("user_id = ? and to_user_id = ?", user.Id, user2.Id).Find(&relation)
	if relation != none {
		return nil
	}
	tx := u.db.Begin()
	user.FollowCount++
	user2.FollowerCount++
	if err := tx.Save(user).Error; err != nil {
		log.Println("用户follow_count++ 出错了")
		tx.Rollback()
		return err
	}
	if err := tx.Save(user2).Error; err != nil {
		log.Println("toUserId follower_count++ 出错了")
		tx.Rollback()
		return err
	}
	if err := tx.Create(follow).Error; err != nil {
		log.Println("插入relation表中出错了")
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (u *UserRepository) RelationRemove(user *model.User, user2 *model.User) error {
	//TODO implement me
	// 1.用户follow_count--
	// 2.toUserId follower_count--
	// 3.从relation表中删除
	user.FollowCount--
	user2.FollowerCount--
	tx := u.db.Begin()
	if err := tx.Save(user).Error; err != nil {
		log.Println("用户follow_count-- 出错了")
		tx.Rollback()
		return err
	}
	if err := tx.Save(user2).Error; err != nil {
		log.Println("toUserId follower_count-- 出错了")
		tx.Rollback()
		return err
	}
	var relation model.Relation
	if err := tx.Where("user_id=? and to_user_id=?", user.Id, user2.Id).Delete(&relation).Error; err != nil {
		log.Println("从relation表中删除出错了")
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
