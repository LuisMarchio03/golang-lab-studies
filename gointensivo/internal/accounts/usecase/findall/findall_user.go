package usecase

import (
	"github/LuisMarchio03/gointensivo/internal/accounts/infra"
	"github/LuisMarchio03/gointensivo/internal/accounts/infra/database"
)

type OutputFindUsersDTO struct {
	ID       string
	Name     string
	Age      int
	Cpf      int
	Birth    string
	Email    string
	Gender   string
	Password string
	Street   string
	City     string
}

type FindAllUsersUsecase struct {
	UserRepository infra.UserRepositoryInterface
}

func NewFindAllUsersUsecase(
	userRepository database.UserRepository,
) *FindAllUsersUsecase {
	return &FindAllUsersUsecase{
		UserRepository: &userRepository,
	}
}

func (f *FindAllUsersUsecase) Execute() ([]*OutputFindUsersDTO, error) {
	users, err := f.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*OutputFindUsersDTO
	for _, user := range users {
		userDTO := &OutputFindUsersDTO{
			ID:       user.ID,
			Name:     user.Name,
			Age:      user.Age,
			Cpf:      user.Cpf,
			Birth:    user.Birth,
			Email:    user.Email,
			Gender:   user.Gender,
			Password: user.Password,
			Street:   user.Street,
			City:     user.City,
		}

		output = append(output, userDTO)
	}

	return output, nil
}
