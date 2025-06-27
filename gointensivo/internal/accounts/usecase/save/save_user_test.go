package usecase

import (
	"database/sql"
	"github/LuisMarchio03/gointensivo/internal/accounts/entity"
	"github/LuisMarchio03/gointensivo/internal/accounts/infra/database"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type SaveUserUsecaseTestSuite struct {
	suite.Suite
	OrderRepository database.UserRepository
	Db              *sql.DB
}

func (suite *SaveUserUsecaseTestSuite) SetupTest() {
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
	suite.OrderRepository = *database.NewUserRepository(db)
	// suite.OrderRepository = database.NewOrderRepository(db)
}

func (suite *SaveUserUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SaveUserUsecaseTestSuite))
}

func (suite *SaveUserUsecaseTestSuite) TestSaveUser() {
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

	inputSaveUser := InputSaveUserDTO{
		Name:     user.Name,
		Age:      user.Age,
		Cpf:      user.Cpf,
		Birth:    user.Birth,
		Email:    user.Email,
		Gender:   user.Gender,
		Password: user.Password,
		Street:   user.Street,
		City:     user.City,
	}

	saveUserUsecase := NewSaveUserUsecase(suite.OrderRepository)
	err = saveUserUsecase.Execute(inputSaveUser)
	suite.NoError(err)
}
