package main

import (
	"net"
	"net/http"
)

type Server struct {
	Addr string // Server address
}

// Function to create the server
func createServer() *Server {
	return &Server{
		Addr: ":8080", // Setting the server address (port 8080)
	}
}

// StartServer function to run the server with a custom handler
func (srv *Server) StartServer() error {
	addr := srv.Addr
	// Listen on the specified address
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// Use the default handler registered in processor.go
	return http.Serve(ln, nil)
}
