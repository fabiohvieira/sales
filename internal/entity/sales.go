package entity

import (
	"github.com/google/uuid"
)

type SalesRepository interface {
	Create(sale *Sale) (*Sale, error)
	FindAll() ([]*Sale, error)
}

type Sale struct {
	ID        string  `json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
}

func NewSale(productID, quantity int64, price float64) *Sale {
	return &Sale{
		ID:        uuid.New().String(),
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		Total:     price * float64(quantity),
	}
}
