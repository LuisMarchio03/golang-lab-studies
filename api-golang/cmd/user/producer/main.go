package main

import (
	"encoding/json"
	"time"

	"github.com/LuisMarchio03/nutri/internal/entity"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublichUser(ch *amqp.Channel, user entity.User) error {
	body, err := json.Marshal(user) // tranforma em JSON
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	for i := 0; i < 15; i++ {
		PublichUser(ch, entity.User{
			Name:     "Luis",
			Email:    "lg@gmail.com",
			Password: "123",
			Weight:   70,
			Height:   1.70,
			Age:      20,
			Gender:   "M",
			Goal:     1,
		})
		time.Sleep(1000 * time.Millisecond)
	}
}
