package usecase

import (
	"database/sql"
	"github/LuisMarchio03/gointensivo/internal/accounts/entity"
	"github/LuisMarchio03/gointensivo/internal/accounts/infra/database"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type FindAllUsersUsecaseTestSuite struct {
	suite.Suite
	UserRepository database.UserRepository
	Db             *sql.DB
}

func (suite *FindAllUsersUsecaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			name TEXT,
			age INTEGER,
			cpf INTEGER,
			birth TEXT,
			email TEXT,
			gender TEXT,
			password TEXT,
			street TEXT,
			city TEXT
		);
	`)
	suite.Db = db
	suite.UserRepository = *database.NewUserRepository(db)
	// suite.UserRepository = database.NewUserRepository(db)
}

func (suite *FindAllUsersUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FindAllUsersUsecaseTestSuite))
}

func (suite *FindAllUsersUsecaseTestSuite) TestFindAllUsers() {
	user, err := entity.NewUser(
		"John",
		20,
		123456789,
		"01/01/2000",
		"john@gmail.com",
		"M",
		"123456",
		"Street",
		"City",
	)
	suite.NoError(err)
	_, err = user.SetID()
	suite.NoError(err)

	err = suite.UserRepository.Save(user)
	suite.NoError(err)

	users, err := suite.UserRepository.FindAll()
	suite.NoError(err)
	suite.Equal(1, len(users))
}
