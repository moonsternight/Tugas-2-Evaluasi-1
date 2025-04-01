package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "ordersaga/proto/order"

	"google.golang.org/grpc"
)

type orderServer struct {
    pb.UnimplementedOrderServiceServer
}

func (s *orderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    fmt.Println("CreateOrder:", req.UserId, req.ItemId)
    return &pb.CreateOrderResponse{
        OrderId: "ORDER123",
        Status:  "PENDING",
    }, nil
}

func (s *orderServer) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
    fmt.Println("CancelOrder:", req.OrderId)
    return &pb.CancelOrderResponse{
        Status: "CANCELLED",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterOrderServiceServer(s, &orderServer{})
    fmt.Println("🚀 OrderService running on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
