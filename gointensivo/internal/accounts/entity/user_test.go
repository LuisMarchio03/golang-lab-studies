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
	user := User{ID: "123"}
	assert.Error(t, user.IsValid(), "invalid name")
}

func TestGivenAnEmptyAge_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John"}
	assert.Error(t, user.IsValid(), "invalid age")
}

func TestGivenAnEmptyCpf_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10}
	assert.Error(t, user.IsValid(), "invalid cpf")
}

func TestGivenAnEmptyBirth_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123}
	assert.Error(t, user.IsValid(), "invalid birth")
}

func TestGivenAnEmptyEmail_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000"}
	assert.Error(t, user.IsValid(), "invalid email")
}

func TestGivenAnEmptyGender_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000", Email: "lg@gmail.com"}
	assert.Error(t, user.IsValid(), "invalid gender")
}

func TestGivenAnEmptyPassword_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000", Email: "lg@gmail.com", Gender: "M"}
	assert.Error(t, user.IsValid(), "invalid password")
}

func TestGivenAnEmptyStreet_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000", Email: "lg@gmail.com", Gender: "M", Password: "123"}
	assert.Error(t, user.IsValid(), "invalid street")
}

func TestGivenAnEmptyCity_WhenCreateANewUser_ThenShouldReceiveAnError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000", Email: "lg@gmail.com", Gender: "M", Password: "123", Street: "Rua 1"}
	assert.Error(t, user.IsValid(), "invalid street")
}

func TestGivenAValidUser_WhenCreateANewUser_ThenShouldReceiveNoError(t *testing.T) {
	user := User{ID: "123", Name: "John", Age: 10, Cpf: 123, Birth: "01/01/2000", Email: "lg@gmail.com", Gender: "M", Password: "123", Street: "Rua 1", City: "São Paulo"}
	assert.Equal(t, "123", user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, 10, user.Age)
	assert.Equal(t, 123, user.Cpf)
	assert.Equal(t, "01/01/2000", user.Birth)
	assert.Equal(t, "lg@gmail.com", user.Email)
	assert.Equal(t, "M", user.Gender)
	assert.Equal(t, "123", user.Password)
	assert.Equal(t, "Rua 1", user.Street)
	assert.Equal(t, "São Paulo", user.City)
	assert.Nil(t, user.IsValid())
}
