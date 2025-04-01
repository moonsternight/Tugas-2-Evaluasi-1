package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "ordersaga/proto/shipping"

	"google.golang.org/grpc"
)

type shippingServer struct {
	pb.UnimplementedShippingServiceServer
}

func (s *shippingServer) StartShipping(ctx context.Context, req *pb.StartShippingRequest) (*pb.StartShippingResponse, error) {
	fmt.Println("StartShipping:", req.OrderId)
	return &pb.StartShippingResponse{
		ShippingId: "SHIP123",
		Status:     "SHIPPED",
	}, nil
}

func (s *shippingServer) CancelShipping(ctx context.Context, req *pb.CancelShippingRequest) (*pb.CancelShippingResponse, error) {
	fmt.Println("CancelShipping:", req.ShippingId)
	return &pb.CancelShippingResponse{
		Status: "CANCELLED",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &shippingServer{})
	fmt.Println("📦 ShippingService running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
