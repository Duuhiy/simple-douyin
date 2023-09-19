package mysql

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
)

func (u *UserRepository) MessageInsert(data *model.Message) error {
	//TODO implement me
	err := u.db.Create(data).Error
	return err
}

func (u *UserRepository) MessageFindOne(id int64) (*model.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) MessageFindByUserToUser(userId int64, toUserId int64) ([]model.Message, error) {
	//TODO implement me
	var msgs []model.Message
	q := fmt.Sprintf("select * from %s where (to_user_id=? and from_user_id=?) or (to_user_id=? and from_user_id=?)", "message")
	err := u.db.Raw(q, toUserId, userId, userId, toUserId).Scan(&msgs).Error
	//fmt.Println("MessageFindByUserToUser", msgs)
	return msgs, err
}

func (u *UserRepository) MessageUpdate(data *model.Message) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) MessageDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}
