package gateway

import (
	"context"
	"strings"

	itemProto "github.com/fiveret/api-gateway/grpc/item-grpc"
	leadProto "github.com/fiveret/api-gateway/grpc/lead-grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, endpoints []string, opts []grpc.DialOption) error {
	for i := 0; i < len(endpoints); i++ {
		parts := strings.Split(endpoints[i], ":")
		switch parts[0] {
		case "lead-service":
			err := leadProto.RegisterLeadServiceHandlerFromEndpoint(ctx, mux, endpoints[i], opts)
			if err != nil {
				return err
			}
			err = leadProto.RegisterLeadProductServiceHandlerFromEndpoint(ctx, mux, endpoints[i], opts)
			if err != nil {
				return err
			}
		case "item-service":
			err := itemProto.RegisterItemServiceHandlerFromEndpoint(ctx, mux, endpoints[i], opts)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
