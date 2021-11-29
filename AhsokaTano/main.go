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
	fmt.Println("BIENVENIDA AHSOKA TANO")
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
			fmt.Printf("Ingrese catidad de reveldes (0 en caso de no especificar cantidad): ")
			fmt.Scanf("%d \n", &valor)
			if elec == 1 {
				r, err := ServiceClient.AddCity(context.Background(), &pb.RequestInf{Planeta: planeta, Ciudad: ciudad, Valor: int32(valor)})
				if err != nil {
					log.Fatalf("%v", err)
				}
				address_server := r.Address
				fmt.Println(address_server)
			}
			if elec == 2 {
				print("Modificar Ciudad")
			}
			if elec == 3 {
				print("Modificar Valor")
			}

		}

		if elec == 4 {
			fmt.Printf("Ingrese Planeta: ")
			fmt.Scanf("%s \n", &planeta)
			fmt.Printf("Ingrese Ciudad: ")
			fmt.Scanf("%s \n", &ciudad)
		}
		if elec == 5 {
			estado = "desconectado"
		}

	}
	fmt.Println("-----------------------------------")
	fmt.Println("DESCONECTADA")
	fmt.Println("-----------------------------------")
}
