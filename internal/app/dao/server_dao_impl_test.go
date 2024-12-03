package dao

import (
	"testing"

	"github.com/jerryTJ/controller/internal/app/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&model.ServerInfo{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestServerDaoImpl_Add(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	testServer := &model.ServerInfo{
		Name:    "test-server",
		Version: "1.0.0",
	}

	err := dao.Add(testServer)
	assert.NoError(t, err)
	assert.NotZero(t, testServer.ID)

	// Verify the record was created
	var found model.ServerInfo
	result := db.First(&found, testServer.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, testServer.Name, found.Name)
	assert.Equal(t, testServer.Version, found.Version)
}

func TestServerDaoImpl_Update(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	// First create a record
	testServer := &model.ServerInfo{
		Name:    "test-server",
		Version: "1.0.0",
	}
	err := dao.Add(testServer)
	assert.NoError(t, err)

	// Update the record
	testServer.Name = "updated-server"
	err = dao.Update(testServer)
	assert.NoError(t, err)

	// Verify the update
	var found model.ServerInfo
	result := db.First(&found, testServer.ID)
	assert.NoError(t, result.Error)
	assert.Equal(t, "updated-server", found.Name)
}

func TestServerDaoImpl_Delete(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	// First create a record
	testServer := &model.ServerInfo{
		Name:    "test-server",
		Version: "1.0.0",
	}
	err := dao.Add(testServer)
	assert.NoError(t, err)

	// Delete the record
	err = dao.Delete(testServer)
	assert.NoError(t, err)

	// Verify the deletion
	var found model.ServerInfo
	result := db.First(&found, testServer.ID)
	assert.Error(t, result.Error)
	assert.True(t, result.Error == gorm.ErrRecordNotFound)
}

func TestServerDaoImpl_List(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	// Create test records
	servers := []model.ServerInfo{
		{Name: "server-1", Version: "1.0.0"},
		{Name: "server-2", Version: "1.0.0"},
		{Name: "other-1", Version: "2.0.0"},
	}
	for _, s := range servers {
		err := dao.Add(&s)
		assert.NoError(t, err)
	}

	// Test listing with filter
	found, err := dao.List(&model.ServerInfo{Name: "server"})
	assert.NoError(t, err)
	assert.Len(t, found, 2)
}

func TestServerDaoImpl_ListById(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	// Create a test record
	testServer := &model.ServerInfo{
		Name:    "test-server",
		Version: "1.0.0",
	}
	err := dao.Add(testServer)
	assert.NoError(t, err)

	// Test finding by ID
	found, err := dao.ListById(testServer.ID)
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, testServer.Name, found.Name)
}

func TestServerDaoImpl_QueryByVersion(t *testing.T) {
	db := setupTestDB(t)
	dao := &ServerDaoImpl{DB: db}

	// Create test records
	servers := []model.ServerInfo{
		{Name: "server-1", Version: "1.0.0"},
		{Name: "server-2", Version: "1.0.0"},
		{Name: "server-3", Version: "2.0.0"},
	}
	for _, s := range servers {
		err := dao.Add(&s)
		assert.NoError(t, err)
	}

	// Test querying by name and version
	found, err := dao.QueryByVersion("server", "1.0.0")
	assert.NoError(t, err)
	assert.NotNil(t, found)
	assert.Equal(t, "1.0.0", found.Version)
	assert.Contains(t, found.Name, "server")
}
