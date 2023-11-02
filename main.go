package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	calculator "main/proto"
	"net"
)

type Server struct {
	calculator.UnimplementedCalculatorServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	srv := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(srv, &Server{})
	reflection.Register(srv)
	fmt.Println("Server is listening on port 50051")
	if err := srv.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
