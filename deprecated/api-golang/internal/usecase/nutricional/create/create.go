package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
)

type CreateNutricionalInput struct {
	FoodName     string  `json:"food_name"`
	Quantity     int     `json:"quantity"`
	CaloriesUnit float64 `json:"calories_unit"`
}

type CreateNutricionalUsecase struct {
	NutricionalRepository infra.NutricionalRepositoryInterface
}

func NewCreateNutricionalUsecase(
	nutricionalRepository repository.NutricionalRepository,
) *CreateNutricionalUsecase {
	return &CreateNutricionalUsecase{
		NutricionalRepository: &nutricionalRepository,
	}
}

func (createNutricionalUsecase *CreateNutricionalUsecase) Execute(createNutricionalInput CreateNutricionalInput) error {
	nutricional, err := entity.NewNutricional(
		createNutricionalInput.FoodName,
		createNutricionalInput.Quantity,
		createNutricionalInput.CaloriesUnit,
	)
	if err != nil {
		return err
	}
	nutricional.GenerateIDNutricional()
	nutricional.CalculeteTotalCalories()

	err = createNutricionalUsecase.NutricionalRepository.Save(nutricional)
	if err != nil {
		return err
	}
	return nil
}
