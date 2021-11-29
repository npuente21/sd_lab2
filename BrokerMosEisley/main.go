package main

import (
	pb "Lab2/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

type BrokerServer struct {
	pb.UnimplementedBrokerServicesServer
}

func (s *BrokerServer) AddCity(context.Context, *pb.RequestInf) (*pb.ResponseBroker, error) {
	return &pb.ResponseBroker{Address: "localhost:40000"}, nil
}

func main() {
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBrokerServicesServer(grpcServer, &BrokerServer{})
	log.Printf("server listening at %v", listner.Addr())

	if err = grpcServer.Serve(listner); err != nil {
		log.Fatalf("Failed to listen on port 50000: %v", err)
	}
}
