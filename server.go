package main

import (
	// 	"net"
	// 	"log"

	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Samrat-Mukherjee/practice-grpc/coffeeshop_proto"
	"google.golang.org/grpc"
	// 	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menurequest *pb.MenuRequest, svr grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Macher Jhol",
		},
		&pb.Item{
			Id:   "2",
			Name: "litti Choka",
		},
		&pb.Item{
			Id:   "3",
			Name: "Idli Dosa",
		},
	}

	for i, _ := range items {
		svr.Send(&pb.Menu{
			Item: items[0 : i+1],
		})
	}
	return nil
}
func (s *server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Recept, error) {
	return &pb.Recept{
		Id: "0xabc",
	}, nil
}
func (s *server) GetOrderStatus(context context.Context, recept *pb.Recept) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: "0xabc",
		Status:  "Pending",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Errorrr Error vagoo %v", err)
	}
	grpcServer := grpc.NewServer()
	fmt.Println("Your Server is Running at :9001")

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	err1 := grpcServer.Serve(lis)

	if err1 != nil {
		log.Fatalf("Server Error %v", err)
	}

}
