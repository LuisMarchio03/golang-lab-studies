package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	userRepository "github.com/LuisMarchio03/nutri/internal/infra/repository/user"
	web "github.com/LuisMarchio03/nutri/internal/infra/web"
	calculateDietUsecase "github.com/LuisMarchio03/nutri/internal/usecase/user/calculate"
	findUserInfosUsecase "github.com/LuisMarchio03/nutri/internal/usecase/user/find"
	updateDietUsecase "github.com/LuisMarchio03/nutri/internal/usecase/user/update"

	"github.com/go-chi/chi/v5"

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

	userRepo := userRepository.NewUserRepository(db)
	calculateDietUC := calculateDietUsecase.CalculateDietUsecase{UserRepository: userRepo}
	updateDietUC := updateDietUsecase.UpdateDietUsecase{UserRepository: userRepo}
	findUserInfosUC := findUserInfosUsecase.FindUserInfosUsecase{UserRepository: userRepo}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, out, "users")

	for i := 1; i <= 10; i++ {
		go worker(out, &calculateDietUC, i)
	}

	usersHandlers := web.NewUserHandlers(&updateDietUC, &findUserInfosUC)

	r := chi.NewRouter()
	r.Put("/v1/user", usersHandlers.UpdateUserHandler)
	r.Get("/v1/user", usersHandlers.ListUserHandler)

	http.ListenAndServe(":8080", r)
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *calculateDietUsecase.CalculateDietUsecase, workerID int) {
	for msg := range deliveryMessage {
		var inputDTO calculateDietUsecase.CalculateDietInput

		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println("Worker", workerID, "received a message:", outputDTO)
		time.Sleep(500 * time.Microsecond)
	}
}
