package gateway

import (
	"context"
	"os"

	itemProto "github.com/fiveret/api-gateway/grpc/item-grpc"
	leadProto "github.com/fiveret/api-gateway/grpc/lead-grpc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RegisterHandlers(ctx context.Context, mux *runtime.ServeMux, opts []grpc.DialOption) error {
	if err := leadProto.RegisterLeadServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("LEAD_SERVICE_URL"),
		opts,
	); err != nil {
		return err
	}

	if err := leadProto.RegisterLeadProductServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("LEAD_SERVICE_URL"),
		opts,
	); err != nil {
		return err
	}

	if err := itemProto.RegisterItemServiceHandlerFromEndpoint(
		ctx,
		mux,
		os.Getenv("ITEM_SERVICE_URL"),
		opts,
	); err != nil {
		return err
	}

	return nil
}
