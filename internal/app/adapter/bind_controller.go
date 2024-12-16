package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

func BindController(enginer *gin.Engine, daoFactory mysql.IDaoFactory) {
	serverInfoController := CreateTrafficColorController(daoFactory)
	routeServerInfoGroup := enginer.Group("/v1")
	routeServerInfoGroup.GET("/serverinfo", serverInfoController.QueryServerInfos)
	routeServerInfoGroup.POST("/chains", serverInfoController.CreateChain)
	routeServerInfoGroup.PUT("/chains/:chainId/server/:serverId", serverInfoController.BindServerInfo)
	routeServerInfoGroup.DELETE("/chains/:chainId/server/:serverId", serverInfoController.UnbindServerInfo)
}
