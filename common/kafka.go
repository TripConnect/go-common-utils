package common

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/tripconnect/go-common-utils/helper"
)

var kafkaAddres, _ = helper.ReadConfig[string]("kafka.host")
var kafkaPort, _ = helper.ReadConfig[int]("kafka.port")
var KafkaConnection = fmt.Sprintf("%s:%d", kafkaAddres, kafkaPort)

var KafkaPublisher = &kafka.Writer{
	Addr:                   kafka.TCP(KafkaConnection),
	Balancer:               &kafka.LeastBytes{},
	AllowAutoTopicCreation: true,
}

func Publish(ctx context.Context, topic string, data interface{}) error {
	// Publish message to Kafka, data should be passed as pointer
	if valueBytes, err := json.Marshal(data); err == nil {
		KafkaPublisher.WriteMessages(ctx, kafka.Message{
			Topic: topic,
			Value: []byte(valueBytes),
		})
	} else {
		return fmt.Errorf("Publish kafka message failed %v", err)
	}

	return nil
}
