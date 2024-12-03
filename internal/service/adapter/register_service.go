package adapter

import (
	"github.com/jerryTJ/controller/internal/app/mysql"
	pb "github.com/jerryTJ/controller/internal/service/grpc"
	"google.golang.org/grpc"
)

type GrpcRegister struct {
	GrpcServer *grpc.Server
	DaoFactory mysql.IDaoFactory
}

func (gr *GrpcRegister) RegisterColoringService() {
	pb.RegisterServerInfoServiceServer(gr.GrpcServer, &ServerInfo{daoFactory: gr.DaoFactory})
}
