package main

import (
	pb "Lab2/proto"
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

const (
	port = ":50023"
)

type FulcrumServer struct {
	pb.UnimplementedFulcrumServicesServer
}

type Reloj struct {
	namePlanet string
	x          int
	y          int
	z          int
}

func RegistroLog(Planeta string, accion string, ciudad string, valor int) {
	f, err := os.OpenFile("Fulcrum/"+Planeta+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if valor == 0 || accion == "DeleteCity" {
		log.Printf("%s %s %s", accion, Planeta, ciudad)
	} else {
		log.Printf("%s %s %s %d", accion, Planeta, ciudad, valor)
	}

}

func (s *FulcrumServer) AddCity(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	b := []byte(in.GetPlaneta() + " " + in.GetCiudad() + " " + strconv.Itoa(int(in.GetValor())) + "\n")
	cont, _ := ioutil.ReadFile("Fulcrum/" + in.GetPlaneta() + ".txt")
	cont = append(cont, b...)
	err := ioutil.WriteFile("Fulcrum/"+in.GetPlaneta()+".txt", cont, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "AddCity", in.GetCiudad(), int(in.GetValor()))
	return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) UpdateName(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	Bytes, err := ioutil.ReadFile("Fulcrum/" + in.GetPlaneta() + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	datos := string(Bytes)
	arr := strings.Split(datos, "\n")
	var aux []byte
	for i := 0; i < len(arr); i++ {
		l := strings.Split(arr[i], " ")
		if len(l) > 1 {
			num, _ := strconv.Atoi(l[2])
			if int32(num) == in.GetValor() {
				l[1] = in.GetCiudad()
				arr[i] = strings.Join(l, " ")
			}
			b := []byte(arr[i] + "\n")
			aux = append(aux, b...)

		}
	}
	err = ioutil.WriteFile("Fulcrum/"+in.GetPlaneta()+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "UpdateName", in.GetCiudad(), int(in.GetValor()))
	return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) UpdateNumber(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	Bytes, err := ioutil.ReadFile("Fulcrum/" + in.GetPlaneta() + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	datos := string(Bytes)
	arr := strings.Split(datos, "\n")
	var aux []byte
	for i := 0; i < len(arr); i++ {
		l := strings.Split(arr[i], " ")
		if len(l) > 1 {
			if strings.Compare(l[1], in.GetCiudad()) == 0 {
				l[2] = strconv.Itoa(int(in.GetValor()))
				arr[i] = strings.Join(l, " ")
			}
			b := []byte(arr[i] + "\n")
			aux = append(aux, b...)

		}
	}
	err = ioutil.WriteFile("Fulcrum/"+in.GetPlaneta()+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "UpdateNumber", in.GetCiudad(), int(in.GetValor()))
	return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) DeleteCity(ctx context.Context, in *pb.RequestDel) (*pb.ResponseFulcrum, error) {
	Bytes, err := ioutil.ReadFile("Fulcrum/" + in.GetPlaneta() + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	datos := string(Bytes)
	arr := strings.Split(datos, "\n")
	var aux []byte
	for i := 0; i < len(arr); i++ {
		l := strings.Split(arr[i], " ")
		if len(l) > 1 {
			if strings.Compare(l[1], in.GetCiudad()) != 0 {
				b := []byte(arr[i] + "\n")
				aux = append(aux, b...)
			}
		}
	}
	err = ioutil.WriteFile("Fulcrum/"+in.GetPlaneta()+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "DeleteCity", in.GetCiudad(), 0)
	return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func main() {

	//lista_relojes := []Reloj{}
	//reloj := Reloj{namePlanet: "asd", x: 0, y: 0, z: 0}
	//lista_relojes = append(lista_relojes, reloj)

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
