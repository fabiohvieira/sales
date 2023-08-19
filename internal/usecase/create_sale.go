package usecase

import (
	"github.com/fabiohvieira/sales/internal/entity"
)

type CreateSaleInputDto struct {
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}

type CreateSaleOutputDto struct {
	ID        string  `json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
}

type CreateSaleUseCase struct {
	SalesRepository entity.SalesRepository
}

func NewCreateSaleUseCase(salesRepository entity.SalesRepository) *CreateSaleUseCase {
	return &CreateSaleUseCase{SalesRepository: salesRepository}
}

func (u *CreateSaleUseCase) Execute(input CreateSaleInputDto) (*CreateSaleOutputDto, error) {
	sale := entity.NewSale(input.ProductID, input.Quantity, input.Price)
	sale, err := u.SalesRepository.Create(sale)
	if err != nil {
		return nil, err
	}

	return &CreateSaleOutputDto{
		ID:        sale.ID,
		ProductID: sale.ProductID,
		Quantity:  sale.Quantity,
		Price:     sale.Price,
		Total:     sale.Total,
	}, nil
}
