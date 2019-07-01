package recovery

import (
	"sync/atomic"

	"github.com/streadway/amqp"
)

type atomicChannel struct {
	channel atomic.Value
}

func newAtomicChannel(v *amqp.Channel) *atomicChannel {
	ch := &atomicChannel{}
	ch.store(v)

	return ch
}

func (c *atomicChannel) load() *amqp.Channel {
	return c.channel.Load().(*amqp.Channel)
}

func (c *atomicChannel) store(v *amqp.Channel) {
	c.channel.Store(v)
}

type atomicConnection struct {
	connection atomic.Value
}

func newAtomicConnection(v *amqp.Connection) *atomicConnection {
	conn := &atomicConnection{}
	conn.store(v)

	return conn
}

func (c *atomicConnection) load() *amqp.Connection {
	return c.connection.Load().(*amqp.Connection)
}

func (c *atomicConnection) store(v *amqp.Connection) {
	c.connection.Store(v)
}
