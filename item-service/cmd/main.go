package main

import (
	"context"
	"log"
	"net"

	proto "github.com/fiveret/product-service/grpc/item-grpc"
	"github.com/fiveret/product-service/internal/db"
	"github.com/fiveret/product-service/internal/models"
	"github.com/fiveret/product-service/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	proto.RegisterItemServiceServer(s, &server{})
	log.Print("server is running on port 50053!")
	if err := s.Serve(lis); err != nil {
		log.Fatal("error serving server:", err)
	}
}

func (s *server) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	dbConn, err := db.Init()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error connecting to db: %v", err)
	}
	createdAt, err := repository.NewItem(&models.Item{
		Name:        &req.Item.Name,
		Category:    &req.Item.Category,
		Price:       &req.Item.Price,
		Description: &req.Item.Description,
		InStock:     &req.Item.InStock,
	}, dbConn)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error creating new item: %v", err)
	}
	resp := proto.CreateItemResponse{
		Message:   "successfully created!",
		CreatedAt: timestamppb.New(*createdAt),
	}
	return &resp, nil
}

func (s *server) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	dbConn, err := db.Init()
	if err != nil {
		return nil, err
	}

}
