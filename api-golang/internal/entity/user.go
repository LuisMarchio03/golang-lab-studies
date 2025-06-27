package entity

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	Height         float64 `json:"height"` // altura
	Weight         float64 `json:"weight"` // peso
	Age            int     `json:"age"`
	Gender         string  `json:"gender"`
	Goal           int     `json:"goal"` // Meta
	TotalCalories  float64 `json:"total_calories"`
	TotalProteinas float64 `json:"total_proteinas"`
	// 1 - Perder peso
	// 2 - Ganho de massa muscular
}

func NewUser(
	name string,
	email string,
	password string,
	height float64,
	weight float64,
	age int,
	gender string,
	goal int,
) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
		Height:   height,
		Weight:   weight,
		Age:      age,
		Gender:   gender,
		Goal:     goal,
	}
	err := user.IsValid()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) SetID() error {
	u.ID = uuid.New().String()
	err := u.IsValid()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) CalculateTotalCalories() error {
	if u.Goal == 1 {
		u.TotalCalories = u.Weight * 25
	}
	if u.Goal == 2 {
		u.TotalCalories = u.Weight * 35
	}
	return nil
}

func (u *User) CalculateTotalProteinas() error {
	if u.Goal == 1 {
		u.TotalProteinas = u.Weight * 2.2
	}
	if u.Goal == 2 {
		u.TotalProteinas = u.Weight * 2.7
	}
	return nil
}

func (u *User) IsValid() error {
	// if u.ID == "" {
	// 	return errors.New("invalid id")
	// }
	if u.Name == "" {
		return errors.New("invalid name")
	}
	if u.Email == "" {
		return errors.New("invalid email")
	}
	if u.Password == "" {
		return errors.New("invalid password")
	}
	if u.Height == 0 {
		return errors.New("invalid height")
	}
	if u.Weight == 0 {
		return errors.New("invalid weight")
	}
	if u.Age == 0 {
		return errors.New("invalid age")
	}
	if u.Gender == "" {
		return errors.New("invalid gender")
	}
	if u.Goal == 0 || u.Goal > 2 || u.Goal < 1 {
		return errors.New("invalid goal")
	}
	return nil
}
