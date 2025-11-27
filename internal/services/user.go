package services

import (
	"test-1/internal/dtos"
	"test-1/internal/entities"
	"test-1/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func Constructor(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(body dtos.CreateUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hashedPassword),
	}
	
	return u.userRepository.CreateUser(user)
}
