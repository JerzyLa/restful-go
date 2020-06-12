package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc_example/protofiles"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot setup connection to server")
	}
	defer conn.Close()
	// Create a client
	c := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Make a server request.
	r, err := c.MakeTransaction(context.Background(),
		&pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatal("cannot make transaction")
	}
	log.Println("Transaction confirmed: ", r.Confirmation)
}
