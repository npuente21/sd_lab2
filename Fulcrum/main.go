package main

import (
	pb "Lab2/proto"
	"context"
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

const (
	port = ":40000"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServicesServer
}

func (s *FulcrumServer) AddCity(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	b := []byte(in.GetPlaneta() + " " + in.GetCiudad() + " " + strconv.Itoa(int(in.GetValor())) + "\n")
	cont, _ := ioutil.ReadFile("Fulcrum/" + in.GetPlaneta() + ".txt")
	cont = append(cont, b...)
	err := ioutil.WriteFile("Fulcrum/"+in.GetPlaneta()+".txt", cont, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func main() {

	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFulcrumServicesServer(grpcServer, &FulcrumServer{})
	log.Printf("server listening at %v", listner.Addr())

	if err = grpcServer.Serve(listner); err != nil {
		log.Fatalf("Failed to listen on port 50023: %v", err)
	}
}
