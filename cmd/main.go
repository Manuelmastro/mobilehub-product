package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Manuelmastro/mobilehub-product/pkg/config"
	"github.com/Manuelmastro/mobilehub-product/pkg/db"
	pb "github.com/Manuelmastro/mobilehub-product/pkg/pb"
	services "github.com/Manuelmastro/mobilehub-product/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", c.Port)

	s := services.ProductServiceServer{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
