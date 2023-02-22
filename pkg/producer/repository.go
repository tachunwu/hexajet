package producer

import (
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func PublishEmitter(nc *nats.Conn, subject string, logger *zap.Logger) {

	for {
		msg := nats.NewMsg(subject)
		msg.Data = []byte("Hi")
		nc.PublishMsg(msg)
		logger.Info(
			"in-process repository publish",
			zap.String("message", "repository push command/event"),
		)
		time.Sleep(5 * time.Second)
	}
}
