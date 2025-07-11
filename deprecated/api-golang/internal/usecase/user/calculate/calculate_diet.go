package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
)

type CalculateDietInput struct {
	Name     string
	Email    string
	Password string
	Height   float64
	Weight   float64
	Age      int
	Gender   string
	Goal     int
}

type CalculateDietUsecase struct {
	UserRepository infra.UserRepositoryInterface
}

func NewCalculateDietUsecase(
	userRepository repository.UserRepository,
	// orderRepository infra.OrderRepositoryInterface,
) *CalculateDietUsecase {
	return &CalculateDietUsecase{
		UserRepository: &userRepository,
	}
}

func (calculateDietUsecase *CalculateDietUsecase) Execute(calculateDietInput CalculateDietInput) (entity.User, error) {
	user, err := entity.NewUser(
		calculateDietInput.Name,
		calculateDietInput.Email,
		calculateDietInput.Password,
		calculateDietInput.Height,
		calculateDietInput.Weight,
		calculateDietInput.Age,
		calculateDietInput.Gender,
		calculateDietInput.Goal,
	)
	if err != nil {
		return entity.User{}, err
	}
	user.SetID()
	user.CalculateTotalCalories()
	user.CalculateTotalProteinas()

	err = calculateDietUsecase.UserRepository.Save(user)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Password:       user.Password,
		Height:         user.Height,
		Weight:         user.Weight,
		Age:            user.Age,
		Gender:         user.Gender,
		Goal:           user.Goal,
		TotalCalories:  user.TotalCalories,
		TotalProteinas: user.TotalProteinas,
	}, nil
}
