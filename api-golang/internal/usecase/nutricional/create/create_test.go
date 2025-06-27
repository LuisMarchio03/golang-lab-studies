package usecase

import (
	"database/sql"
	"testing"

	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CreateNutricionalUsecaseTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *CreateNutricionalUsecaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE nutricional (id VARCHAR(255) NOT NULL, id_taco VARCHAR(255) NOT NULL, food_name VARCHAR(255) NOT NULL, quantity FLOAT NOT NULL, calories_unit FLOAT NOT NULL, total_calories FLOAT NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *CreateNutricionalUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CreateNutricionalUsecaseTestSuite))
}

func (suite *CreateNutricionalUsecaseTestSuite) TestCreateNutricional() {
	nutricionalRepository := repository.NutricionalRepository{Db: suite.Db}

	createNutricionalInput := CreateNutricionalInput{
		FoodName:     "Test",
		Quantity:     10,
		CaloriesUnit: 100.0,
	}

	createNutricionalUsecase := NewCreateNutricionalUsecase(nutricionalRepository)
	err := createNutricionalUsecase.Execute(createNutricionalInput)
	suite.NoError(err)
}
