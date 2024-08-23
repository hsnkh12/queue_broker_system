package broker

import (
	"errors"
	cn "queue_system/connection"
	"queue_system/queue"
	"queue_system/utils"
	"sync"

	"github.com/google/uuid"
)

type Broker struct {
	Connections map[string]*cn.Connection
	BrokerId    string
	Key         string
	Queue       *queue.Queue
	Mu          sync.Mutex
}

func NewBroker() *Broker {
	brokerId := uuid.New().String()
	key := utils.GenerateKey()
	return &Broker{BrokerId: brokerId, Key: key, Queue: queue.New(), Connections: make(map[string]*cn.Connection)}
}

func (b *Broker) Connect(con *cn.Connection) (*Broker, error) {

	b.Mu.Lock()
	defer b.Mu.Unlock()

	if b.BrokerId != con.BrokerId {
		return nil, errors.New("borker id is invalid")
	}

	if b.Key != con.Key {
		return nil, errors.New("key is invalid")
	}

	if con.Status != cn.New {
		return nil, errors.New("connection status must be new")
	}

	if _, ok := b.Connections[con.SessionId]; ok {
		return nil, errors.New("connection with this session already exists")
	}
	b.Connections[con.SessionId] = con

	return b, nil
}

func (b *Broker) Disconnect(con *cn.Connection) error {

	b.Mu.Lock()
	defer b.Mu.Unlock()

	if _, ok := b.Connections[con.SessionId]; !ok {
		return errors.New("connection is invalid")
	}

	delete(b.Connections, con.SessionId)

	return nil
}

func (b *Broker) Publish(con *cn.Connection, payload interface{}) error {

	b.Queue.Mu.Lock()
	defer b.Queue.Mu.Unlock()

	if _, ok := b.Connections[con.SessionId]; !ok {
		return errors.New("connection is invalid")
	}

	if con.Type != cn.Publisher {
		return errors.New("user must be a publisher")
	}

	b.Queue.Enqueue(payload)

	return nil
}

func (b *Broker) Consume(con *cn.Connection) (interface{}, error) {
	b.Queue.Mu.Lock()
	defer b.Queue.Mu.Unlock()

	if _, ok := b.Connections[con.SessionId]; !ok {
		return nil, errors.New("connection is invalid")
	}

	if con.Type != cn.Consumer {
		return nil, errors.New("user must be a consumer")
	}

	if b.Queue.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	data := b.Queue.Dequeue()
	return data, nil
}

func (b *Broker) Purge(con *cn.Connection, force bool) error {

	if !force {
		if _, ok := b.Connections[con.SessionId]; !ok {
			return errors.New("connection is invalid")
		}

		if con.Type != cn.Publisher {
			return errors.New("user must be a publisher")
		}
	}

	clear(b.Connections)
	b.Queue = nil
	b = nil
	return nil
}
