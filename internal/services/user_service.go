package services

import (
	"errors"
	"sync"

	"crm-app/internal/models"
)

type UserService struct {
	mu    sync.Mutex
	users map[int]models.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]models.User),
		nextID: 1,
	}
}

func (s *UserService) CreateUser(name, email, password string) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := models.User{
		ID:       s.nextID,
		Name:     name,
		Email:    email,
		Password: password,
	}
	s.users[s.nextID] = user
	s.nextID++
	return user
}

func (s *UserService) GetUser(id int) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, name, email, password string) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	user.Password = password
	s.users[id] = user
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}