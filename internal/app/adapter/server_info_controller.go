package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

type ServerInfoController struct {
	factory mysql.IDaoFactory
}

func NewServerInfoController(daoFactory mysql.IDaoFactory) ServerInfoController {
	return ServerInfoController{
		factory: daoFactory,
	}
}
func (ns *ServerInfoController) QueryServerInfos(ctx *gin.Context) {
	name := ctx.Param("name")
	version := ctx.Param("version")
	serverInfoDao := ns.factory.GetServerInfoDao()
	info, err := serverInfoDao.QueryByVersion(name, version)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, info)
}
