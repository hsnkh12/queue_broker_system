package connection

import (
	"errors"
)

type ConnectionManager struct {
	Connections map[string]*Connection
}

func NewConnectinoManager() *ConnectionManager {
	return &ConnectionManager{Connections: make(map[string]*Connection)}
}

func (m *ConnectionManager) AddConnection(con *Connection) error {

	if _, ok := m.Connections[con.SessionId]; ok {
		return errors.New("connection already exists")
	}

	m.Connections[con.SessionId] = con
	return nil
}

func (m *ConnectionManager) RemoveConnection(con *Connection) error {
	delete(m.Connections, con.SessionId)
	return nil
}

func (m *ConnectionManager) Purge() {
	clear(m.Connections)
	m.Connections = nil
	m = nil
}
