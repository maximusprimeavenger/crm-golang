package transport

import (
	"context"
	"fmt"

	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	grpcModels "github.com/fiveret/crm-golang/grpc/models"
	"github.com/fiveret/crm-golang/internal/helpers"
	"github.com/fiveret/crm-golang/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (h *GRPCHandler) CreateLead(ctx context.Context, req *proto.NewLeadRequest) (*proto.NewLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	if req.Lead == nil {
		return nil, status.Error(codes.InvalidArgument, "lead payload is required")
	}

	lead := helpers.LeadGRPCToModels(req.Lead)

	if lead == nil {
		return nil, status.Error(codes.InvalidArgument, "lead conversion failed")
	}
	createdAt, err := h.leadService.NewLead(lead)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create lead: %v", err)
	}

	return &proto.NewLeadResponse{CreatedAt: timestamppb.New(*createdAt)}, nil
}

func (h *GRPCHandler) GetLead(ctx context.Context, req *proto.GetLeadRequest) (*proto.GetLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	id := req.Id
	lead, err := h.leadService.GetLead(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "lead not found: %v", err)
	}
	leadResp := helpers.ModelsToLeadGRPC(lead)
	return &proto.GetLeadResponse{
		Id:        uint32(lead.ID),
		Lead:      leadResp,
		CreatedAt: timestamppb.New(lead.CreatedAt),
		UpdatedAt: timestamppb.New(lead.UpdatedAt),
	}, nil
}

func (h *GRPCHandler) GetLeads(ctx context.Context, req *proto.GetLeadsRequest) (*proto.GetLeadsResponse, error) {
	leads := h.leadService.GetLeads()

	respLeads := make([]*grpcModels.Lead, 0, len(leads))
	for _, l := range leads {
		respLeads = append(respLeads, helpers.ModelsToLeadGRPC(l))
	}

	return &proto.GetLeadsResponse{Leads: respLeads}, nil
}

/*
	func (h *GRPCHandler) UpdateLead(ctx context.Context, req *proto.PutLeadRequest) (*proto.UpdateLeadResponse, error) {
		err := req.Validate()
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
		}


		updatedLead, err := h.leadService.UpdateLead(lead)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not update lead: %v", err)
		}

		return &proto.PutLeadResponse{updatedLead.ID}, nil
	}
*/
func (h *GRPCHandler) DeleteLead(ctx context.Context, req *proto.DeleteLeadRequest) (*proto.DeleteLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	msg, err := h.leadService.DeleteLead(req.Id)
	if err != nil {
		return &proto.DeleteLeadResponse{Message: msg}, status.Errorf(codes.Internal, "could not delete lead: %v", err)
	}

	return &proto.DeleteLeadResponse{Message: msg}, nil
}
