package main

import (
	pb "Lab2/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

const (
	//address = "10.6.40.181:50000"
	address = "localhost:50000"
)

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("BIENVENIDA PRINCESA LEIA")
	fmt.Println("-----------------------------------")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
		fmt.Printf("1.Cantidad de Rebeldes\n2.Cerrar Sesión\n")
		fmt.Scanf("%d \n", &elec)
		fmt.Println("-----------------------------------")
		planeta := ""
		ciudad := ""
		if elec == 1 {

			fmt.Printf("Ingrese Planeta: ")
			fmt.Scanf("%s \n", &planeta)
			fmt.Printf("Ingrese Ciudad: ")
			fmt.Scanf("%s \n", &ciudad)
			r, err := ServiceClient.GetNumberRebelds(context.Background(), &pb.RequestLeia{Planeta: planeta, Ciudad: ciudad})
			if err != nil {
				log.Fatalf("%v", err)
			}
			fmt.Printf("La cantidad de rebeldes es: %d \n", r.Valor)
		}
		if elec == 2 {
			estado = "desconectado"
		}
	}

}
