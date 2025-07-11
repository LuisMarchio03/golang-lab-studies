package usecase

import (
	"database/sql"
	"testing"

	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type FindAllNutricionalUsecaseTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *FindAllNutricionalUsecaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE nutricional (id VARCHAR(255) NOT NULL, id_taco VARCHAR(255) NOT NULL, food_name VARCHAR(255) NOT NULL, quantity FLOAT NOT NULL, calories_unit FLOAT NOT NULL, total_calories FLOAT NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *FindAllNutricionalUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FindAllNutricionalUsecaseTestSuite))
}

func (suite *FindAllNutricionalUsecaseTestSuite) TestFindAllNutricional() {
	nutricionalRepository := repository.NutricionalRepository{Db: suite.Db}

	_, err := suite.Db.Exec("INSERT INTO nutricional (id, id_taco, food_name, quantity, calories_unit, total_calories) VALUES (?, ?, ?, ?, ?, ?)", "123", "345", "Test Food Name", 2, 100.0, 200.0)
	suite.NoError(err)

	findAllNutricionalUsecase := NewFindAllNutricionalUsecase(nutricionalRepository)
	output, err := findAllNutricionalUsecase.Execute()
	suite.NoError(err)
	suite.Equal(1, len(output))
}
