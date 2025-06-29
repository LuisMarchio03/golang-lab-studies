package infra

import "github/LuisMarchio03/gointensivo/internal/order/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	GetTotal() (int, error)
}
