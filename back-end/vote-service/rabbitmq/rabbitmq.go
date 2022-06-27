package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func PublishVoteMessage(body string) {
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	failOnError(err, "Failed to connect RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"vote",   // name
		"direct", // type
		true,     // durable
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	err = channel.Publish(
		"vote",          // exchange name
		"vote_increase", // routing key name
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")
}
