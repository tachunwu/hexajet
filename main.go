package main

import (
	"runtime"

	"github.com/tachunwu/hexajet/pkg/consumer"
	"github.com/tachunwu/hexajet/pkg/jetstream"
	"github.com/tachunwu/hexajet/pkg/producer"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	nc := jetstream.NewNATS()

	go producer.PublishEmitter(nc, "repo", logger)
	consumer.RepositoryConsumer(nc, "repo", logger)

	for {
		runtime.Gosched()
	}
}
