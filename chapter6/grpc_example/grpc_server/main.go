package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc_example/protofiles"
	"log"
	"net"
)

// server is used to create MoneyTransactionServer.
type server struct{}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	// Use in.Amount, in.From, in.To and perform transaction logic
	return &pb.TransactionResponse{Confirmation: true}, nil
}

var port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("cannot start server on port ", port)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
