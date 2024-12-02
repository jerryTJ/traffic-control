package adapter

import (
	"context"

	"github.com/jerryTJ/controller/internal/service/grpc"
)

type IServer interface {
	GetColoringInfo(ctx context.Context, req *grpc.ServerRequest) (*grpc.ServerReply, error)
}
