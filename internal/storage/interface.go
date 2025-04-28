package storage

import (
	"github.com/dollarkillerx/xauth_backend/pkg/models"
	"gorm.io/gorm"
)

type Interface interface {
	DB() *gorm.DB

	CreateTodo(todo *models.Todo) error                    // 创建 todo
	UpdateTodoByID(id string, todo *models.Todo) error     // 更新 todo
	GetTodosByUserId(userId string) ([]models.Todo, error) // 获取 todo

	RegisterUser(email string, password string) error              // 注册用户
	LoginUser(email string, password string) (*models.User, error) // 登录

}
