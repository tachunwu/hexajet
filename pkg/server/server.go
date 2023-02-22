package server

import (
	"log"

	server "github.com/nats-io/nats-server/v2/server"
)

type Server struct {
	Server *server.Server
}

func NewServer() *Server {
	ns, err := server.NewServer(&server.Options{
		ServerName:   "hexajet",
		StoreDir:     "./data/js",
		JetStream:    true,
		NoLog:        false,
		Trace:        true,
		NoSigs:       true,
		Debug:        false,
		TraceVerbose: false,
		DontListen:   true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	ns.ConfigureLogger()
	go ns.Start()

	return &Server{
		Server: ns,
	}
}

func (s *Server) WaitForShutdown() {
	s.Server.WaitForShutdown()
}
