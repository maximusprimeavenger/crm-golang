package gateway

import (
	"context"

	proto "github.com/fiveret/api-gateway/grpc/item-grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, endpoints []string, opts []grpc.DialOption) error {
	for i := 0; i < len(endpoints); i++ {
		err := proto.RegisterItemServiceHandlerFromEndpoint(ctx, mux, endpoints[i], opts)
		if err != nil {
			return err
		}
	}
	return nil
}
