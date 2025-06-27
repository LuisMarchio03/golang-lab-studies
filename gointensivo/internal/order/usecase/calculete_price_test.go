package usecase

import (
	"database/sql"
	"github/LuisMarchio03/gointensivo/internal/order/entity"
	"github/LuisMarchio03/gointensivo/internal/order/infra/database"
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type CalculateFinalPriceUsecaseTestSuite struct {
	suite.Suite
	// OrderRepository infra.OrderRepositoryInterface
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculateFinalPriceUsecaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	_, err = db.Exec(`
		CREATE TABLE orders (
			id TEXT PRIMARY KEY,	
			price REAL,
			tax REAL,
			final_price REAL
		);
	`)
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
	// suite.OrderRepository = database.NewOrderRepository(db)
}

func (suite *CalculateFinalPriceUsecaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculateFinalPriceUsecaseTestSuite))
}

func (suite *CalculateFinalPriceUsecaseTestSuite) TestCalculeFinalPrice() {
	order, err := entity.NewOrder("1", 10, 2)
	suite.NoError(err)
	order.CalculateFinalPrice()

	calculateFinalPriceInput := OrderInputDTO{
		ID:    order.ID,
		Price: order.Price,
		Tax:   order.Tax,
	}

	calculateFinalPriceUsecase := NewCalculeteFinalPriceUsecase(suite.OrderRepository)
	output, err := calculateFinalPriceUsecase.Execute(calculateFinalPriceInput)
	suite.NoError(err)

	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)
}
