package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestUserCreation(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&User{})

	user := User{Email: "test@example.com", Password: "password123"}
	result := db.Create(&user)

	assert.NoError(t, result.Error)
	assert.NotEqual(t, 0, user.ID)
}
