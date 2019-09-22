package main

import (
	"context"
	"log"

	"envoy-grpc-json/sample/msglib"
	"google.golang.org/grpc"
)

func main() {
	// conn ==> RPC Client, RPC Client.rpcMethod
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := msglib.NewPersonRegistryClient(conn)
	person, err := client.Lookup(context.Background(), &msglib.Person{})
	if err != nil {
		log.Fatalf("Error fetching person details: %v", err)
	}
	log.Printf("Found person: %s %d\n", *person.Name, *person.Age)
}
