package rabbitmq

import (
	"candidate-service/config"
	"candidate-service/model"
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ReceiveVote() {
	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	failOnError(err, "Failed to connect RabbitMQ")
	defer connection.Close()

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"vote",   // exchange name
		"direct", // type
		true,     // durable
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	queue, err := channel.QueueDeclare(
		"vote", // queue name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.QueueBind(
		queue.Name,      // queue name
		"vote_increase", // routing key
		"vote",          // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	messages, err := channel.Consume(
		queue.Name,          // queue name
		"candidate_service", // consumer name
		true,                // auto ack
		false,               // exclusive
		false,               // no local
		false,               // no wait
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for delivery := range messages {
			body := string(delivery.Body)
			UpdateVoteCount(body)
		}
	}()

	<-forever
}

func UpdateVoteCount(body string) {
	db := config.DB

	// Parser json
	var vote model.Vote
	err := json.Unmarshal([]byte(body), &vote)
	if err != nil {
		log.Println("Can not parse json to vote model")
	}

	// Find candidate
	var candidate model.Candidate
	err = db.Model(&model.Candidate{}).Where("id = ?", vote.CandidateID).First(&candidate).Error
	if err != nil {
		log.Println("Candidate not found")
	}

	// Update voted count
	votedCount := *candidate.VotedCount + 1
	err = db.Model(&model.Candidate{}).Where("id = ? AND updated_at = ?", candidate.ID, candidate.UpdatedAt).Update("voted_count", votedCount).Error
	if err != nil {
		log.Println("Can not update voted count")
	}
}
