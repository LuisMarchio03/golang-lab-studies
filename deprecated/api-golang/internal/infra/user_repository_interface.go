package infra

import "github.com/LuisMarchio03/nutri/internal/entity"

type UserRepositoryInterface interface {
	Save(user *entity.User) error
	FindAll() ([]*entity.User, error)
	FindInfosUser(userID string) (entity.User, error)
	Update(user *entity.User, userID string) error
}
