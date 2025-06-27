package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
)

type FindAllNutricionalUsecase struct {
	NutricionalRepository infra.NutricionalRepositoryInterface
}

func NewFindAllNutricionalUsecase(nutricionalRepository repository.NutricionalRepository) *FindAllNutricionalUsecase {
	return &FindAllNutricionalUsecase{
		NutricionalRepository: &nutricionalRepository,
	}
}

func (f *FindAllNutricionalUsecase) Execute() ([]*entity.Nutricional, error) {
	output, err := f.NutricionalRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return output, nil
}
