package queuesystemmanager

import (
	"queue_system/broker"
	"queue_system/connection"
)

type QueueSystemManager struct {
	ConnectionManager *connection.ConnectionManager
	Exchanger         *broker.Exchanger
}

func Init() *QueueSystemManager {
	conManager := connection.NewConnectinoManager()
	exchanger := broker.NewExcahnger()
	return &QueueSystemManager{ConnectionManager: conManager, Exchanger: exchanger}
}

func (qm *QueueSystemManager) Purge() {

	qm.Exchanger.Purge()
	qm.ConnectionManager.Purge()

}
