package adapter

import (
	"context"
	"log"

	"github.com/jerryTJ/controller/internal/app/mysql"
	pb "github.com/jerryTJ/controller/internal/service/grpc"
)

type ServerInfo struct {
	pb.UnimplementedServerInfoServiceServer
	dataSource *mysql.DataSource
}

// 实现 SayHello 方法
func (s *ServerInfo) GetColoringInfo(ctx context.Context, req *pb.ServerRequest) (*pb.ServerReply, error) {
	log.Printf("Received: %v", req.GetName())
	info, err := s.dataSource.ServerInfoDao().QueryByVersion(req.GetName(), req.GetVersion())
	if err != nil {
		return &pb.ServerReply{}, nil
	}
	return &pb.ServerReply{Color: info.Color, Name: info.Name, Version: info.Version}, nil
}
