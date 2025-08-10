package transport

import (
	"context"

	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	"github.com/fiveret/crm-golang/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	proto.UnimplementedLeadServiceServer
	proto.UnimplementedLeadProductServiceServer
	leadService        service.LeadService
	leadProductService service.LeadProductService
}

func NewGRPCHandler(serv1 service.LeadService, serv2 service.LeadProductService) *GRPCHandler {
	return &GRPCHandler{leadService: serv1, leadProductService: serv2}
}

func (h *GRPCHandler) AddProductsToLead(ctx context.Context, req *proto.AddProductsToLeadRequest) (*proto.AddProductsToLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error sending post request: %v", err)
	}
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is null")
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
