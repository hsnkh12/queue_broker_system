package broker

import (
	"fmt"
	"math/rand"
)

type Consumer struct {
	UserId int
	Broker *Broker
}

func NewConsumer() *Consumer {
	id := rand.Intn(101)
	return &Consumer{UserId: id}
}

func (c *Consumer) Subscribe(broker *Broker) {
	c.Broker = broker
}

func (c *Consumer) Consume() interface{} {
	defer c.Broker.Unlock()
	c.Broker.Lock()
	value := c.Broker.Queue.Dequeue()
	fmt.Println("Consumed: ", value)
	return value
}

func (c *Consumer) Unsubscribe() {
	c.Broker = nil
}
