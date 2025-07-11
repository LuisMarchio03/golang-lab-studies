package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
)

type UpdateNutricionalInputDTO struct {
	ID           string
	FoodName     string  `json:"food_name"`
	Quantity     int     `json:"quantity"`
	CaloriesUnit float64 `json:"calories_unit"`
}

type UpateNutricionalUsecase struct {
	NutricionalRepository infra.NutricionalRepositoryInterface
}

func NewUpateNutricionalUsecase(
	nutricionalRepository repository.NutricionalRepository,
) *UpateNutricionalUsecase {
	return &UpateNutricionalUsecase{
		NutricionalRepository: &nutricionalRepository,
	}
}

func (u *UpateNutricionalUsecase) Execute(input UpdateNutricionalInputDTO) error {
	nutricional, err := entity.NewNutricional(input.FoodName, input.Quantity, input.CaloriesUnit)
	if err != nil {
		return err
	}
	nutricional.CalculeteTotalCalories()
	err = u.NutricionalRepository.Update(nutricional)
	if err != nil {
		return err
	}
	return nil
}
