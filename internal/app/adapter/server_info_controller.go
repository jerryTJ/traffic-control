package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

type ServerInfoController struct {
	dataSource *mysql.DataSource
}

func NewServerInfoController(dataSource *mysql.DataSource) ServerInfoController {
	return ServerInfoController{
		dataSource: dataSource,
	}

}
func (ns *ServerInfoController) QueryServerInfos(ctx *gin.Context) {
	name := ctx.Param("name")
	version := ctx.Param("version")
	ns.dataSource.ServerInfoDao().QueryByVersion(name, version)
}
