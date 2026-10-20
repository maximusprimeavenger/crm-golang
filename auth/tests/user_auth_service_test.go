package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockCollection struct{}

func NewMongoRepository(collection *MockCollection) *MongoRepository {
	return &MongoRepository{collection: collection}
}

type MongoRepository struct {
	collection *MockCollection
}

func (repo *MongoRepository) FindByID(id string) (*Item, error) {
	return &Item{ID: id}, nil
}

type Item struct {
	ID string
}

func GetItemHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "name": "Item 1"})
}

func TestGetItem(t *testing.T) {
	router := gin.Default()
	router.GET("/items/:id", GetItemHandler)

	req, _ := http.NewRequest("GET", "/items/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Item 1")
}

func TestFindItemByID(t *testing.T) {
	mockCollection := &MockCollection{}
	repo := NewMongoRepository(mockCollection)

	item, err := repo.FindByID("testID")
	assert.NoError(t, err)
	assert.Equal(t, "testID", item.ID)
}
