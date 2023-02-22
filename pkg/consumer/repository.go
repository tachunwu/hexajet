package consumer

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func RepositoryConsumer(nc *nats.Conn, subject string, logger *zap.Logger) {

	_, err := nc.Subscribe(subject, func(m *nats.Msg) {
		logger.Info(
			"in-process subscribe",
			zap.String("message", "Hello"),
			zap.String("data", string(m.Data)),
		)
		m.Ack()
	})

	if err != nil {
		logger.Info(
			"in-process subscribe error",
			zap.String("error", err.Error()),
		)
	}

}
