package usecase

import (
	"database/sql"
	"testing"

	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type UpdateDietUsecaseTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *UpdateDietUsecaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id VARCHAR(255) NOT NULL,name VARCHAR(255) NOT NULL,email VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,height FLOAT NOT NULL,weight FLOAT NOT NULL,age INT NOT NULL,gender VARCHAR(255) NOT NULL,goal INT NOT NULL, total_calories FLOAT, total_proteinas FLOAT, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *UpdateDietUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UpdateDietUsecaseTestSuite))
}

func (suite *UpdateDietUsecaseTestSuite) TestUpdateDiet() {
	userRepository := repository.UserRepository{Db: suite.Db}
	updateDietInput := UpdateDietInput{
		Name:     "Luis",
		Email:    "luis@gmail.com",
		Password: "123456",
		Height:   1.80,
		Weight:   80.2,
		Age:      20,
		Gender:   "M",
		Goal:     1,
	}

	calculateDietUsecase := NewUpdateDietUsecase(userRepository)
	user, err := calculateDietUsecase.Execute("1", updateDietInput)
	suite.NoError(err)

	suite.Equal(user.Name, updateDietInput.Name)
	suite.Equal(user.Email, updateDietInput.Email)
	suite.Equal(user.Password, updateDietInput.Password)
	suite.Equal(user.Height, updateDietInput.Height)
	suite.Equal(user.Weight, updateDietInput.Weight)
	suite.Equal(user.Age, updateDietInput.Age)
	suite.Equal(user.Gender, updateDietInput.Gender)
	suite.Equal(user.Goal, updateDietInput.Goal)
}
