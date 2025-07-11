package usecase

import (
	"database/sql"
	"testing"

	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculateDietUsecaseTestSuite struct {
	suite.Suite
	UserRepository repository.UserRepository
	Db             *sql.DB
}

func (suite *CalculateDietUsecaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id VARCHAR(255) NOT NULL,name VARCHAR(255) NOT NULL,email VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,height FLOAT NOT NULL,weight FLOAT NOT NULL,age INT NOT NULL,gender VARCHAR(255) NOT NULL,goal INT NOT NULL, total_calories FLOAT, total_proteinas FLOAT, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *CalculateDietUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateDietUsecaseTestSuite))
}

func (suite *CalculateDietUsecaseTestSuite) TestCalculateDiet() {
	userRepository := repository.UserRepository{Db: suite.Db}
	calculateDietInput := CalculateDietInput{
		Name:     "Luis",
		Email:    "luis@gmail.com",
		Password: "123456",
		Height:   1.80,
		Weight:   80.2,
		Age:      20,
		Gender:   "M",
		Goal:     1,
	}

	calculateDietUsecase := NewCalculateDietUsecase(userRepository)
	user, err := calculateDietUsecase.Execute(calculateDietInput)
	suite.NoError(err)

	suite.Equal(user.Name, calculateDietInput.Name)
	suite.Equal(user.Email, calculateDietInput.Email)
	suite.Equal(user.Password, calculateDietInput.Password)
	suite.Equal(user.Height, calculateDietInput.Height)
	suite.Equal(user.Weight, calculateDietInput.Weight)
	suite.Equal(user.Age, calculateDietInput.Age)
	suite.Equal(user.Gender, calculateDietInput.Gender)
	suite.Equal(user.Goal, calculateDietInput.Goal)
}
