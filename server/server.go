package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	qsm "queue_system/queue_system_manager"
	"sync"
	"syscall"
)

type Server struct {
	Address            string
	Listener           net.Listener
	QuitChannel        chan struct{}
	ReceiveBuffer      chan Message
	SendBuffer         chan string
	Wg                 sync.WaitGroup
	ActiveConnections  map[net.Conn]struct{}
	ActiveConnsMux     sync.Mutex
	QueueSystemManager *qsm.QueueSystemManager
}

func CreateServer(addr string, q *qsm.QueueSystemManager) *Server {
	return &Server{
		Address:            addr,
		QuitChannel:        make(chan struct{}),
		ReceiveBuffer:      make(chan Message, 10),
		SendBuffer:         make(chan string, 10),
		ActiveConnections:  make(map[net.Conn]struct{}),
		QueueSystemManager: q,
	}
}

func (s *Server) Listen() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.Listener = listener

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signalChan
		fmt.Printf("Received signal %s, shutting down...\n", sig)

		close(s.QuitChannel)

		s.QueueSystemManager.Purge()
		s.CloseAllConnections()
	}()

	s.Wg.Add(1)
	go s.AcceptConnections()

	fmt.Println("Server is listening on", s.Address)

	<-s.QuitChannel

	close(s.ReceiveBuffer)
	close(s.SendBuffer)

	return nil
}

func (s *Server) AcceptConnections() {
	defer s.Wg.Done()

	for {
		select {
		case <-s.QuitChannel:
			return
		default:
		}

		conn, err := s.Listener.Accept()

		if err != nil {
			log.Printf("error accepting connection: %v", err)
			continue
		}

		s.ActiveConnsMux.Lock()
		s.ActiveConnections[conn] = struct{}{}
		s.ActiveConnsMux.Unlock()

		log.Println("new connection:", conn.RemoteAddr())

		s.Wg.Add(1)
		go s.ReadConneciton(conn)
	}
}

func (s *Server) ReadConneciton(conn net.Conn) {
	defer conn.Close()
	defer s.Wg.Done()

	buffer := make([]byte, 2050)

	for {
		bytesRead, err := conn.Read(buffer)

		if err != nil {
			log.Printf("Connection closed: %s", conn.RemoteAddr().Network())

			s.ActiveConnsMux.Lock()
			delete(s.ActiveConnections, conn)
			s.ActiveConnsMux.Unlock()

			return
		}

		header := HeaderMessage{
			FromAddress: conn.RemoteAddr().String(),
		}

		s.ReceiveBuffer <- Message{
			Header:  header,
			Payload: buffer[:bytesRead],
		}

		respond := <-s.SendBuffer

		conn.Write([]byte(respond))
	}
}

func (s *Server) CloseAllConnections() {
	s.ActiveConnsMux.Lock()
	defer s.ActiveConnsMux.Unlock()

	for conn := range s.ActiveConnections {
		conn.Close()
		delete(s.ActiveConnections, conn)
	}
}
