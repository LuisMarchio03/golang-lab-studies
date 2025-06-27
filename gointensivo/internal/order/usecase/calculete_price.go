package usecase

import (
	"github/LuisMarchio03/gointensivo/internal/order/entity"
	"github/LuisMarchio03/gointensivo/internal/order/infra"
	"github/LuisMarchio03/gointensivo/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculeteFinalPriceUsecase struct {
	OrderRepository infra.OrderRepositoryInterface
}

func NewCalculeteFinalPriceUsecase(
	orderRepository database.OrderRepository,
	// orderRepository infra.OrderRepositoryInterface,
) *CalculeteFinalPriceUsecase {
	return &CalculeteFinalPriceUsecase{
		OrderRepository: &orderRepository,
	}
}

func (c *CalculeteFinalPriceUsecase) Execute(
	input OrderInputDTO,
) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
