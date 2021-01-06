package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("rabbitmq consumer")

	conn, err := amqp.Dial("amqps://username:password@url/")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer conn.Close()

	fmt.Println("connected to the rabbitmq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message is: %s\n", d.Body)
		}
	}()

	<-forever
}
