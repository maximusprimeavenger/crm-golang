package clients

import (
	itemProto "github.com/fiveret/api-gateway/grpc/item-grpc"
	proto "github.com/fiveret/api-gateway/grpc/lead-grpc"
	"google.golang.org/grpc"
)

type GRPCClients struct {
	LeadClient        proto.LeadServiceClient
	LeadProductClient proto.LeadProductServiceClient
	ItemClient        itemProto.ItemServiceClient
}

func InitClients() (*GRPCClients, error) {
	conn, err := grpc.NewClient(":50051")
	if err != nil {
		return nil, err
	}
	return &GRPCClients{
			LeadClient:        proto.NewLeadServiceClient(conn),
			LeadProductClient: proto.NewLeadProductServiceClient(conn),
			ItemClient:        itemProto.NewItemServiceClient(conn),
		},
		nil
}
