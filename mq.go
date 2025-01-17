package gorabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/duffya90/go-rabbitmq/connection"
	"github.com/duffya90/go-rabbitmq/exchange"
	"github.com/duffya90/go-rabbitmq/queue"
)

type MQ struct {
	connection *connection.Connection
	channel    *amqp.Channel
	queue      *queue.Queue
	exchange   *exchange.Exchange
}

func New(url string) (*MQ, error) {
	conn, err := connection.New(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		connection: conn,
		channel:    ch,
		queue:      queue.New(ch),
		exchange:   exchange.New(ch),
	}, nil
}

func NewFromConnection(conn *connection.Connection) (*MQ, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		connection: conn,
		channel:    ch,
	}, nil
}

func (mq *MQ) Connection() *connection.Connection {
	return mq.connection
}

func (mq *MQ) Channel() *amqp.Channel {
	return mq.channel
}

func (mq *MQ) Queue() *queue.Queue {
	return mq.queue
}

func (mq *MQ) Exchange() *exchange.Exchange {
	return mq.exchange
}

func (mq *MQ) Publish(publish *MQConfigPublish) error {
	return mq.channel.Publish(
		publish.Exchange,
		publish.RoutingKey,
		publish.Mandatory,
		publish.Immediate,
		publish.Message,
	)
}

func (mq *MQ) Close() {
	mq.channel.Close()
	mq.connection.Close()
}
