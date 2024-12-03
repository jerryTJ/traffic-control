package mysql

import (
	"github.com/jerryTJ/controller/internal/app/dao"
	"gorm.io/gorm"
)

type IDaoFactory interface {
	GetServerInfoDao() dao.ServerDao
}

type DaoFactory struct {
	DB *gorm.DB
}

func (df *DaoFactory) GetServerInfoDao() dao.ServerDao {
	return &dao.ServerDaoImpl{
		DB: df.DB,
	}
}
