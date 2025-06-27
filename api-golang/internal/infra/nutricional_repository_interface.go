package infra

import "github.com/LuisMarchio03/nutri/internal/entity"

type NutricionalRepositoryInterface interface {
	Save(n *entity.Nutricional) error
	FindAll() ([]*entity.Nutricional, error)
	FindUnique(ID string) (*entity.Nutricional, error)
	Update(n *entity.Nutricional) error
	Delete(ID string) error
}
