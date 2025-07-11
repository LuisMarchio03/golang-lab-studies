package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGivenAnEmptyID_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
// 	user := User{}
// 	assert.Error(t, user.IsValid(), "invalid id")
// }

func TestGivenAnEmptyName_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1"}
	assert.Error(t, user.IsValid(), "invalid name")
}

func TestGivenAnEmptyEmail_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test"}
	assert.Error(t, user.IsValid(), "invalid email")
}

func TestGivenAnEmptyPassword_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com"}
	assert.Error(t, user.IsValid(), "invalid password")
}

func TestGivenAnEmptyHeight_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456"}
	assert.Error(t, user.IsValid(), "invalid height")
}

func TestGivenAnEmptyWeight_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80}
	assert.Error(t, user.IsValid(), "invalid weight")
}

func TestGivenAnEmptyAge_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80}
	assert.Error(t, user.IsValid(), "invalid age")
}

func TestGivenAnEmptyGender_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20}
	assert.Error(t, user.IsValid(), "invalid gender")
}

func TestGivenAnEmptyGoal_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M"}
	assert.Error(t, user.IsValid(), "invalid goal")
}

func TestGivenAnInvalidGoal_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 3}
	assert.Error(t, user.IsValid(), "invalid goal")
}

func TestGivenAValidUser_WhenCreateANewUser_ThenShouldReceiveNoError(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 2}
	assert.NoError(t, user.IsValid())
}

func TestGivenAValidUser_WhenCalculeteTotalCalorias_ThenShouldReceiveTheCorrectValue(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 1}
	user.CalculateTotalCalories()
	assert.Equal(t, 2000.0, user.TotalCalories)
}

func TestGivenAValidUser_WhenCalculeteTotalCalorias_ThenShouldReceiveTheCorrectValue2(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 2}
	user.CalculateTotalCalories()
	assert.Equal(t, 2800.0, user.TotalCalories)
}

func TestGivenAValidUser_WhenCalculeteTotalProteinas_ThenShouldReceiveTheCorrectValue(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 2}
	user.CalculateTotalProteinas()
	assert.Equal(t, 216.0, user.TotalProteinas)
}

func TestGivenAValidUser_WhenCalculeteTotalProteinas_ThenShouldReceiveTheCorrectValue2(t *testing.T) {
	user := User{ID: "1", Name: "Test", Email: "test@email.com", Password: "123456", Height: 1.80, Weight: 80, Age: 20, Gender: "M", Goal: 1}
	user.CalculateTotalProteinas()
	assert.Equal(t, 176.0, user.TotalProteinas)
}
