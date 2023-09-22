package mysql

import (
	"database/sql"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Insert(data *model.User) (*model.User, error)
	FindOne(id int64) (*model.User, error)
	FindOneByName(name string) (string, int64, error)
	FindOneByToken(name string) (int64, error)
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
	VideoFindOneByVideo(id int64) (*model.Video, error)

	CommentInsert(data *model.Comment) error
	CommentFindOne(id int64) (*model.Comment, error)
	CommentFindByVideo(videoId int64) ([]model.Comment, error)
	CommentFindByUserVideo(userId int64, videoId int64) (*model.Comment, error)
	CommentUpdate(data *model.Comment) error
	CommentDelete(id int64) error
	CommentAdd(data *model.Comment) error
	CommentRemove(id, videoId int64) error

	RelationInsert(data *model.Relation) (*model.Relation, error)
	RelationFindOne(id int64) (*model.Relation, error)
	RelationFindOneByUserToUser(userId int64, toUserId int64) (*model.Relation, error)
	RelationFindByUser(userId int64) ([]model.Relation, error)
	RelationFindByToUser(toUserId int64) ([]model.Relation, error)
	RelationUpdate(data *model.Relation) error
	RelationDelete(id int64) error
	RelationDeleteByUser(userId int64, toUserId int64) error
	RelationAdd(user *model.User, user2 *model.User, follow *model.Relation) error
	RelationRemove(user *model.User, user2 *model.User) error
	UserFindByIdList(idList string) ([]model.User, error)

	MessageInsert(data *model.Message) error
	MessageFindOne(id int64) (*model.Message, error)
	MessageFindByUserToUser(userId int64, toUserId int64) ([]model.Message, error)
	MessageUpdate(data *model.Message) error
	MessageDelete(id int64) error
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) UserFindByIdList(idList string) ([]model.User, error) {
	//TODO implement me
	var users []model.User
	q := fmt.Sprintf("select * from %s where id in %s", "user", idList)
	//fmt.Println(q)
	err := u.db.Raw(q).Scan(&users).Error
	//err := u.db.Where("id in ?", idList).Find(&users).Error
	//fmt.Println(users)
	return users, err
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

func (u *UserRepository) FindOneByName(name string) (string, int64, error) {
	//TODO implement me
	var user model.User
	fmt.Println("FindOneByToken", name)
	q := fmt.Sprintf("select id, password from %s where name = ?", "user")
	err := u.db.Raw(q, name).Scan(&user).Error
	fmt.Println("FindOneByToken", user.Id)
	//log.Println(err)
	return user.Password, user.Id, err
}

func (u *UserRepository) FindOneByToken(name string) (int64, error) {
	//TODO implement me
	var user model.User
	fmt.Println("FindOneByToken", name)
	q := fmt.Sprintf("select id from %s where name = ?", "user")
	err := u.db.Raw(q, name).Scan(&user).Error
	fmt.Println("FindOneByToken", user.Id)
	//log.Println(err)
	return user.Id, err
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
