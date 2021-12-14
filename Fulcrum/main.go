package main

import (
	pb "Lab2/proto"
	"context"
	"fmt"
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
	x          int32
	y          int32
	z          int32
}

var lista_relojes = []Reloj{}

func encontrarVector(Planeta string) Reloj {
	for _, reloj := range lista_relojes {
		if reloj.namePlanet == Planeta {
			return reloj
		}
	}
	return Reloj{namePlanet: "notfound", x: -1, y: -1, z: -1}
}

func actualizarReloj(Planeta string, ip string) {
	var flag int32 = 0
	fmt.Println("--actreloj--")
	fmt.Println(Planeta)
	fmt.Println(ip)
	fmt.Println("----")
	for _, reloj := range lista_relojes {
		if reloj.namePlanet == Planeta {
			flag = 1
			if ip == "1" {
				reloj.x = reloj.x + 1
			}
			if ip == "2" {
				reloj.y = reloj.y + 1
			}
			if ip == "4" {
				reloj.z = reloj.z + 1
			}
		}
	}
	//planeta no estaba en la lista
	if flag == 0 {
		if ip == "1" {
			reloj := Reloj{namePlanet: Planeta, x: 1, y: 0, z: 0}
			lista_relojes = append(lista_relojes, reloj)
		}
		if ip == "2" {
			reloj := Reloj{namePlanet: Planeta, x: 0, y: 1, z: 0}
			lista_relojes = append(lista_relojes, reloj)
		}
		if ip == "4" {
			reloj := Reloj{namePlanet: Planeta, x: 0, y: 0, z: 1}
			lista_relojes = append(lista_relojes, reloj)
		}

	}
	//probando local
	if ip == "t" {
		reloj := Reloj{namePlanet: Planeta, x: -1, y: -1, z: -1}
		lista_relojes = append(lista_relojes, reloj)
	}
	fmt.Println(lista_relojes)
}

func RegistroLog(Planeta string, accion string, ciudad string, valor int) {
	ip := string(Planeta[len(Planeta)-1])
	Planeta = Planeta[:len(Planeta)-1]
	//fmt.Println("PLANETA " + Planeta)

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

	actualizarReloj(Planeta, ip)
	//fmt.Println(lista_relojes)
}

func (s *FulcrumServer) AddCity(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	inGetPlaneta := in.GetPlaneta()[:len(in.GetPlaneta())-1]
	b := []byte(inGetPlaneta + " " + in.GetCiudad() + " " + strconv.Itoa(int(in.GetValor())) + "\n")
	cont, _ := ioutil.ReadFile("Fulcrum/" + inGetPlaneta + ".txt")
	cont = append(cont, b...)
	err := ioutil.WriteFile("Fulcrum/"+inGetPlaneta+".txt", cont, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "AddCity", in.GetCiudad(), int(in.GetValor()))
	reloj_aux := encontrarVector(inGetPlaneta)
	vector := "(" + strconv.Itoa(int(reloj_aux.x)) + strconv.Itoa(int(reloj_aux.y)) + strconv.Itoa(int(reloj_aux.z)) + ")"
	return &pb.ResponseFulcrum{Vector: vector}, nil
	//return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) UpdateName(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	inGetPlaneta := in.GetPlaneta()[:len(in.GetPlaneta())-1]
	Bytes, err := ioutil.ReadFile("Fulcrum/" + inGetPlaneta + ".txt")
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
	err = ioutil.WriteFile("Fulcrum/"+inGetPlaneta+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "UpdateName", in.GetCiudad(), int(in.GetValor()))
	reloj_aux := encontrarVector(inGetPlaneta)
	vector := "(" + strconv.Itoa(int(reloj_aux.x)) + ", " + strconv.Itoa(int(reloj_aux.y)) + ", " + strconv.Itoa(int(reloj_aux.z)) + ")"
	return &pb.ResponseFulcrum{Vector: vector}, nil
	//return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) UpdateNumber(ctx context.Context, in *pb.RequestInf) (*pb.ResponseFulcrum, error) {
	inGetPlaneta := in.GetPlaneta()[:len(in.GetPlaneta())-1]
	Bytes, err := ioutil.ReadFile("Fulcrum/" + inGetPlaneta + ".txt")
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
	err = ioutil.WriteFile("Fulcrum/"+inGetPlaneta+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "UpdateNumber", in.GetCiudad(), int(in.GetValor()))
	reloj_aux := encontrarVector(inGetPlaneta)
	vector := "(" + strconv.Itoa(int(reloj_aux.x)) + ", " + strconv.Itoa(int(reloj_aux.y)) + ", " + strconv.Itoa(int(reloj_aux.z)) + ")"
	return &pb.ResponseFulcrum{Vector: vector}, nil
	//return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) DeleteCity(ctx context.Context, in *pb.RequestDel) (*pb.ResponseFulcrum, error) {
	inGetPlaneta := in.GetPlaneta()[:len(in.GetPlaneta())-1]
	Bytes, err := ioutil.ReadFile("Fulcrum/" + inGetPlaneta + ".txt")
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
	err = ioutil.WriteFile("Fulcrum/"+inGetPlaneta+".txt", aux, 0644)
	if err != nil {
		log.Fatalf("Failed to write in File")
	}
	RegistroLog(in.GetPlaneta(), "DeleteCity", in.GetCiudad(), 0)
	reloj_aux := encontrarVector(inGetPlaneta)
	vector := "(" + strconv.Itoa(int(reloj_aux.x)) + ", " + strconv.Itoa(int(reloj_aux.y)) + ", " + strconv.Itoa(int(reloj_aux.z)) + ")"
	return &pb.ResponseFulcrum{Vector: vector}, nil
	//return &pb.ResponseFulcrum{Vector: "OK"}, nil
}

func (s *FulcrumServer) GetNumberRebelds(ctx context.Context, in *pb.RequestLeia) (*pb.ResponseRebelds, error) {
	inGetPlaneta := in.GetPlaneta()[:len(in.GetPlaneta())-1]
	Bytes, err := ioutil.ReadFile("Fulcrum/" + inGetPlaneta + ".txt")
	if err != nil {
		log.Fatal(err)
	}

	datos := string(Bytes)
	arr := strings.Split(datos, "\n")
	for i := 0; i < len(arr); i++ {
		l := strings.Split(arr[i], " ")
		if len(l) > 1 {
			if strings.Compare(l[1], in.GetCiudad()) == 0 {
				num, _ := strconv.Atoi(l[2])
				return &pb.ResponseRebelds{Valor: int32(num), Vector: "OK"}, nil
			}

		}
	}
	return &pb.ResponseRebelds{Vector: "OK"}, nil
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
