package main

import (
	"github.com/nats-io/nats.go"
	"github.com/tachunwu/hexajet/pkg/consumer"
	"github.com/tachunwu/hexajet/pkg/producer"

	"go.uber.org/zap"
)

var subject = "repository"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Client
	nc, _ := nats.Connect(nats.DefaultURL)
	// Subscribe
	consumer.RepositoryConsumer(nc, subject, logger)
	// Publish
	producer.PublishEmitter(nc, subject, logger)

}
