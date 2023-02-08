package main

import (
	"fmt"
	"log"
	"net"

	pb "groupchat/generated/groupchatpb"
	"groupchat/internal/api"
	"groupchat/internal/client"
	"groupchat/internal/group"

	"google.golang.org/grpc"
)

const (
	// port default port for TCP listener
	port = 13111
)

func main() {
	// Init TCP listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Services and implementation
	clientService := client.New()
	groupService := group.New()
	impl := api.NewAPI(clientService, groupService)

	// Server and registration
	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, impl)

	// Starting tcp listener
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
