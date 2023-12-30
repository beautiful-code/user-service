package user

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUserHandler(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&User{})
	service := NewService(db)

	server := httptest.NewServer(http.HandlerFunc(service.CreateUserHandler))
	defer server.Close()

	payload := []byte(`{"Email":"test@example.com","Password":"password123"}`)
	resp, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
