package broker

import (
	"math/rand"
	"queue_system/queue"
	"sync"

	"github.com/google/uuid"
)

type BrokerI interface {
	Lock()
	Unlock()
}

type Broker struct {
	BrokerId string
	Key      string
	Queue    *queue.Queue
	Mu       sync.Mutex
}

func NewBroker() *Broker {
	brokerId := uuid.New().String()
	key := generateKey()
	return &Broker{BrokerId: brokerId, Key: key, Queue: queue.New()}
}

func (b *Broker) Lock() {
	b.Mu.Lock()
}

func (b *Broker) Unlock() {
	b.Mu.Unlock()
}

func generateKey() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
