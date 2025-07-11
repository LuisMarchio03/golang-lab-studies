package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/LuisMarchio03/nutri/internal/entity"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Publich(ch *amqp.Channel, nutricional entity.Nutricional) error {
	body, err := json.Marshal(nutricional) // tranforma em JSON
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"nutricional",
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

	for i := 0; i < 5; i++ {
		Publich(ch, entity.Nutricional{
			FoodName:     "Arroz",
			Quantity:     100,
			CaloriesUnit: 1000.0,
		})
		fmt.Println("Enviando mensagem")
		time.Sleep(1000 * time.Millisecond)
	}
}
