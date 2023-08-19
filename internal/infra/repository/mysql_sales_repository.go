package repository

import (
	"database/sql"
	"fmt"

	"github.com/fabiohvieira/sales/internal/entity"
)

type MySQLSalesRepository struct {
	Conn *sql.DB
}

func NewMySQLSalesRepository(conn *sql.DB) *MySQLSalesRepository {
	return &MySQLSalesRepository{Conn: conn}
}

func (r *MySQLSalesRepository) Create(sale *entity.Sale) (*entity.Sale, error) {
	stmt, err := r.Conn.Prepare(`INSERT INTO sales (id, product_id, quantity, price, total) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = stmt.Exec(sale.ID, sale.ProductID, sale.Quantity, sale.Price, sale.Total)
	if err != nil {
		return nil, err
	}

	return sale, nil
}

func (r *MySQLSalesRepository) FindAll() ([]*entity.Sale, error) {
	rows, err := r.Conn.Query(`SELECT id, product_id, quantity, price, total FROM sales`)
	if err != nil {
		return nil, err
	}

	var sales []*entity.Sale
	for rows.Next() {
		sale := &entity.Sale{}
		err := rows.Scan(&sale.ID, &sale.ProductID, &sale.Quantity, &sale.Price, &sale.Total)
		if err != nil {
			return nil, err
		}

		sales = append(sales, sale)
	}

	return sales, nil
}
