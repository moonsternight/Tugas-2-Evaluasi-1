package main

import (
	"context"
	"log"
	"time"

	orderpb "ordersaga/proto/order"
	paymentpb "ordersaga/proto/payment"
	shippingpb "ordersaga/proto/shipping"

	"google.golang.org/grpc"
)

func main() {
	// Connect ke OrderService
	orderConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to OrderService: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// Connect ke PaymentService
	paymentConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to PaymentService: %v", err)
	}
	defer paymentConn.Close()
	paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)

	// Connect ke ShippingService
	shippingConn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to ShippingService: %v", err)
	}
	defer shippingConn.Close()
	shippingClient := shippingpb.NewShippingServiceClient(shippingConn)

	// 1. Create Order
	orderRes, err := orderClient.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{
		UserId: "user1",
		ItemId: "item1",
	})
	if err != nil {
		log.Fatalf("CreateOrder failed: %v", err)
	}
	log.Println("🛒 Order Created:", orderRes.OrderId)

	// 2. Process Payment
	paymentRes, err := paymentClient.ProcessPayment(context.Background(), &paymentpb.ProcessPaymentRequest{
		OrderId: orderRes.OrderId,
		Amount:  100.0,
	})
	if err != nil || paymentRes.Status != "SUCCESS" {
		log.Println("💳 Payment failed, cancelling order...")
		orderClient.CancelOrder(context.Background(), &orderpb.CancelOrderRequest{OrderId: orderRes.OrderId})
		return
	}
	log.Println("💳 Payment Success")

	// 3. Start Shipping
	shipRes, err := shippingClient.StartShipping(context.Background(), &shippingpb.StartShippingRequest{
		OrderId: orderRes.OrderId,
	})
	if err != nil || shipRes.Status != "SHIPPED" {
		log.Println("📦 Shipping failed, doing compensation...")
		shippingClient.CancelShipping(context.Background(), &shippingpb.CancelShippingRequest{ShippingId: shipRes.ShippingId})
		paymentClient.RefundPayment(context.Background(), &paymentpb.RefundPaymentRequest{OrderId: orderRes.OrderId})
		orderClient.CancelOrder(context.Background(), &orderpb.CancelOrderRequest{OrderId: orderRes.OrderId})
		return
	}

	log.Println("📦 Shipping Success, Saga Completed 🎉")
	time.Sleep(time.Second)
}
