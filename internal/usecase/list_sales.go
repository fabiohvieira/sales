package usecase

import (
	"github.com/fabiohvieira/sales/internal/entity"
)

type FilterSalesInputDto struct {
	ID        string `json:"id"`
	ProductID int64  `json:"product_id"`
}

type ListSalesInputDto struct {
	ID        string  `json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
}

type ListSalesUseCase struct {
	SalesRepository entity.SalesRepository
}

func NewListSalesUseCase(salesRepository entity.SalesRepository) *ListSalesUseCase {
	return &ListSalesUseCase{SalesRepository: salesRepository}
}

func (u *ListSalesUseCase) Execute(filters FilterSalesInputDto) ([]*ListSalesInputDto, error) {
	sales, err := u.SalesRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var salesDto []*ListSalesInputDto
	for _, sale := range sales {
		salesDto = append(salesDto, &ListSalesInputDto{
			ID:        sale.ID,
			ProductID: sale.ProductID,
			Quantity:  sale.Quantity,
			Price:     sale.Price,
			Total:     sale.Total,
		})
	}

	return salesDto, nil
}
