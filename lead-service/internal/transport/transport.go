package transport

import (
	proto "github.com/fiveret/crm-golang/grpc/lead-grpc"
	"github.com/fiveret/crm-golang/internal/service"
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
