package transport

import (
	"context"

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

func (h *GRPCHandler) NewLead(ctx context.Context, req *proto.NewLeadRequest) (*proto.NewLeadResponse, error) {
	lead := req.Lead
	if lead == nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occured during validating, lead is nil")
	}
	msg, createdAt, err := h.leadService.NewLead(ctx, helpers.LeadRequest(lead))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occured during saving the lead: %v", err)
	}
	return &proto.NewLeadResponse{
		Message:   msg,
		CreatedAt: timestamppb.New(*createdAt),
	}, nil
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
	leadResp := helpers.LeadResponse(lead)
	return &proto.GetLeadResponse{
		Lead:      leadResp,
		CreatedAt: timestamppb.New(lead.CreatedAt),
		UpdatedAt: timestamppb.New(lead.UpdatedAt),
	}, nil
}

func (h *GRPCHandler) GetLeads(ctx context.Context, req *proto.GetLeadsRequest) (*proto.GetLeadsResponse, error) {
	leads := h.leadService.GetLeads()

	respLeads := make([]*grpcModels.LeadResponse, 0, len(leads))
	for _, l := range leads {
		respLeads = append(respLeads, helpers.LeadResponse(l))
	}

	return &proto.GetLeadsResponse{Leads: respLeads}, nil
}

func (h *GRPCHandler) PutLead(ctx context.Context, req *proto.PutLeadRequest) (*proto.PutLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	lead := helpers.LeadRequest(req.Lead)
	lead.ID = uint(req.Id)
	updatedLead, err := h.leadService.UpdateLead(lead)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not update lead: %v", err)
	}
	leadResponse := helpers.LeadResponse(updatedLead)

	return &proto.PutLeadResponse{Lead: leadResponse, UpdatedAt: timestamppb.New(updatedLead.UpdatedAt), CreatedAt: timestamppb.New(updatedLead.CreatedAt)}, nil
}

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
