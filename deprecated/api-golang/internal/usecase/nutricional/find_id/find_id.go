package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
)

type FindIdNutricionalUsecase struct {
	NutricionalRepository infra.NutricionalRepositoryInterface
}

func NewFindIdNutricionalUsecase(
	nutricionalRepository repository.NutricionalRepository,
) *FindIdNutricionalUsecase {
	return &FindIdNutricionalUsecase{
		NutricionalRepository: &nutricionalRepository,
	}
}

func (f *FindIdNutricionalUsecase) Execute(ID string) (*entity.Nutricional, error) {
	output, err := f.NutricionalRepository.FindUnique(ID)
	if err != nil {
		return &entity.Nutricional{}, err
	}
	return output, err
}
