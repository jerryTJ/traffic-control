package adapter

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jerryTJ/controller/internal/app/entity"
	"github.com/jerryTJ/controller/internal/app/model"
	"github.com/jerryTJ/controller/internal/app/mysql"
)

var (
	instance *TraficcColorController
	once     sync.Once
)

// 单例
func CreateTrafficColorController(daoFactory mysql.IDaoFactory) *TraficcColorController {
	once.Do(func() {
		instance = &TraficcColorController{
			factory: daoFactory,
		}
	})
	return instance
}

type TraficcColorController struct {
	factory mysql.IDaoFactory
}

func (ns *TraficcColorController) QueryServerInfos(ctx *gin.Context) {
	name := ctx.Query("name")
	version := ctx.Query("version")
	serverInfoDao := ns.factory.GetServerInfoDao()
	info, err := serverInfoDao.QueryByVersion(name, version)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, info)
}

// add query chain
func (ns *TraficcColorController) CreateChain(ctx *gin.Context) {
	var chainVo entity.ChainVo
	if err := ctx.BindJSON(&chainVo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	chain := &model.Chain{Name: chainVo.Name, Version: chainVo.Version}

	ns.factory.GetChainDao().Add(chain)
	for _, serverVo := range chainVo.ServerInfo {
		serverInfo := model.ServerInfo{Name: serverVo.Name, ID: uint(serverVo.ID), ChainId: chain.ID}
		ns.factory.GetServerInfoDao().Update(&serverInfo)
	}
	// 返回解析后的数据
	ctx.JSON(http.StatusOK, gin.H{
		"message": "save success",
		"data":    chainVo,
	})
}

func (ns *TraficcColorController) UnbindServerInfo(ctx *gin.Context) {
	chainId, err1 := strconv.ParseInt(ctx.Param("chainId"), 10, 10)
	serverId, err2 := strconv.ParseInt(ctx.Param("serverId"), 10, 10)
	if err1 != nil || err2 != nil {
		// 返回解析后的数据
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save success",
			"data":    nil,
		})
		err := ns.factory.GetServerInfoDao().BindChain(uint(chainId), uint(serverId))
		if err != nil {

		}
	}

}

func (ns *TraficcColorController) BindServerInfo(ctx *gin.Context) {
	var defaultChainId uint = 1
	serverId, err := strconv.ParseInt(ctx.Param("serverId"), 10, 10)
	if err != nil {
		// 返回解析后的数据
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "save success",
			"data":    nil,
		})
		err := ns.factory.GetServerInfoDao().BindChain(defaultChainId, uint(serverId))
		if err != nil {

		}
	}
}
