package mysql

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(data *model.User) (*model.User, error)
	FindOne(id int64) (*model.User, error)
	FindOneByToken(name string, password string) (*model.User, error)
	Update(data *model.User) error
	Delete(id int64) error

	VideoInsert(data *model.Video) error
	VideoFindOne(id int64) (*model.Video, error)
	VideoFindOneByUser(name string, password string) (*model.Video, error)
	VideoFindAll() ([]model.Video, error)
	VideoUpdate(data *model.Video) error
	VideoDelete(id int64) error
	Publish(video *model.Video, author *model.User, data []byte) error
	FindAllByAuthor(userId int64) ([]model.Video, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u UserRepository) Insert(data *model.User) (*model.User, error) {
	//TODO implement me
	err := u.db.Create(data).Error
	return data, err
}

func (u UserRepository) FindOne(id int64) (*model.User, error) {
	//TODO implement me
	var user model.User
	err := u.db.Where("id=?", id).Find(&user).Error
	return &user, err
}

func (u UserRepository) FindOneByToken(name string, password string) (*model.User, error) {
	//TODO implement me
	var user model.User
	err := u.db.Scopes().Where("name=? and password=?", name, password).Find(&user).Error
	return &user, err
}

func (u UserRepository) Update(data *model.User) error {
	//TODO implement me
	err := u.db.Save(data).Error
	return err
}

func (u UserRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewUUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
