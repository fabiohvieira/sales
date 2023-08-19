package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fabiohvieira/sales/internal/usecase"
)

type SalesHandlers struct {
	CreateProductUseCase *usecase.CreateSaleUseCase
	ListSalesUseCase     *usecase.ListSalesUseCase
}

func NewSalesHandlers(createProductUseCase *usecase.CreateSaleUseCase, listSalesUseCase *usecase.ListSalesUseCase) *SalesHandlers {
	return &SalesHandlers{
		CreateProductUseCase: createProductUseCase,
		ListSalesUseCase:     listSalesUseCase,
	}
}

func (s *SalesHandlers) CreateSaleHandlers(w http.ResponseWriter, r *http.Request) {
	dto := usecase.CreateSaleInputDto{}
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sale, err := s.CreateProductUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sale)
}

func (s *SalesHandlers) ListSalesHandlers(w http.ResponseWriter, r *http.Request) {
	product_id, err := strconv.ParseInt(r.URL.Query().Get("product_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	filters := usecase.FilterSalesInputDto{
		ID:        r.URL.Query().Get("id"),
		ProductID: product_id,
	}

	sales, err := s.ListSalesUseCase.Execute(filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sales)
}
