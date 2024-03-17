package repositories

import (
	"context"

	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) CreateUser(ctx context.Context, user *entities.User) error {
	return u.DB.Create(&user).Error
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
