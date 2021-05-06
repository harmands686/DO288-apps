package main

import (
	"log"

	"amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-1.rabbitmq.rabbits.svc.cluster.local:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
                "HelloExchange",   // name
                "topic", // type
                true,     // durable
                false,    // auto-deleted
                false,    // internal
                false,    // no-wait
                nil,      // arguments
        )
	failOnError(err, "Failed to declare an exchange")
	
	q, err := ch.QueueDeclare(
		"HelloQueue", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
  		q.Name, // queue name
  		q.Name,     // routing key
  		"HelloExchange", // exchange
  		false,
  		nil,
	)
	failOnError(err, "Failed to bind the queue to exchange")
	
	body := "Hello World!"
	err = ch.Publish(
		"HelloExchange",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
