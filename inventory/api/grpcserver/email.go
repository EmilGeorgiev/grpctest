package grpcserver

import (
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"

	pb "github.com/proxiad/grpctest/api/inventory/v1"
)

type emailInventoryServiceServer struct {
	store map[string]*pb.Email
}

// NewInventoryServiceServer return and initialize  pb.TrackintObjectRepositoryServer
func NewInventoryServiceServer() pb.EmailInventoryServiceServer {
	return emailInventoryServiceServer{store: make(map[string]*pb.Email)}
}

func (ei emailInventoryServiceServer) CreateEmail(ctx context.Context, dto *pb.CreateEmailRequest) (*empty.Empty, error) {
	e := &pb.Email{
		Recipient: dto.Recipient,
		Subject: dto.Subject,
		Content: dto.Content,
	}

	ei.store[e.Recipient] = e

	log.Printf("Email with recipient: %s is stored successfully\n", e.Recipient)

	return new(empty.Empty), nil
}

func (ei emailInventoryServiceServer) FindEmail(ctx context.Context, req *pb.FindEmailRequest) (*pb.FindEmailResponse, error) {
	e, ok := ei.store[req.Recipient]
	if !ok {
		log.Printf("Email with recipient: %s is not found\n", req.Recipient)
	}

	return &pb.FindEmailResponse {
		Emial: e,
	}, nil
}
