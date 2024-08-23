package connection

import (
	"sync"

	"github.com/google/uuid"
)

type Connection struct {
	BrokerId  string
	Key       string
	Type      int8 // 1: Publisher, 2: Consumer
	SessionId string
	Status    int8 // 0: New, 1: Accepted, 2: Zmobie
	Mu        sync.Mutex
}

func NewConnection(brokerId string, key string, type_ int8) *Connection {

	sessionId := uuid.New().String()
	return &Connection{BrokerId: brokerId, Key: key, Type: type_, Status: 0, SessionId: sessionId}
}

func (con *Connection) Accept() {
	con.Mu.Lock()
	con.Status = 1
	con.Mu.Unlock()
}

func (con *Connection) Zombify() {
	con.Mu.Lock()
	con.Status = 2
	con.Mu.Unlock()
}
