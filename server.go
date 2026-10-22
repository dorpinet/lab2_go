package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

const acceptRetryDelay = 100 * time.Millisecond

type Server struct {
	listener    net.Listener
	registry    *Registry
	commands    map[string]command
	commandHint string
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	s := &Server{listener: listener, registry: NewRegistry()}
	s.registerCommands()
	return s, nil
}

func (s *Server) Run() {
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			log.Printf(acceptFailedFormat, err)
			time.Sleep(acceptRetryDelay)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	client := NewClient(conn)

	log.Printf(connectedFormat, client.addr)
	defer log.Printf(disconnectedFormat, client.addr)

	if err := client.Send(welcomeMsg); err != nil {
		log.Printf(sendFailedFormat, client.addr, err)
		return
	}

	s.registry.Add(client)
	defer func() {
		s.registry.Remove(client)
		s.registry.Broadcast(fmt.Sprintf(leftFormat, client.addr), nil)
	}()
	s.registry.Broadcast(fmt.Sprintf(joinedFormat, client.addr), client)

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString(lineDelimiter)
		if err != nil {
			return
		}
		if !s.dispatch(client, strings.TrimSpace(line)) {
			return
		}
	}
}
