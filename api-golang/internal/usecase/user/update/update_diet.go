package usecase

import (
	"github.com/LuisMarchio03/nutri/internal/entity"
	"github.com/LuisMarchio03/nutri/internal/infra"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
)

type UpdateDietInput struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Height   float64 `json:"height"`
	Weight   float64 `json:"weight"`
	Age      int     `json:"age"`
	Gender   string  `json:"gender"`
	Goal     int     `json:"goal"`
}

type UpdateDietUsecase struct {
	UserRepository infra.UserRepositoryInterface
}

func NewUpdateDietUsecase(
	userRepository repository.UserRepository,
	// orderRepository infra.OrderRepositoryInterface,
) *UpdateDietUsecase {
	return &UpdateDietUsecase{
		UserRepository: &userRepository,
	}
}

func (u *UpdateDietUsecase) Execute(
	ID string,
	input UpdateDietInput,
) (entity.User, error) {
	user, err := entity.NewUser(
		input.Name,
		input.Email,
		input.Password,
		input.Height,
		input.Weight,
		input.Age,
		input.Gender,
		input.Goal,
	)
	if err != nil {
		return entity.User{}, err
	}
	user.CalculateTotalCalories()
	user.CalculateTotalProteinas()

	err = u.UserRepository.Update(user, ID)
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
