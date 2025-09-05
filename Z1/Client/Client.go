package main

import (
	"context"
	"log"
	"time"

	"Z1/protobufapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := protobufapi.NewChatServiceClient(conn)

	for i := 0; i < 3; i++ {
		resp, err := c.SayHello(context.Background(), &protobufapi.Message{Body: "66666"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %v", err)
		}
		log.Printf("Response from server: %s", resp.Body)

		resp2, err := c.Test(context.Background(), &protobufapi.Message{Body: "123"})
		if err != nil {
			log.Fatalf("Error when calling Test: %v", err)
		}
		log.Printf("Response from server: %s", resp2.Body)

		time.Sleep(1 * time.Second)
	}
}
