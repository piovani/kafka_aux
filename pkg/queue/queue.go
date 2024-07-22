package queue

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/piovani/kafka_aux/config"
)

type Queue struct {
	brokers []string
	topic   string
}

func NewQueue() *Queue {
	return &Queue{
		brokers: []string{config.Env.KafkaHost},
		topic:   config.Env.KafkaTopic,
	}
}

func (q *Queue) Publish(msg string) error {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(q.brokers, config)
	if err != nil {
		return err
	}

	msgK := &sarama.ProducerMessage{
		Topic:     q.topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}

	partition, offset, err := producer.SendMessage(msgK)
	if err != nil {
		return err
	}

	fmt.Println(partition)
	fmt.Println(offset)

	return nil
}
