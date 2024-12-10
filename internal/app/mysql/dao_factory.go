package mysql

import (
	"github.com/jerryTJ/controller/internal/app/dao"
	"gorm.io/gorm"
)

type IDaoFactory interface {
	GetServerInfoDao() dao.ServerDao
	GetChainDao() dao.ChainDao
}

type DaoFactory struct {
	DB *gorm.DB
}

func (df *DaoFactory) GetServerInfoDao() dao.ServerDao {
	return &dao.ServerDaoImpl{
		DB: df.DB,
	}
}
func (df *DaoFactory) GetChainDao() dao.ChainDao {
	return &dao.ChainDaoImpl{
		DB: df.DB,
	}
}

// mock Dao factory to test controller
type MockDaoFactory struct {
	MockServerDao *dao.MockServerDao
	MockChainDao  *dao.ChainDao
}

func (df *MockDaoFactory) GetServerInfoDao() dao.ServerDao {
	return df.MockServerDao
}

func (df *MockDaoFactory) GetChainDao() dao.ChainDao {
	return *df.MockChainDao
}
