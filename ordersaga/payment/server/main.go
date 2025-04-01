package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "ordersaga/proto/payment"

	"google.golang.org/grpc"
)

type paymentServer struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *paymentServer) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	fmt.Println("ProcessPayment:", req.OrderId, req.Amount)
	return &pb.ProcessPaymentResponse{
		Status: "SUCCESS",
	}, nil
}

func (s *paymentServer) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
	fmt.Println("RefundPayment:", req.OrderId)
	return &pb.RefundPaymentResponse{
		Status: "REFUNDED",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &paymentServer{})
	fmt.Println("💳 PaymentService running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
