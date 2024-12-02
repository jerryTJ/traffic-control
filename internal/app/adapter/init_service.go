package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

func InitServerController(enginer *gin.Engine, dataSource *mysql.DataSource) {
	serverInfoController := NewServerInfoController(dataSource)
	routeServerInfoGroup := enginer.Group("server")
	routeServerInfoGroup.GET("/:name/:version", serverInfoController.QueryServerInfos)
}
