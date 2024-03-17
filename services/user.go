package services

import (
	"context"

	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"github.com/marifsulaksono/eticketing-test-mkp/repositories"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo repositories.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
}

func NewUserService(r repositories.UserRepository) *userService {
	return &userService{Repo: r}
}

func (u *userService) CreateUser(ctx context.Context, user *entities.User) error {
	// encrypt password with defaul cost (10)
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPwd)

	if user.Role == "" {
		user.Role = "customer"
	}

	return u.Repo.CreateUser(ctx, user)
}

func (u *userService) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	return u.Repo.GetUserByUsername(ctx, username)
}
