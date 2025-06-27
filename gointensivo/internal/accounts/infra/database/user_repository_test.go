package database

import (
	"database/sql"
	"github/LuisMarchio03/gointensivo/internal/accounts/entity"
	"testing"

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
	db.Exec("CREATE TABLE users (id varchar(255) NOT NULL, name varchar(255) NOT NULL, age int NOT NULL, cpf int NOT NULL, birth varchar(255) NOT NULL, email varchar(255) NOT NULL, gender varchar(255) NOT NULL, password varchar(255) NOT NULL, street varchar(255) NOT NULL, city varchar(255) NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestGivenAnUser_ThenShouldUser() {
	suite.Run("TestGivenAnUser_WithSave_ThenShouldReturnUsers", func() {
		user, err := entity.NewUser(
			"John",
			20,
			123456789,
			"01/01/2000",
			"john@email.com",
			"M",
			"123456",
			"Street",
			"City",
		)
		suite.NoError(err)
		id, err := user.SetID()
		suite.NoError(err)
		repo := NewUserRepository(suite.Db)
		err = repo.Save(user)
		suite.NoError(err)

		var userResult entity.User
		err = suite.Db.QueryRow("Select id, name, age, cpf, birth, email, gender, password, street, city from users where id = ?", id).
			Scan(&userResult.ID, &userResult.Name, &userResult.Age, &userResult.Cpf, &userResult.Birth, &userResult.Email, &userResult.Gender, &userResult.Password, &userResult.Street, &userResult.City)

		suite.NoError(err)
		suite.Equal(id, userResult.ID)
		suite.Equal(user.Name, userResult.Name)
		suite.Equal(user.Age, userResult.Age)
		suite.Equal(user.Cpf, userResult.Cpf)
		suite.Equal(user.Birth, userResult.Birth)
		suite.Equal(user.Email, userResult.Email)
		suite.Equal(user.Gender, userResult.Gender)
		suite.Equal(user.Password, userResult.Password)
		suite.Equal(user.Street, userResult.Street)
		suite.Equal(user.City, userResult.City)
	})

	suite.Run("TestGivenAnUser_WithFindAll_ThenShouldReturnUsers", func() {
		stmt, err := suite.Db.Prepare("INSERT INTO users (id, name, age, cpf, birth, email, gender, password, street, city) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		suite.NoError(err)
		_, err = stmt.Exec(
			"1234",
			"John2",
			20,
			123456789,
			"01/01/2000",
			"john2@email.com",
			"M",
			"123456",
			"Street",
			"City",
		)
		suite.NoError(err)
		repo := NewUserRepository(suite.Db)
		users, err := repo.FindAll()
		suite.NoError(err)

		suite.Equal(2, len(users))
	})
}
