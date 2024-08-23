package main

import (
	"fmt"
	qsm "queue_system/queue_system_manager"
	"queue_system/server"
)

func main() {

	q := qsm.Init()
	server := server.CreateServer("localhost:3000", q)

	go func() {
		for msg := range server.ReceiveBuffer {

			fmt.Printf("< Message \n < Headers: address: %s > \n < Payload: %s > \n >", msg.Header.FromAddress, msg.Payload)
			response := "< Message from " + msg.Header.FromAddress + " Received>"

			select {
			case server.SendBuffer <- response:
			case <-server.QuitChannel:
				return
			}

		}
	}()

	server.Listen()

}
