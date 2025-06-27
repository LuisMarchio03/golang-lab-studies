package main

import (
	"encoding/json"
	"github/LuisMarchio03/gointensivo/internal/order/entity"
	"math/rand"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publich(ch *amqp.Channel, order entity.Order) error {
	body, err := json.Marshal(order) // tranforma em JSON
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

func GenerateOrders() entity.Order { // Gera um order de forma aleatoria
	return entity.Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100, // valor aleatorio
		Tax:   rand.Float64() * 10,  // valor aleatorio
	}
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

	for i := 0; i < 1000; i++ {
		Publich(ch, GenerateOrders())
		time.Sleep(300 * time.Millisecond)
	}
}
