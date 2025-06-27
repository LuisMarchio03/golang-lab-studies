package usecase

import (
	"github/LuisMarchio03/gointensivo/internal/accounts/entity"
	"github/LuisMarchio03/gointensivo/internal/accounts/infra"
	"github/LuisMarchio03/gointensivo/internal/accounts/infra/database"
)

type InputSaveUserDTO struct {
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

// type OutputSaveUserDTO struct {
// 	ID       string
// 	Name     string
// 	Age      int
// 	Cpf      int
// 	Birth    string
// 	Email    string
// 	Gender   string
// 	Password string
// 	Street   string
// 	City     string
// }

type SaveUserUsecase struct {
	UserRepository infra.UserRepositoryInterface
}

func NewSaveUserUsecase(
	userRepository database.UserRepository,
) *SaveUserUsecase {
	return &SaveUserUsecase{
		UserRepository: &userRepository,
	}
}

func (s *SaveUserUsecase) Execute(
	input InputSaveUserDTO,
) error {
	user, err := entity.NewUser(input.Name, input.Age, input.Cpf, input.Birth, input.Email, input.Gender, input.Password, input.Street, input.City)
	if err != nil {
		return err
	}
	_, err = user.SetID()
	if err != nil {
		return err
	}
	err = s.UserRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}
