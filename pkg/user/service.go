package user

import (
	"fmt"
	"log"
	"regexp"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(email, password string) (*User, error) {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if regex.MatchString(email) {
		return nil, fmt.Errorf("Invalid email format")
	}
	if len(password) <= 6 {
		return nil, fmt.Errorf("Invalid password")
	}
	user := User{Email: email, Password: password}
	result := s.db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &user, nil
}
