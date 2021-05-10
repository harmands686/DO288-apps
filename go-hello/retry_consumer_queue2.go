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
                false,   // auto-ack
                false,  // exclusive
                false,  // no-local
                false,  // no-wait
                nil,    // args
        )
        failOnError(err, "Failed to register a subscriber2")
   for d := range msgs {
                log.Printf("Received a message: %s", d.Body)
                d.Ack(false)

                //if any error then execute below

                log.Printf("Failure processing message\n\n")

                 err := ch.Publish(
                "RetryExchange",     // exchange
                "retry_delay_queue", // routing key
                 false,  // mandatory
                false,  // immediate
                amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        d.Body,
                })
        failOnError(err, "Failed to publish a message")
        }
}
