package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

func InitServerController(enginer *gin.Engine, daoFactory mysql.IDaoFactory) {
	serverInfoController := NewServerInfoController(daoFactory)
	routeServerInfoGroup := enginer.Group("server")
	routeServerInfoGroup.GET("/:name/:version", serverInfoController.QueryServerInfos)
}
