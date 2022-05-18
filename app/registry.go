package app

import (
	"github.com/Jehanv60/gotoko/app/models"
)

type Model struct {
	Model interface{}
}

func Registermodels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Category{}},
		{Model: models.Section{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Payment{}},
		{Model: models.Shipment{}},
	}
}
