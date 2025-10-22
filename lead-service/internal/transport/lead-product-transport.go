package transport

import (
	"context"

	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	"github.com/fiveret/crm-golang/internal/helpers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *GRPCHandler) NewLead(ctx context.Context, req *proto.NewLeadRequest) (*proto.NewLeadResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occured during validating")
	}
	lead := req.Lead
	createdAt, err := h.leadService.NewLead(helpers.LeadGRPCToModels(lead))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error occured during saving the lead")
	}
	return &proto.NewLeadResponse{
		Lead:      lead,
		CreatedAt: timestamppb.New(*createdAt),
	}, nil
}
