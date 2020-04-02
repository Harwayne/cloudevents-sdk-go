package nats

import (
	"errors"
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/nats-io/nats.go"
)

var ErrInvalidQueueName = errors.New("invalid queue name for QueueSubscriber")

// NatsOptions is a helper function to group a variadic stan.ProtocolOption into
// []stan.Option that can be used by either Sender, Consumer or Protocol
func NatsOptions(opts ...nats.Option) []nats.Option {
	return opts
}

// ProtocolOption is the function signature required to be considered an nats.ProtocolOption.
type ProtocolOption func(*Protocol) error

func WithConsumerOptions(opts ...ConsumerOption) ProtocolOption {
	return func(p *Protocol) error {
		p.consumerOptions = opts
		return nil
	}
}

func WithSenderOptions(opts ...SenderOption) ProtocolOption {
	return func(p *Protocol) error {
		p.senderOptions = opts
		return nil
	}
}

type SenderOption func(*Sender) error

// WithTransformer adds a transformer, which Sender uses while encoding a binding.Message to an nats.Message
func WithTransformer(transformer binding.TransformerFactory) SenderOption {
	return func(s *Sender) error {
		s.Transformers = append(s.Transformers, transformer)
		return nil
	}
}

type ConsumerOption func(*Consumer) error

// WithQueueSubscriber configures the Consumer to join a queue group when subscribing
func WithQueueSubscriber(queue string) ConsumerOption {
	return func(c *Consumer) error {
		if queue == "" {
			return ErrInvalidQueueName
		}
		c.Subscriber = &QueueSubscriber{Queue: queue}
		return nil
	}
}
