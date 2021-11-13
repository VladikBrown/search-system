package mq

import (
	"github.com/streadway/amqp"
	"log"
)

func SetUpMQConnection(url string, queueName string) (
	conn *amqp.Connection,
	ch *amqp.Channel,
	q amqp.Queue,
	err error,
) {
	conn, err = amqp.Dial(url)
	ch, err = conn.Channel()
	q, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	log.Println("Connection established")
	return
}

func ReceiveFromQueue(ch *amqp.Channel, q amqp.Queue, dataFunc func([]byte)) {
	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
			dataFunc(d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
