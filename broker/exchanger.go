package broker

import (
	"fmt"
	cn "queue_system/connection"
)

type Exchanger struct {
	Brokers []*Broker
}

func NewExcahnger() *Exchanger {
	return &Exchanger{}
}

func (e *Exchanger) AddBroker(br *Broker) {
	e.Brokers = append(e.Brokers, br)
}

func (e *Exchanger) FindBroker(brokerId string) (*Broker, error) {

	for i := 0; i < len(e.Brokers); i++ {

		if e.Brokers[i].BrokerId == brokerId {
			return e.Brokers[i], nil
		}
	}

	return nil, fmt.Errorf("broker with is id is not found")
}

func (e *Exchanger) RemoveBroker(con *cn.Connection, br *Broker) error {
	broker, err := e.FindBroker(con.BrokerId)

	if err != nil {
		return err
	}

	err = broker.Purge(con, false)

	if err != nil {
		return err
	}

	return nil

}

func (e *Exchanger) Exchange(con *cn.Connection, ExchangeType int8, payload *interface{}) (interface{}, error) {

	broker, err := e.FindBroker(con.BrokerId)

	if err != nil {
		return nil, err
	}

	if broker != nil {

		switch ExchangeType {

		case Connect:

			_, err = broker.Connect(con)

			if err != nil {
				return nil, err
			}

			con.Accept()

		case Disconnect:

			err = broker.Disconnect(con)

			if err != nil {
				return nil, err
			}
			con.Zombify()

		case Publish:
			err = broker.Publish(con, payload)

			if err != nil {
				return nil, err
			}

		case Consume:
			data, err := broker.Consume(con)

			if err != nil {
				return nil, err
			}

			return data, nil
		}

	}

	return nil, nil

}

func (e *Exchanger) Purge() {
	for i := 0; i < len(e.Brokers); i++ {
		e.Brokers[i].Purge(nil, true)
	}

	e.Brokers = nil
	e = nil
}
