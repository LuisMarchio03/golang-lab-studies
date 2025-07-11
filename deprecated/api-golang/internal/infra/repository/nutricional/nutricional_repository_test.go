package repository

import (
	"database/sql"
	"testing"

	"github.com/LuisMarchio03/nutri/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type NutricionalRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *NutricionalRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE nutricional (id VARCHAR(255) NOT NULL, id_taco VARCHAR(255) NOT NULL, food_name VARCHAR(255) NOT NULL, quantity FLOAT NOT NULL, calories_unit FLOAT NOT NULL, total_calories FLOAT NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *NutricionalRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(NutricionalRepositoryTestSuite))
}

func (suite *NutricionalRepositoryTestSuite) TestSave() {
	suite.Run("Should save a nutricional", func() {
		nutricional, err := entity.NewNutricional(
			"test",
			10,
			100.0,
		)
		suite.NoError(err)
		nutricional.GenerateIDNutricional()
		nutricional.IDTaco = "123"
		nutricional.FoodName = "Test"
		nutricional.Quantity = 1
		nutricional.CaloriesUnit = 100.0
		nutricional.CalculeteTotalCalories()

		repository := NewNutricionalRepository(suite.Db)
		err = repository.Save(nutricional)
		suite.NoError(err)
	})

	suite.Run("find all", func() {
		nutricional, err := entity.NewNutricional(
			"test 2",
			12,
			100.0,
		)
		suite.NoError(err)
		nutricional.GenerateIDNutricional()
		nutricional.IDTaco = "123"
		nutricional.FoodName = "Test"
		nutricional.Quantity = 1
		nutricional.CaloriesUnit = 100.0
		nutricional.CalculeteTotalCalories()

		repository := NewNutricionalRepository(suite.Db)
		err = repository.Save(nutricional)
		suite.NoError(err)

		nutricionals, err := repository.FindAll()
		suite.NoError(err)
		suite.Equal(2, len(nutricionals))
	})

	suite.Run("find unique", func() {
		nutricional, err := entity.NewNutricional(
			"test 3",
			13,
			100.0,
		)
		suite.NoError(err)
		nutricional.GenerateIDNutricional()
		nutricional.IDTaco = "123"
		nutricional.FoodName = "Test"
		nutricional.Quantity = 1
		nutricional.CaloriesUnit = 100.0
		nutricional.CalculeteTotalCalories()

		repository := NewNutricionalRepository(suite.Db)
		err = repository.Save(nutricional)
		suite.NoError(err)

		nutricional, err = repository.FindUnique(nutricional.ID)
		suite.NoError(err)
		suite.Equal("123", nutricional.IDTaco)
		suite.Equal("Test", nutricional.FoodName)
		suite.Equal(1, nutricional.Quantity)
		suite.Equal(100.0, nutricional.CaloriesUnit)
		suite.Equal(100.0, nutricional.TotalCalories)
	})

	suite.Run("update", func() {
		nutricional, err := entity.NewNutricional(
			"test 4",
			14,
			100.0,
		)
		suite.NoError(err)
		nutricional.GenerateIDNutricional()
		nutricional.IDTaco = "123"
		nutricional.FoodName = "Test"
		nutricional.Quantity = 1
		nutricional.CaloriesUnit = 100.0
		nutricional.CalculeteTotalCalories()

		repository := NewNutricionalRepository(suite.Db)
		err = repository.Save(nutricional)
		suite.NoError(err)

		nutricional.FoodName = "Test2"
		err = repository.Update(nutricional)
		suite.NoError(err)

		nutricional, err = repository.FindUnique(nutricional.ID)
		suite.NoError(err)
		suite.Equal("123", nutricional.IDTaco)
		suite.Equal("Test2", nutricional.FoodName)
		suite.Equal(1, nutricional.Quantity)
		suite.Equal(100.0, nutricional.CaloriesUnit)
		suite.Equal(100.0, nutricional.TotalCalories)
	})

	suite.Run("delete", func() {
		nutricional, err := entity.NewNutricional(
			"test 5",
			15,
			100.0,
		)
		suite.NoError(err)
		nutricional.GenerateIDNutricional()
		nutricional.IDTaco = "123"
		nutricional.FoodName = "Test"
		nutricional.Quantity = 1
		nutricional.CaloriesUnit = 100.0
		nutricional.CalculeteTotalCalories()

		repository := NewNutricionalRepository(suite.Db)
		err = repository.Save(nutricional)
		suite.NoError(err)

		err = repository.Delete(nutricional.ID)
		suite.NoError(err)

		nutricional, err = repository.FindUnique(nutricional.ID)
		suite.Error(err)
		suite.Nil(nutricional)
	})
}
