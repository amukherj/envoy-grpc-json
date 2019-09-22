package main

import (
	"fmt"
	"net"

	"envoy-grpc-json/sample/msglib"
	"google.golang.org/grpc"
)

func main() {
	// 1. Create an instance of the service implementation.
	// 2. Create an instance of the grpc server.
	// 3. Associate 1 with 2, using the appropriate Register...Server call.
	// 4. Create a listener.
	// 5. Serve the grpc server using the listener.
	// So:
	//     service --> grpc server <== listener

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		fmt.Printf("Failed to start listening: %v\n", err)
	}

	svr := msglib.PersonRegistryService{}
	grpcSvr := grpc.NewServer()
	msglib.RegisterPersonRegistryServer(grpcSvr, &svr)
	if err := grpcSvr.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
