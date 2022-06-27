package rabbitmq

import (
	"election-service/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
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
		"write_csv", // queue name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
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
			body := string(delivery.Body)
			WriteToCSV(body)
		}
	}()

	<-forever

}

func WriteToCSV(body string) {
	// Parser json
	var vote model.Vote
	err := json.Unmarshal([]byte(body), &vote)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to parser json", err)
	}

	// Read
	file, err := os.Open(os.Getenv("CSV_FILE"))
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open csv file", err)
	}

	data, err := csv.NewReader(file).ReadAll()
	file.Close()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to create csv reader", err)
	}

	// Prepare data
	candidateId := fmt.Sprintf("%v", vote.CandidateID)
	nationalId := vote.NationalID

	row := []string{candidateId, nationalId}
	data = append(data, row)

	// Create csv file
	file, err = os.Create(os.Getenv("CSV_FILE"))
	if err != nil {
		log.Fatalf("%s: %s", "Failed to create csv file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.WriteAll(data)
}
