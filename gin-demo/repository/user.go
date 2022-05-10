package repository

import (
	"gotrunks/gin-demo/model"

	"github.com/gin-gonic/gin"
)

//userRepository the struct of userRepository
type userRepository struct {
	baseRepository
}

// NewUserRepository create an instance of *userRepository
func NewUserRepository(ctx *gin.Context) *userRepository {
	userRepository := userRepository{}
	userRepository.New(ctx)
	return &userRepository
}

// 将用户数据创建到数据库中
func (ur *userRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

//TODO 从数据库中查找用户记录
func (ur *userRepository) Find(ctx *gin.Context, user *model.User) error {
	username := ctx.DefaultPostForm("username", "aa")
	return ur.db.Where("username=?", username).Find(user).Error
}
