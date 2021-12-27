package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	bassoon "github.com/MaxencePPierre/effective-bassoon/message"
	"github.com/nsqio/go-nsq"
	zlog "github.com/rs/zerolog/log"
)

type messageHandler struct{}

const (
	localhost string = "127.0.0.1"
	port      string = "4161"
)

func main() {
	//The only valid way to instantiate the Config
	config := nsq.NewConfig()

	//Tweak several common setup in config
	// Maximum number of times this consumer will attempt to process a message before giving up
	config.MaxAttempts = 10
	// Maximum number of messages to allow in flight
	config.MaxInFlight = 5
	// Maximum duration when REQueueing
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0
	//Init topic name and channel
	channel := "Channel_Example"

	//Creating the consumer
	consumer, err := nsq.NewConsumer(bassoon.Topic, channel, config)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Failed to create new consumer")
	}

	// Set the Handler for messages received by this Consumer.
	consumer.AddHandler(&messageHandler{})
	//Use nsqlookupd to find nsqd instances
	err = consumer.ConnectToNSQLookupd(localhost + ":" + port)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Connection failed")
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	// Gracefully stop the consumer.
	consumer.Stop()
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	var request bassoon.Message
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)
		return err
	}

	log.Println("Message")
	log.Println("--------------------")
	log.Println("Name : ", request.Name)
	log.Println("Content : ", request.Content)
	log.Println("Timestamp : ", request.Timestamp)
	log.Println("--------------------")
	log.Println("")

	// Will automatically set the message as finish
	return nil
}
