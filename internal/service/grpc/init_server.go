package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func CreateGRPCServer() *grpc.Server {
	// 加载服务端证书和私钥
	serverCert, err := tls.LoadX509KeyPair("internal/service/ssl/server.crt", "internal/service/ssl/server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	// 加载 CA 证书
	caCert, err := os.ReadFile("internal/service/ssl/ca.crt")
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		log.Fatalf("Failed to append CA certificate")
	}

	// 配置 TLS，要求客户端提供证书（双向 TLS）
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	creds := credentials.NewTLS(tlsConfig)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	return grpcServer
}
func StartGrpcServer(grpcServer *grpc.Server, addr string) {

	// 创建 gRPC 服务端，启用双向 TLS

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server is listening on port " + addr)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
