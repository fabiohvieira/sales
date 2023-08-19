package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/fabiohvieira/sales/internal/infra/akafka"
	"github.com/fabiohvieira/sales/internal/infra/repository"
	"github.com/fabiohvieira/sales/internal/infra/web"
	"github.com/fabiohvieira/sales/internal/usecase"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/sales")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repository := repository.NewMySQLSalesRepository(conn)
	createSalesUseCase := usecase.NewCreateSaleUseCase(repository)
	listSalesUseCase := usecase.NewListSalesUseCase(repository)

	salesHandlers := web.NewSalesHandlers(createSalesUseCase, listSalesUseCase)

	r := chi.NewRouter()
	r.Post("/sales", salesHandlers.CreateSaleHandlers)
	r.Get("/sales", salesHandlers.ListSalesHandlers)

	go http.ListenAndServe(":3000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"sales"}, "kafka:9092", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateSaleInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			log.Fatal(err)
		}
		_, err = createSalesUseCase.Execute(dto)
		if err != nil {
			log.Fatal(err)
		}
	}
}
