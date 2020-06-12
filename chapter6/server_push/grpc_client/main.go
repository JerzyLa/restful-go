package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "server_push/protofiles"
)

const (
	address = "localhost:50051"
)

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	// Listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// If there are no more messages, get out of loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status,
			response.Description)
	}
}

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
	ReceiveStream(c, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
