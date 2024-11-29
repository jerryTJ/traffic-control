package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/cmd"
	"github.com/jerryTJ/controller/init/db"
	"github.com/jerryTJ/controller/tools"
	pb "github.com/jerryTJ/controller/web/app"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCoroingServiceServer
}

// 实现 SayHello 方法
func (s *server) GetColoringInfo(ctx context.Context, req *pb.ServerRequest) (*pb.ServerReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.ServerReply{Color: "red", Name: req.Name}, nil
}

func startGRPCServer() {
	// 加载服务端证书和私钥
	serverCert, err := tls.LoadX509KeyPair("configs/ssl/server.crt", "configs/ssl/server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificate: %v", err)
	}

	// 加载 CA 证书
	caCert, err := os.ReadFile("configs/ssl/ca.crt")
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

	// 创建 gRPC 服务端，启用双向 TLS
	creds := credentials.NewTLS(tlsConfig)
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterCoroingServiceServer(grpcServer, &server{})

	listener, err := net.Listen("tcp", ":10080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server is listening on port 10080...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func main() {
	cmd.Execute()
	db.Init(cmd.DB_USER, cmd.DB_PWD, cmd.DB_URL, cmd.DB_NAME)
	go startGRPCServer()
	router()
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now() // 记录开始时间
		// 处理请求
		c.Next()
		// 计算延迟
		duration := time.Since(startTime).Seconds()
		// 更新指标
		tools.HttpRequestsTotal.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
			http.StatusText(c.Writer.Status()),
		).Inc()
		tools.HttpRequestDuration.WithLabelValues(
			c.Request.Method,
			c.FullPath(),
		).Observe(duration)
	}
}

func router() {
	router := gin.Default()
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	// write log
	f, _ := os.Create("./coloring_controller.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(PrometheusMiddleware())
	router.Use(gin.Recovery())
	store := memstore.NewStore([]byte("1"))
	router.Use(sessions.Sessions("session", store))
	router.Use(Authorize())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":9091")

	//router.RunTLS(":443", "./resources/ssl/3440711_assistant.albk.tech.pem", "./resources/ssl/3440711_assistant.albk.tech.key")

}
