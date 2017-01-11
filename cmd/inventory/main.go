package main

import (
	"log"

	"net"
	"flag"

	"google.golang.org/grpc"
	"github.com/proxiad/grpctest/inventory/api/grpcserver"

	pb "github.com/proxiad/grpctest/api/inventory/v1"
)

var (
	rpcPort = flag.String("rpc", "8686", "RPC Server port")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()

	errc := make(chan error)

	go func() {
		addr := ":" + *rpcPort
		l, err := net.Listen("tcp", addr)

		if err != nil {
			errc <- err
			return
		}

		log.Println("Application started successfully.")

		s := grpc.NewServer()

		pb.RegisterEmailInventoryServiceServer(s, grpcserver.NewInventoryServiceServer())

		if err := s.Serve(l); err != nil {
			log.Printf("The RPC server cannot be started - %s", err)
			errc <- err
		}
	}()

	err := <-errc
	if err != nil {
		log.Printf("could not start due: %v\n", err)
	}
}
