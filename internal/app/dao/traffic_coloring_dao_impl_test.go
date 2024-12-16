package dao

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/jerryTJ/controller/internal/app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info), // 设置日志级别
	})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the schemas
	//db.Debug().Create(&model.ServerInfo{})
	err = db.AutoMigrate(&model.ServerInfo{}, &model.Chain{}, &model.ChainServer{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestServerDaoImpl(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	t.Run("Add", func(t *testing.T) {
		server := &model.ServerInfo{
			Name:    "test-server",
			Version: "1.0",
		}
		err := dao.Add(server)
		assert.NoError(t, err)
		assert.NotZero(t, server.ID)
	})

	t.Run("Update", func(t *testing.T) {
		server := &model.ServerInfo{
			Name:    "test-server-update",
			Version: "1.0",
		}
		err := dao.Add(server)
		assert.NoError(t, err)

		server.Name = "updated-name"
		err = dao.Update(server)
		assert.NoError(t, err)

		found, err := dao.ListById(server.ID)
		assert.NoError(t, err)
		assert.Equal(t, "updated-name", found.Name)
	})

	t.Run("Delete", func(t *testing.T) {
		server := &model.ServerInfo{
			Name:    "test-server",
			Version: "1.0",
		}
		err := dao.Add(server)
		assert.NoError(t, err)
		err = dao.Delete(server)
		assert.NoError(t, err)
		assert.NotZero(t, server.ID)
	})

	t.Run("List", func(t *testing.T) {
		// Clean up
		db.Exec("DELETE FROM t_server_infos")

		server1 := &model.ServerInfo{Name: "test1", Version: "1.0"}
		server2 := &model.ServerInfo{Name: "test2", Version: "1.0"}
		dao.Add(server1)
		dao.Add(server2)

		results, err := dao.GetPaginatedServerInfos(&model.ServerInfo{Name: "test"}, 10, 1)
		assert.NoError(t, err)
		assert.Len(t, results, 2)
	})

	t.Run("QueryByVersion", func(t *testing.T) {
		server := &model.ServerInfo{
			Name:    "version-test",
			Version: "2.0",
		}
		err := dao.Add(server)
		assert.NoError(t, err)

		found, err := dao.QueryByVersion("version", "2.0")
		assert.NoError(t, err)
		assert.Equal(t, server.Name, found.Name)
	})
	t.Run("BindChain", func(t *testing.T) {
		server := &model.ServerInfo{
			Name:    "version-test",
			Version: "2.0",
		}
		err := dao.Add(server)
		assert.NoError(t, err)
		dao.BindChain(2, server.ID)
		server, err = dao.ListById(server.ID)
		assert.NoError(t, err)
		assert.Equal(t, uint(2), server.ChainId)
		dao.UnBindChain(server.ID, 1)
		server, err = dao.ListById(server.ID)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), server.ChainId)
	})
}

func TestChainDaoImpl(t *testing.T) {
	db := setupTestDB(t)
	dao := &ChainDaoImpl{DB: db}
	daoServerInfo := &ServerDaoImpl{DB: db}

	t.Run("Update and QueryById", func(t *testing.T) {
		chain := &model.Chain{
			Name: "test-chain",
		}

		err := dao.Add(chain)
		assert.NoError(t, err)
		err = dao.Update(chain)
		assert.NoError(t, err)
		assert.NotZero(t, chain.ID)

		found, err := dao.QueryById(chain.ID)
		assert.NoError(t, err)
		assert.Equal(t, chain.Name, found.Name)
	})

	t.Run("Query", func(t *testing.T) {
		chain := &model.Chain{
			Name: "query-test-chain",
		}
		err := dao.Add(chain)
		assert.NoError(t, err)
		serverInfo := model.ServerInfo{
			Namespace: "default",
			Name:      "test-query",
		}
		daoServerInfo.Add(&serverInfo)
		var serverInfos = []model.ServerInfo{serverInfo}

		dao.AssociationServerInfo(chain, serverInfos)
		found, err := dao.Query("query-test")
		assert.NoError(t, err)
		assert.Equal(t, 1, len(chain.ServerInfos))
		assert.Equal(t, chain.Name, found.Name)
	})

	t.Run("AssociationServerInfo", func(t *testing.T) {
		chain := model.Chain{
			Name:    "association-test",
			IfClean: false,
		}

		si1 := model.ServerInfo{Namespace: "default", Name: "test1"}
		si1.ID = 10
		daoServerInfo.Add(&si1)
		si2 := model.ServerInfo{Namespace: "default", Name: "test1"}
		si2.ID = 20
		daoServerInfo.Add(&si2)
		serverInfos := []model.ServerInfo{si1, si2}
		chain.ServerInfos = serverInfos
		//err := dao.Add(&chain)
		//assert.NoError(t, err)
		err := dao.AssociationServerInfo(&chain, serverInfos)
		assert.NoError(t, err)

		results := dao.QueryServerInfos(chain.ID)
		assert.NotNil(t, results)
	})
	t.Run("GetPaginatedChains", func(t *testing.T) {
		for i := 0; i < 93; i++ {
			chain := model.Chain{
				Name:    fmt.Sprintf("association-test%d", i),
				Version: fmt.Sprintf("v%d", i),
			}
			err := dao.Add(&chain)
			assert.NoError(t, err)
		}
		paginated := dao.GetPaginatedChains(&model.Chain{Name: "association", Version: "v"}, 1, 20)
		assert.Equal(t, int64(93), paginated.TotalCount)
		assert.Equal(t, 5, paginated.TotalPage)
	})
}
func TestReflect(t *testing.T) {
	t.Run(
		"demo",
		func(t *testing.T) {
			var x int = 42
			tp := reflect.TypeOf(x)
			v := reflect.ValueOf(x)
			fmt.Println("type:", tp)
			fmt.Println("value:", v)

			v = reflect.ValueOf(&x)
			fmt.Printf("%p", &x)
			hexStr := "140001acfb8"
			decValue, _ := strconv.ParseInt(hexStr, 16, 64)
			fmt.Printf("dec---%d", decValue)
			hexStr = fmt.Sprintf("%x", decValue)
			fmt.Printf("10进制 %d 转为16进制: %s\n", decValue, hexStr)

			if v.Kind() == reflect.Pointer {
				v = v.Elem()
			}
			v.SetInt(100)
			fmt.Println(x)

		},
	)
}
