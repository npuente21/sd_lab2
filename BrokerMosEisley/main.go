package main

import (
	pb "Lab2/proto"
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/grpc"
)

//var address_1 = []string{"10.6.40.181:50023", "10.6.40.182:50023", "10.6.40.184:50023"}

var address_1 = []string{"localhost:50023", "localhost:50023", "localhost:50023"}

const (
	port = ":50000"
)

type BrokerServer struct {
	pb.UnimplementedBrokerServicesServer
}

func choose_number() int32 {
	rand.Seed(time.Now().UTC().UnixNano())
	elec := int32(rand.Intn(3))
	return elec
}

func (s *BrokerServer) AddCity(context.Context, *pb.RequestInf) (*pb.ResponseBroker, error) {
	add := address_1[choose_number()]
	return &pb.ResponseBroker{Address: add}, nil
}

func (s *BrokerServer) UpdateName(context.Context, *pb.RequestInf) (*pb.ResponseBroker, error) {
	add := address_1[choose_number()]
	return &pb.ResponseBroker{Address: add}, nil
}

func (s *BrokerServer) UpdateNumber(context.Context, *pb.RequestInf) (*pb.ResponseBroker, error) {
	add := address_1[choose_number()]
	return &pb.ResponseBroker{Address: add}, nil
}

func (s *BrokerServer) DeleteCity(ctx context.Context, in *pb.RequestDel) (*pb.ResponseBroker, error) {
	add := address_1[choose_number()]
	return &pb.ResponseBroker{Address: add}, nil
}

func (s *BrokerServer) GetNumberRebelds(ctx context.Context, in *pb.RequestLeia) (*pb.ResponseRebelds, error) {
	add := address_1[choose_number()]
	conn, err := grpc.Dial(add, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	Fulcrum := pb.NewFulcrumServicesClient(conn)
	r, err := Fulcrum.GetNumberRebelds(context.Background(), &pb.RequestLeia{Planeta: in.GetPlaneta(), Ciudad: in.GetCiudad()})
	return &pb.ResponseRebelds{Valor: r.Valor, Vector: r.Vector}, nil
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
