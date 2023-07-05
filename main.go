package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Hoc vien ky thuat mat ma")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Success")
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)

	err = ch.Publish(
		"", "TestQueue", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		})
	if err != nil {
		panic(err)
	}
}
