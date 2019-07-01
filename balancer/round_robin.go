package balancer

import (
	"sync"

	"github.com/kismia/rabbitmq-go-client"
)

type roundRobinBalancer struct {
	addrs  []string
	nextMx sync.Mutex
	next   int
}

func NewRoundRobinBalancer(addrs []string) rabbitmq.Balancer {
	return &roundRobinBalancer{
		addrs: addrs,
	}
}

func (b *roundRobinBalancer) Balance() string {
	b.nextMx.Lock()

	addr := b.addrs[b.next]

	b.next = (b.next + 1) % len(b.addrs)

	b.nextMx.Unlock()

	return addr
}
