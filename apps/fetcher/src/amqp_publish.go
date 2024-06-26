package main

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type ExchangeRateMessage struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

func publishToExchangeRate(message ExchangeRateMessage) error {
	conn, err := amqp091.Dial(AMQP_CONNECTION)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"exchange-rate", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(
		"exchange-rate-queue", // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	err = ch.QueueBind(queue.Name, "", "exchange-rate", false, nil)

	if err != nil {
		return err
	}

	// Serialize the message to JSON
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Publish a message to the exchange
	err = ch.Publish(
		"exchange-rate", // exchange
		"",              // routing key
		false,           // mandatory
		false,           // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", body)
	return nil
}
