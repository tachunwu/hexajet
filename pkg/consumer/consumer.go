package consumer

import "github.com/nats-io/nats.go"

type Consumer struct {
	Subject  string
	HandleFn func(*nats.Msg)
}
