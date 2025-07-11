package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	nutriRepository "github.com/LuisMarchio03/nutri/internal/infra/repository/nutricional"
	createNutriUsecase "github.com/LuisMarchio03/nutri/internal/usecase/nutricional/create"
	"github.com/LuisMarchio03/nutri/pkg/rabbitmq"

	_ "github.com/mattn/go-sqlite3"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	nutriRepo := nutriRepository.NewNutricionalRepository(db)
	createNutriUC := createNutriUsecase.CreateNutricionalUsecase{NutricionalRepository: nutriRepo}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, out, "nutricional")

	for i := 1; i <= 10; i++ {
		go worker(out, &createNutriUC, i)
	}

	http.ListenAndServe(":8080", nil)
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *createNutriUsecase.CreateNutricionalUsecase, workerID int) {
	for msg := range deliveryMessage {
		var inputDTO createNutriUsecase.CreateNutricionalInput

		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		err = uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Worker", workerID, "received a message:", inputDTO)
		time.Sleep(500 * time.Microsecond)
	}
}
