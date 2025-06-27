package infra

import "github/LuisMarchio03/gointensivo/internal/accounts/entity"

type UserRepositoryInterface interface {
	Save(user *entity.User) error
	FindAll() ([]*entity.User, error)
}
