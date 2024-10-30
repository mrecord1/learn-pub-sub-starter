package pubsub

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	// Marshal the val to JSON bytes
	body, err := json.Marshal(val)
	if err != nil {
		return err
	}

	// Publish the message to the exchange with the routing key
	ctx := context.Background()
	err = ch.PublishWithContext(ctx, // context
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", body)
	return nil
}

func PublishGOB[T any](ch *amqp.Channel, exchange, key string, val T) error {
	// Encode to GOB
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(val)
	if err != nil {
		return nil
	}
	body := buffer.Bytes()

	// Publish the message to the exchange with the routing key
	ctx := context.Background()
	err = ch.PublishWithContext(ctx, // context
		exchange, // exchange
		key,      // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/gob",
			Body:        body,
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", body)
	return nil
}
