package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	desc "github.com/obeismo/auth/grpc/pkg/auth/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err)
	}
	defer conn.Close()

	c := desc.NewAuthV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: noteID})
	if err != nil {
		log.Fatalf("failed to get user by id: %s", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetInfo()))
}
