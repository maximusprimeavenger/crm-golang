package main

import (
	"log"
	"net"

	proto "github.com/fiveret/product-service/grpc/item-grpc"
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/repository"
	"github.com/fiveret/product-service/internal/service"
	"github.com/fiveret/product-service/internal/transport"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatal("error listening on port 50053:", err)
	}
	s := grpc.NewServer()

	dbConn, err := db.Init()
	if err != nil {
		log.Fatal("error while connecting to db")
	}
	repo := repository.NewItemRepo(dbConn)
	svc := service.NewItemService(repo)
	handler := transport.NewGRPCHandler(svc)

	proto.RegisterItemServiceServer(s, handler)

	log.Print("server is running on port 50053!")
	if err := s.Serve(lis); err != nil {
		log.Fatal("error serving server:", err)
	}
}
