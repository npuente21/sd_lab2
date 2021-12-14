package main

import (
	pb "Lab2/proto"
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
)

const (
	address = "10.6.40.183:50000"
	//address = "localhost:50000"
)

var informante = []string{"AHSOKA TANO", "ALMIRANTE THRAWN"}

func main() {
	choice := 0
	fmt.Println("-----------------------------------")
	fmt.Println("IDENTIFIQUESE")
	fmt.Println("1.AHOSOKA TANO\n2.ALMIRANTE THRAWN")
	fmt.Scanf("%d \n", &choice)
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Printf("BIENVENIDA/O %s \n", informante[choice-1])
	fmt.Println("-----------------------------------")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()) //conexion informante-broker

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	ServiceClient := pb.NewBrokerServicesClient(conn)

	estado := "conectada"
	elec := 0
	for estado == "conectada" {
		fmt.Println("-----------------------------------")
		fmt.Println("Estado: Conectada a la red")
		fmt.Println("-----------------------------------")
		fmt.Printf("1.AddCity \n2.UpdateName\n3.UpdateNumber\n4.DeleteCity\n5.Cerrar Sesión\n")
		fmt.Scanf("%d \n", &elec)
		fmt.Println("-----------------------------------")
		planeta := ""
		ciudad := ""
		valor := 0
		if elec == 1 || elec == 2 || elec == 3 {

			fmt.Printf("Ingrese Planeta: ")
			fmt.Scanf("%s \n", &planeta)
			fmt.Printf("Ingrese Ciudad: ")
			fmt.Scanf("%s \n", &ciudad)
			fmt.Printf("Ingrese catidad de rebeldes (0 en caso de no especificar cantidad): ")
			fmt.Scanf("%d \n", &valor)
			if elec == 1 { //addcity
				r, err := ServiceClient.AddCity(context.Background(), &pb.RequestInf{Planeta: planeta, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				address_server := r.Address
				//fmt.Println(address_server)
				ip := strings.Split(address_server, ":")[0] //info para modificar reloj en x, y o z
				ip = string(ip[len(ip)-1])                  //same as above
				conn, err := grpc.Dial(address_server, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("Did not connect: %v", err)
				}
				defer conn.Close()
				ServiceFulcrum := pb.NewFulcrumServicesClient(conn)
				r1, err := ServiceFulcrum.AddCity(context.Background(), &pb.RequestInf{Planeta: planeta + ip, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				if r1 != nil {
					fmt.Printf("Ciudad añadida \n")
				}

			}
			if elec == 2 { //updatename
				r, err := ServiceClient.UpdateName(context.Background(), &pb.RequestInf{Planeta: planeta, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				address_server := r.Address
				ip := strings.Split(address_server, ":")[0] //info para modificar reloj en x, y o z
				ip = string(ip[len(ip)-1])                  //same as above
				fmt.Println(address_server)
				conn, err := grpc.Dial(address_server, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("Did not connect: %v", err)
				}
				defer conn.Close()
				ServiceFulcrum := pb.NewFulcrumServicesClient(conn)
				r1, err := ServiceFulcrum.UpdateName(context.Background(), &pb.RequestInf{Planeta: planeta + ip, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				if r1 != nil {
					fmt.Printf("Nombre actualizado \n")
				}
			}
			if elec == 3 { //updatenumber
				r, err := ServiceClient.UpdateNumber(context.Background(), &pb.RequestInf{Planeta: planeta, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				address_server := r.Address
				ip := strings.Split(address_server, ":")[0] //info para modificar reloj en x, y o z
				ip = string(ip[len(ip)-1])                  //same as above
				fmt.Println(address_server)
				conn, err := grpc.Dial(address_server, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("Did not connect: %v", err)
				}
				defer conn.Close()
				ServiceFulcrum := pb.NewFulcrumServicesClient(conn)
				r1, err := ServiceFulcrum.UpdateNumber(context.Background(), &pb.RequestInf{Planeta: planeta + ip, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				if r1 != nil {
					fmt.Printf("Modificación realizada \n")
				}
			}

		}

		if elec == 4 {
			fmt.Printf("Ingrese Planeta: ")
			fmt.Scanf("%s \n", &planeta)
			fmt.Printf("Ingrese Ciudad: ")
			fmt.Scanf("%s \n", &ciudad)
			r, err := ServiceClient.DeleteCity(context.Background(), &pb.RequestDel{Planeta: planeta, Ciudad: ciudad})
			if err != nil {
				log.Fatalf("%v", err)
			}
			address_server := r.Address
			fmt.Println(address_server)
			ip := strings.Split(address_server, ":")[0] //info para modificar reloj en x, y o z
			ip = string(ip[len(ip)-1])                  //same as above
			conn, err := grpc.Dial(address_server, grpc.WithInsecure(), grpc.WithBlock())
			if err != nil {
				log.Fatalf("Did not connect: %v", err)
			}
			defer conn.Close()
			ServiceFulcrum := pb.NewFulcrumServicesClient(conn)
			r1, err := ServiceFulcrum.DeleteCity(context.Background(), &pb.RequestDel{Planeta: planeta + ip, Ciudad: ciudad})
			if err != nil {
				log.Fatalf("%v", err)
			}
			if r1 != nil {
				fmt.Printf("Ciudad eliminada \n")
			}
		}
		if elec == 5 {
			estado = "desconectado"
		}

	}
	fmt.Println("-----------------------------------")
	fmt.Println("DESCONECTADA/O")
	fmt.Println("-----------------------------------")
}
