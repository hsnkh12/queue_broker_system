package broker

import (
	"fmt"
	"math/rand"
)

type Producer struct {
	UserId int
	Broker *Broker
}

func NewProducer(borker *Broker) *Producer {
	id := rand.Intn(101)
	return &Producer{UserId: id, Broker: borker}
}

func (p *Producer) Publish(data interface{}) {
	defer p.Broker.Unlock()
	p.Broker.Lock()
	p.Broker.Queue.Enqueue(data)
	fmt.Println("Published: ", data)
}
