package usecase

import (
	"database/sql"
	"testing"

	"github.com/LuisMarchio03/nutri/internal/entity"
	repository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type FindUserInfosUsecaseTestSuite struct {
	suite.Suite
	UserRepository repository.UserRepository
	Db             *sql.DB
}

func (suite *FindUserInfosUsecaseTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id VARCHAR(255) NOT NULL,name VARCHAR(255) NOT NULL,email VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,height FLOAT NOT NULL,weight FLOAT NOT NULL,age INT NOT NULL,gender VARCHAR(255) NOT NULL,goal INT NOT NULL, total_calories FLOAT, total_proteinas FLOAT, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *FindUserInfosUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FindUserInfosUsecaseTestSuite))
}

func (suite *FindUserInfosUsecaseTestSuite) TestFindUserInfos() {
	userRepository := repository.UserRepository{Db: suite.Db}
	user := &entity.User{
		Name:     "Teste",
		Email:    "teste@email.com",
		Password: "123456",
		Height:   1.80,
		Weight:   80.0,
		Age:      20,
		Gender:   "M",
		Goal:     1,
	}
	err := userRepository.Save(user)
	suite.NoError(err)

	FindUserInfosUsecase := NewFindUserInfosUsecase(userRepository)
	res, err := FindUserInfosUsecase.Execute(user.ID)
	suite.NoError(err)

	suite.Equal(res.Name, user.Name)
	suite.Equal(res.Email, user.Email)
	suite.Equal(res.Height, user.Height)
	suite.Equal(res.Weight, user.Weight)
	suite.Equal(res.Age, user.Age)
	suite.Equal(res.Gender, user.Gender)
	suite.Equal(res.Goal, user.Goal)
}
