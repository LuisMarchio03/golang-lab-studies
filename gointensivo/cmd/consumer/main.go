package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/LuisMarchio03/gointensivo/internal/order/infra/database"
	"github/LuisMarchio03/gointensivo/internal/order/usecase"
	"github/LuisMarchio03/gointensivo/pkg/rabbitmq"
	"time"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.CalculeteFinalPriceUsecase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery) // Channel
	go rabbitmq.Consumer(ch, out)   // T2

	for msg := range out {
		var inputDTO usecase.OrderInputDTO

		err := json.Unmarshal(msg.Body, &inputDTO) // Converte um JSON
		if err != nil {
			panic(err)
		}
		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println(outputDTO)
		time.Sleep(500 * time.Microsecond)
	}
}
