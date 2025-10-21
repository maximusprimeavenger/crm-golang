package transport

import (
	"context"
	"errors"
	"fmt"

	proto "github.com/fiveret/product-service/grpc/item-grpc"
	"github.com/fiveret/product-service/internal/helpers"
	"github.com/fiveret/product-service/internal/models"
	"github.com/fiveret/product-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	proto.UnimplementedItemServiceServer
	service service.ItemService
}

func NewGRPCHandler(s service.ItemService) *GRPCHandler {
	return &GRPCHandler{service: s}
}

func (h *GRPCHandler) CreateItem(ctx context.Context, req *proto.CreateItemRequest) (*proto.CreateItemResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending post request: %v", err)
	}
	if req.Item == nil {
		return nil, status.Errorf(codes.InvalidArgument, "item is nil")
	}

	item := &models.Item{
		Name:        &req.Item.Name,
		Category:    &req.Item.Category,
		Price:       &req.Item.Price,
		Description: &req.Item.Description,
		InStock:     &req.Item.InStock,
	}
	createdAt, err := h.service.CreateItem(ctx, item)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error creating new item: %v", err)
	}
	resp := proto.CreateItemResponse{
		Message:   "successfully created!",
		CreatedAt: timestamppb.New(*createdAt),
	}
	return &resp, nil
}

func (h *GRPCHandler) DeleteItem(ctx context.Context, req *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending delete request")
	}
	id := &req.Id
	msg, err := h.service.DeleteItem(ctx, id)
	if err != nil {
		if errors.Is(err, fmt.Errorf("error finding item for deleting")) {
			return &proto.DeleteItemResponse{
				Message: fmt.Sprintf("item with id %d successfully deleted", id),
			}, nil
		}
		return nil, status.Errorf(codes.InvalidArgument, "error deleting an item")
	}
	return &proto.DeleteItemResponse{
		Message: msg,
	}, nil
}

func (h *GRPCHandler) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending get request: %v", err)
	}
	id := &req.Id
	item, err := h.service.GetItem(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error finding an item: %v", err)
	}
	respItem := helpers.ConvertModelsToGRPC(item)
	resp := &proto.GetItemResponse{
		Item: respItem,
	}
	return resp, nil
}

func (h *GRPCHandler) GetItems(ctx context.Context, req *proto.GetItemsRequest) (*proto.GetItemsResponse, error) {
	items, err := h.service.GetItems(ctx)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, status.Errorf(codes.NotFound, "no items found")
	}
	return &proto.GetItemsResponse{
		Items: items,
	}, nil
}

func (h *GRPCHandler) PutItem(ctx context.Context, req *proto.PutItemRequest) (*proto.PutItemResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending put request: %v", err)
	}
	id, item := &req.Id, req.Item
	updatedItem, createdAt, updatedAt, err := h.service.PutItem(ctx, id, helpers.ConvertGRPCToModels(item))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error putting item: %v", err)
	}
	resp := &proto.PutItemResponse{
		CreatedAt: timestamppb.New(*createdAt),
		UpdatedAt: timestamppb.New(*updatedAt),
		Item:      helpers.ConvertModelsToGRPC(updatedItem),
	}
	return resp, nil
}
