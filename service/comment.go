package service

import (
	"github.com/RaymondCode/simple-demo/repository/mysql"
)

type ICommentService interface {
}

type CommentService struct {
	CommentRepository mysql.IUserRepository
}

func NewCommentService(repository mysql.IUserRepository) ICommentService {
	return &CommentService{CommentRepository: repository}
}
