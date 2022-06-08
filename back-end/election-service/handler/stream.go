package handler

import (
	"election-service/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/streadway/amqp"
)

func CandidateVoteStream(ws *websocket.Conn) {
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
		"",   // queue name
		true, // durable
		true, // delete when unused
		true, // exclusive
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
		queue.Name,         // queue name
		"election_service", // consumer name
		true,               // auto ack
		false,              // exclusive
		false,              // no local
		false,              // no wait
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for delivery := range messages {
			// Delay for candidate-service update
			time.Sleep(time.Millisecond * 100)

			// Parser json
			body := string(delivery.Body)

			var vote model.Vote
			err := json.Unmarshal([]byte(body), &vote)
			if err != nil {
				log.Fatalf("%s: %s", "Failed to parser json", err)
			}

			// Check ID
			candidateId := ws.Params("candidateId")
			voteId := fmt.Sprintf("%v", vote.CandidateID)

			if candidateId == voteId {
				count := getCandidate(candidateId)
				// Send to web socket connection
				err = ws.WriteJSON(count)
				if err != nil {
					log.Fatalf("%s: %s", "Failed to write json", err)
				}
			}
		}
	}()

	<-forever
}

func AllCandidateVoteStream(ws *websocket.Conn) {
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
		"",    // queue name
		true,  // durable
		false, // delete when unused
		true,  // exclusive
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
		queue.Name,         // queue name
		"election_service", // consumer name
		true,               // auto ack
		false,              // exclusive
		false,              // no local
		false,              // no wait
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for delivery := range messages {
			// Delay for candidate-service update
			time.Sleep(time.Millisecond * 100)

			// Parser json
			body := string(delivery.Body)

			var vote model.Vote
			err := json.Unmarshal([]byte(body), &vote)
			if err != nil {
				log.Fatalf("%s: %s", "Failed to parser json", err)
			}

			candidateId := fmt.Sprintf("%v", vote.CandidateID)
			count := getCandidate(candidateId)

			// Send to web socket connection
			err = ws.WriteJSON(count)
			if err != nil {
				log.Fatalf("%s: %s", "Failed to write json", err)
			}
		}
	}()

	<-forever
}

func getCandidate(id string) model.Count {
	// Get all candidate
	var count model.Count
	httpRequest("http://localhost:5002/api/candidates/"+id, &count)

	return count
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
