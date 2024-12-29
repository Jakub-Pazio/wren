package server

import (
	"fmt"
	"log/slog"
	"net"
)

type Server struct {
	users      map[int]connection
	cp         clientPool
	serverName string
}

func New() Server {
	return Server{
		users: make(map[int]connection),
		cp:    NewClientPool(),
	}
}

func (s *Server) Run() error {
	l, err := net.Listen("tcp", ":6667")
	if err != nil {
		panic(err)
	}

	connCounter := 0

	connChan := make(chan channelCommand)
	respChan := make(chan serverReply)

	defer func() {
		l.Close()
	}()
	go func() {
		for {
			c, err := l.Accept()
			slog.Info("new connection, trying to create connection object")
			if err != nil {
				fmt.Println(err)
				continue
			}
			conn := connection{
				netConn: c,
				id:      connCounter,
				ch:      connChan,
				rch:     respChan,
			}
			s.users[connCounter] = conn
			connCounter++
			go conn.Run()
			slog.Info("connection established for user")
		}
	}()
	for {
		msg := <-connChan
		switch msg.Name {
		case "NewNICK":
			err := s.cp.AddClientNick(msg.Args[0], msg.ConnId)
			if err != nil {
				s.users[msg.ConnId].rch <- serverReply{Err: err}
				continue
			}
			s.users[msg.ConnId].rch <- serverReply{Response: "ok"}
		case "UpdateNICK":
			err := s.cp.UpdateClientNick(msg.Args[0], msg.Args[1], msg.ConnId)
			if err != nil {
				s.users[msg.ConnId].rch <- serverReply{Err: err}
				continue
			}
			s.users[msg.ConnId].rch <- serverReply{Response: "ok"}

		case "USER":
			s.users[msg.ConnId].rch <- serverReply{Response: "ok"}
		case "QUIT":
			s.cp.RemoveClient(msg.Args[0])
		}
		slog.Info("Clients in the server", "users", fmt.Sprintf("%v", s.cp.ListUsers()))
	}
}
