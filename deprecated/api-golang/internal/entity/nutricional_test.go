package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyFoodName_WhenCreateANewNutricional_ThenShouldReceiveAnError(t *testing.T) {
	n := Nutricional{ID: "1"}
	assert.Error(t, n.IsValidNutricional(), "invalid food name")
}

func TestGivenAnEmptyQuantity_WhenCreateANewNutricional_ThenShouldReceiveAnError(t *testing.T) {
	n := Nutricional{ID: "1", FoodName: "Food Name"}
	assert.Error(t, n.IsValidNutricional(), "invalid quantity")
}
