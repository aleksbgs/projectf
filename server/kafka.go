package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Notificater interface {
	CreateTopic(topic string) error
	StartWriter()
	Publish(input string) error
	Initialize() error
}

type KafkaInstance struct {
	Ctx     context.Context
	Brokers []string
	Topic   string
	Writer  *kafka.Writer
}

func NewKafkaInstance() KafkaInstance {
	return KafkaInstance{}
}

// CreateTopic creates a topic in kafka environment
func (k *KafkaInstance) CreateTopic(topic string) error {
	_, err := kafka.DialLeader(k.Ctx, "tcp", k.Brokers[0], topic, 0)
	if err != nil {
		return err
	}
	return nil
}

func (k *KafkaInstance) StartWriter() {
	k.Writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.Brokers,
		Topic:   k.Topic,
	})
}

func (k *KafkaInstance) Publish(input string) error {
	if err := k.Writer.WriteMessages(k.Ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(1)),
		Value: []byte(input),
	}); err != nil {
		return err
	}
	return nil
}

func (k *KafkaInstance) Initialize() error {
	log.Println("Initializing Kafka connection ...")

	KAFKA_BROKER := "broker"

	KAFKA_BROKER_PORT := "29092"

	brokerstr := fmt.Sprintf("%s:%s", KAFKA_BROKER, KAFKA_BROKER_PORT)

	c := make(chan Result)
	waitfor := 60

	for i := 0; i < waitfor; i++ {
		time.Sleep(1 * time.Second)

		go func(broker string) {
			kafka := KafkaInstance{
				Brokers: []string{broker},
				Ctx:     context.Background(),
				Topic:   "user_events",
			}

			if err := kafka.CreateTopic(kafka.Topic); err != nil {
				log.Println("error creating topic, kafka is not ready")
				result := Result{
					Success: false,
					Error:   err,
					Kafka:   nil,
				}
				c <- result
				return
			}
			kafka.StartWriter()

			result := Result{
				Success: true,
				Error:   nil,
				Kafka:   &kafka,
			}
			c <- result
		}(brokerstr)

		select {
		case res := <-c:
			if res.Success {
				fmt.Println("***** Kafka is ready *****")
				k.Writer = res.Kafka.Writer
				k.Brokers = res.Kafka.Brokers
				k.Ctx = res.Kafka.Ctx
				k.Topic = res.Kafka.Topic
				return nil
			}
		case <-time.After(time.Duration(waitfor) * time.Second):
			fmt.Println("timeout %n", waitfor)
		}
	}
	return fmt.Errorf("error establishing kafka connection")
}

type DataEvent struct {
	Action string
}

type Result struct {
	Success bool
	Error   error
	Kafka   *KafkaInstance
}

// User is a struct representing the user entity
