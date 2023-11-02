package server

import (
	"context"
	"errors"
	calculator "main/proto"
)

// server struct

type Server struct {
	calculator.UnimplementedCalculatorServiceServer
}

// sum grpc Request handler

func (s *Server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	var result int32
	result = 0
	for _, i := range req.A {
		result += i
	}
	return &calculator.SumResponse{Result: result}, nil
}

// Subtraction grpc Request handler

func (s *Server) Subtraction(ctx context.Context, req *calculator.SubtractionRequest) (*calculator.SubtractionResponse, error) {
	var result int32
	result = req.A[0]
	for _, i := range req.A {
		if result == i {
			continue
		}
		result -= i
	}
	return &calculator.SubtractionResponse{Result: result}, nil
}

// Multiplication grpc Request handler

func (s *Server) Multiplication(ctx context.Context, req *calculator.MultiplicationRequest) (*calculator.MultiplicationResponse, error) {
	var result int32
	result = 1
	for _, i := range req.A {

		result *= i
	}
	return &calculator.MultiplicationResponse{Result: result}, nil
}

// Multiplication grpc Request handler

func (s *Server) Division(ctx context.Context, req *calculator.DivisionRequest) (*calculator.DivisionResponse, error) {
	var result int32
	result = req.A[0]
	for _, i := range req.A {
		if result == i {
			continue
		}
		if i == 0 {
			return nil, errors.New("error: number Division zero")
		}
		result /= i
	}
	return &calculator.DivisionResponse{Result: result}, nil
}
