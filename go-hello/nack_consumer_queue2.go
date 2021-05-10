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
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-balancer:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	

	msgs, err := ch.Consume(
		"hello_queue2", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a subscriber")

	for d := range msgs {
  		d.Nack(false,false)	
  		log.Printf("Nack::Requeue=false >> Received a message: %s", d.Body)
	}

}
