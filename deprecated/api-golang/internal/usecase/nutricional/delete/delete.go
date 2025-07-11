package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
)

type DeleleNutricionalUsecase struct {
	NutricionalRepository infra.NutricionalRepositoryInterface
}

func NewDeleleNutricionalUsecase(
	nutricionalRepository repository.NutricionalRepository,
) *DeleleNutricionalUsecase {
	return &DeleleNutricionalUsecase{
		NutricionalRepository: &nutricionalRepository,
	}
}

func (d *DeleleNutricionalUsecase) Execute(ID string) error {
	err := d.NutricionalRepository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
