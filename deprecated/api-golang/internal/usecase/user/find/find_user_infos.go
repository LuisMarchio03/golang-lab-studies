package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
)

type FindUserInfosOutputDTO struct {
	ID             string
	Name           string
	Email          string
	Height         float64
	Weight         float64
	Age            int
	Gender         string
	Goal           int
	TotalCalories  float64
	TotalProteinas float64
}

type FindUserInfosUsecase struct {
	UserRepository infra.UserRepositoryInterface
}

func NewFindUserInfosUsecase(
	userRepository repository.UserRepository,
	// orderRepository infra.OrderRepositoryInterface,
) *FindUserInfosUsecase {
	return &FindUserInfosUsecase{
		UserRepository: &userRepository,
	}
}

func (findUserInfosUsecase *FindUserInfosUsecase) Execute(userID string) (FindUserInfosOutputDTO, error) {
	output, err := findUserInfosUsecase.UserRepository.FindInfosUser(userID)
	if err != nil {
		return FindUserInfosOutputDTO{}, err
	}
	return FindUserInfosOutputDTO{
		ID:             output.ID,
		Name:           output.Name,
		Email:          output.Email,
		Height:         output.Height,
		Weight:         output.Weight,
		Age:            output.Age,
		Gender:         output.Gender,
		Goal:           output.Goal,
		TotalCalories:  output.TotalCalories,
		TotalProteinas: output.TotalProteinas,
	}, nil
}
