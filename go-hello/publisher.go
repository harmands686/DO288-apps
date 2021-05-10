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
        //conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-1.rabbitmq.rabbits.svc.cluster.local:5672/")
        conn, err := amqp.Dial("amqp://guest:guest@rabbitmq-balancer:5672/")
        failOnError(err, "Failed to connect to RabbitMQ")
        defer conn.Close()

        ch, err := conn.Channel()
        failOnError(err, "Failed to open a channel")
        defer ch.Close()

        //*************Exchange declaration*************

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

        //*****************Queues declarations*************************
	
		//********adding x-arguments for queue***********
		args := make(amqp.Table)
		args["x-queue-type"] = "quorum"

	q1, err := ch.QueueDeclare(
                "hello_queue1", // name
                true,   // durable
                false,   // delete when unused
                false,   // exclusive
                false,   // no-wait
                args,     // arguments
        )
        failOnError(err, "Failed to declare a queue")

        q2, err := ch.QueueDeclare(
               "hello_queue2", // name
                true,   // durable
                false,   // delete when unused
                false,   // exclusive
                false,   // no-wait
                args,     // arguments
        )
        failOnError(err, "Failed to declare a queue")


        //********************Queues bindings**************************
        err = ch.QueueBind(
                q1.Name, // queue name
                "hello.*",     // routing key
                "HelloExchange", // exchange
                false,    //nowait
                nil,
        )
        failOnError(err, "Failed to bind the queue to exchange")

         err = ch.QueueBind(
                q2.Name, // queue name
                "hello.*",     // routing key
                "HelloExchange", // exchange
                false,    //nowait
                nil,
        )
        failOnError(err, "Failed to bind the queue to exchange")


        //**************Message publish***********************
        body := "Hello World!"
        err = ch.Publish(
                "HelloExchange",     // exchange
                "hello.world", // routing key
		 false,  // mandatory
                false,  // immediate
                amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte(body),
                })
        failOnError(err, "Failed to publish a message")
        log.Printf(" [Routing.Key=hello.world] Sent %s", body)
}
