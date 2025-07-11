package repository

import (
	"database/sql"
	"testing"

	"github.com/LuisMarchio03/nutri/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE users (id VARCHAR(255) NOT NULL,name VARCHAR(255) NOT NULL,email VARCHAR(255) NOT NULL,password VARCHAR(255) NOT NULL,height FLOAT NOT NULL,weight FLOAT NOT NULL,age INT NOT NULL,gender VARCHAR(255) NOT NULL,goal INT NOT NULL, total_calories FLOAT, total_proteinas FLOAT, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestGivenAnUser_ThenShouldUser() {
	suite.Run("Should create a user", func() {
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
		userRepository := NewUserRepository(suite.Db)
		err := userRepository.Save(user)
		suite.NoError(err)

		var id string
		err = suite.Db.QueryRow("SELECT id FROM users WHERE id = ?", user.ID).Scan(&id)
		suite.NoError(err)
		suite.Equal(user.ID, id)

		var total_calories float64
		err = suite.Db.QueryRow("SELECT total_calories FROM users WHERE id = ?", user.ID).Scan(&total_calories)
		suite.NoError(err)
		suite.Equal(user.TotalCalories, total_calories)

		var total_proteinas float64
		err = suite.Db.QueryRow("SELECT total_proteinas FROM users WHERE id = ?", user.ID).Scan(&total_proteinas)
		suite.NoError(err)
		suite.Equal(user.TotalProteinas, total_proteinas)
	})

	suite.Run("Should find all users", func() {
		userRepository := NewUserRepository(suite.Db)
		users, err := userRepository.FindAll()
		suite.NoError(err)
		suite.Equal(1, len(users))
	})

	suite.Run("Should find infos user", func() {
		user2 := &entity.User{
			Name:     "Teste2",
			Email:    "teste2@email.com",
			Password: "1234567",
			Height:   1.87,
			Weight:   80.5,
			Age:      22,
			Gender:   "M",
			Goal:     1,
		}
		user2.SetID()
		userRepository := NewUserRepository(suite.Db)
		err := userRepository.Save(user2)
		suite.NoError(err)
		res, err := userRepository.FindInfosUser(user2.ID)
		suite.NoError(err)

		suite.Equal(user2.Name, res.Name)
		suite.Equal(user2.Email, res.Email)
		suite.Equal(user2.Height, res.Height)
		suite.Equal(user2.Weight, res.Weight)
		suite.Equal(user2.Age, res.Age)
		suite.Equal(user2.Gender, res.Gender)
		suite.Equal(user2.Goal, res.Goal)
	})

	suite.Run("Should update user", func() {
		user := &entity.User{
			Name:     "Teste Update",
			Email:    "teste_update@email.com",
			Password: "123456",
			Height:   1.80,
			Weight:   80.0,
			Age:      20,
			Gender:   "M",
			Goal:     1,
		}

		userRepository := NewUserRepository(suite.Db)
		err := userRepository.Update(user, user.ID)
		suite.NoError(err)
	})

}
