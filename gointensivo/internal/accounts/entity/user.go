package entity

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
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

func NewUser(name string, age int, cpf int, birth string, email string, gender string, password string, street string, city string) (*User, error) {
	user := &User{
		Name:     name,
		Age:      age,
		Cpf:      cpf,
		Birth:    birth,
		Email:    email,
		Gender:   gender,
		Password: password,
		Street:   street,
		City:     city,
	}
	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) IsValid() error {
	if u.Name == "" {
		return errors.New("name is empty")
	}
	if u.Age == 0 {
		return errors.New("age is empty")
	}
	if u.Cpf == 0 {
		return errors.New("cpf is empty")
	}
	if u.Birth == "" {
		return errors.New("birth is empty")
	}
	if u.Email == "" {
		return errors.New("email is empty")
	}
	if u.Gender == "" {
		return errors.New("gender is empty")
	}
	if u.Password == "" {
		return errors.New("password is empty")
	}
	if u.Street == "" {
		return errors.New("street is empty")
	}
	if u.City == "" {
		return errors.New("city is empty")
	}
	return nil
}

func (u *User) SetID() (string, error) {
	u.ID = uuid.New().String()
	return u.ID, nil
}
