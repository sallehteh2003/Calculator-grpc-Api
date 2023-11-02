package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	calculator "main/proto"
	"main/server"
	"net"
)

const Port = 3000

func main() {
	// Create a tcp server listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", Port))
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}

	// Create grpc server
	srv := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(srv, &server.Server{})
	reflection.Register(srv)
	fmt.Println(fmt.Sprintf("Server is listening on port:%v", Port))
	if err := srv.Serve(listener); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
