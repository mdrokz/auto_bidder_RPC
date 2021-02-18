package main

import (
	pb "auto_bidder/RPC/protos"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
)

const (
	port = ":50051"
)

type bidService struct {
	pb.UnimplementedBidServer
}

func (s *bidService) GetBids(ctx context.Context, in *pb.BidInput) (*pb.BidOutput, error) {
	log.Printf("Received: %v", in.Refresh)
	return &pb.BidOutput{Bids: 8}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	bidServer := grpc.NewServer()
	pb.RegisterBidServer(bidServer, &bidService{})

	if err := bidServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
