package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/jerryTJ/controller/cmd"
	controller "github.com/jerryTJ/controller/internal/app/adapter"
	service "github.com/jerryTJ/controller/internal/service/adapter"
	pb "github.com/jerryTJ/controller/internal/service/grpc"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
	"github.com/jerryTJ/controller/tools"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cmd.Execute()
	//db.Init(cmd.DB_USER, cmd.DB_PWD, cmd.DB_URL, cmd.DB_NAME)
	//init datasource
	service.Start("/Users/jerry/.kube/config")
	factory := mysql.DaoFactory{DB: mysql.CreateDB()}
	grpcServer := pb.CreateGRPCServer()
	grpcRegister := service.GrpcRegister{GrpcServer: grpcServer, DaoFactory: &factory}
	grpcRegister.RegisterColoringService()
	go pb.StartGrpcServer(grpcServer, cmd.GrpcPort)

	createHttpServer(&factory)
}

func createHttpServer(factory mysql.IDaoFactory) {
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

	// register route
	controller.BindController(router, factory)

	router.Use(Authorize())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Run(":9091")

	//router.RunTLS(":443", "./resources/ssl/3440711_assistant.albk.tech.pem", "./resources/ssl/3440711_assistant.albk.tech.key")
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
