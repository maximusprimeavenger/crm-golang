package main

import (
	"log"
	"net"

	proto "github.com/fiveret/product-service/grpc/item-grpc"
	"github.com/fiveret/product-service/internal/db"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedItemServiceServer
}

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

	proto.RegisterItemServiceServer(s, &server{})
	log.Print("server is running on port 50053!")
	if err := s.Serve(lis); err != nil {
		log.Fatal("error serving server:", err)
	}
}
