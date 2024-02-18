package main

import (
	"fmt"
	"queue_system/broker"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	b := broker.NewBroker()

	producer := broker.NewProducer(b)

	consumer := broker.NewConsumer()

	consumer.Subscribe(b)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			producer.Publish(fmt.Sprintf("message %d", i+1))
		}
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		for {
			consumer.Consume()
			if consumer.Broker.Queue.IsEmpty() {
				break
			}
		}

	}()

	wg.Wait()

}
