package main

import (
	"flag"
	"log"
	"fmt"

	"google.golang.org/grpc"
	"golang.org/x/net/context"

	pb "github.com/proxiad/grpctest/api/inventory/v1"
)

var (
	listenerRPCHost = flag.String("listener-rpc", "localhost:8686", "listener gRPC server host")
)

func main () {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	gconn, err := grpc.Dial(*listenerRPCHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("The client gRPC connection cannot be created - %v", err)
	}

	defer gconn.Close()

	emailRPC := pb.NewEmailInventoryServiceClient(gconn)

	req1 := &pb.CreateEmailRequest{
		Recipient:  "recipient@gmail.com",
		Subject: "gRPC",
		Content: "Hello World!",
	}
	emailRPC.CreateEmail(context.Background(), req1)

	req2 := &pb.FindEmailRequest{
		Recipient: "recipient@gmail.com",
	}

	resp, err := emailRPC.FindEmail(context.Background(), req2)

	fmt.Printf("Response from server: %#v\n", resp.Emial)
}


