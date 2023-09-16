package mysql

import (
	"database/sql"
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
	VideoFindByIdList(list string) ([]model.Video, error)

	FavoriteInsert(data *model.Favorite) (sql.Result, error)
	FavoriteFindOne(id int64) (*model.Favorite, error)
	FavoriteFindOneByUserVideo(userId int64, videoId int64) (*model.Favorite, error)
	FavoriteFindByUser(userId int64) ([]model.Favorite, error)
	FavoriteUpdate(data *model.Favorite) error
	FavoriteDelete(id int64) error
	FavoriteAction(username, password string, videoId, actionType int64) error

	CommentInsert(data *model.Comment) (sql.Result, error)
	CommentFindOne(id int64) (*model.Comment, error)
	CommentFindByVideo(videoId int64) ([]model.Comment, error)
	CommentFindByUserVideo(userId int64, videoId int64) (*model.Comment, error)
	CommentUpdate(data *model.Comment) error
	CommentDelete(id int64) error
	VideoFindOneByVideo(id int64) (*model.Video, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Insert(data *model.User) (*model.User, error) {
	//TODO implement me
	err := u.db.Create(data).Error
	return data, err
}

func (u *UserRepository) FindOne(id int64) (*model.User, error) {
	//TODO implement me
	var user model.User
	err := u.db.Where("id=?", id).Find(&user).Error
	return &user, err
}

func (u *UserRepository) FindOneByToken(name string, password string) (*model.User, error) {
	//TODO implement me
	var user model.User
	err := u.db.Scopes().Where("name=? and password=?", name, password).Find(&user).Error
	return &user, err
}

func (u *UserRepository) Update(data *model.User) error {
	//TODO implement me
	err := u.db.Save(data).Error
	return err
}

func (u *UserRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func NewUUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
