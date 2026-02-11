package transport

import (
	"context"
	"fmt"

	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *GRPCHandler) AddProductsToLead(ctx context.Context, req *proto.AddProductsToLeadRequest) (*proto.AddProductsToLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending post request: %v", err)
	}
	if len(req.ProductIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "product ids is null")
	}
	message, err := h.leadProductService.AddProductsToLead(req.Id, req.ProductIds)
	if err != nil {
		return &proto.AddProductsToLeadResponse{Message: message}, err
	}
	return &proto.AddProductsToLeadResponse{Message: message}, nil
}

func (h *GRPCHandler) DeleteLeadProducts(ctx context.Context, req *proto.DeleteLeadProductsRequest) (*proto.DeleteLeadProductsResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending delete request: %v", err)
	}
	msg, err := h.leadProductService.DeleteLeadProducts(req.Id)
	if err != nil {
		return &proto.DeleteLeadProductsResponse{Message: msg}, status.Errorf(codes.InvalidArgument, "couldn't delete lead's products: %v", err)
	}
	return &proto.DeleteLeadProductsResponse{Message: msg}, nil
}

func (h *GRPCHandler) DeleteLeadProduct(ctx context.Context, req *proto.DeleteLeadProductRequest) (*proto.DeleteLeadProductResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending delete request: %v", err)
	}
	msg, err := h.leadProductService.DeleteLeadProduct(req.Id, req.ProductId)
	if err != nil {
		return &proto.DeleteLeadProductResponse{Message: msg}, status.Error(codes.InvalidArgument, fmt.Sprint(err))
	}
	return &proto.DeleteLeadProductResponse{Message: msg}, nil
}
