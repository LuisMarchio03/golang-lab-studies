package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Nutricional struct {
	ID            string  `json:"id"`
	IDTaco        string  `json:"id_taco"`
	FoodName      string  `json:"food_name"`
	Quantity      int     `json:"quantity"`
	CaloriesUnit  float64 `json:"calories_unit"`
	TotalCalories float64 `json:"total_calories"`
}

func NewNutricional(
	foodName string,
	quantity int,
	caloriesUnit float64,
) (*Nutricional, error) {
	nutricional := &Nutricional{
		FoodName:     foodName,
		Quantity:     quantity,
		CaloriesUnit: caloriesUnit,
	}
	err := nutricional.IsValidNutricional()
	if err != nil {
		return nil, err
	}
	return nutricional, nil
}

func (n *Nutricional) GenerateIDNutricional() {
	n.ID = uuid.New().String()
}

func (n *Nutricional) CalculeteTotalCalories() {
	n.TotalCalories = float64(n.Quantity) * n.CaloriesUnit
}

func (n *Nutricional) IsValidNutricional() error {
	if n.FoodName == "" {
		return errors.New("invalid food name")
	}
	if n.Quantity == 0.0 {
		return errors.New("invalid quantity")
	}
	return nil
}
