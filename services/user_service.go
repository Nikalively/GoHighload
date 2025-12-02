package services

import (
	"gohighload/models"
	"sync"
)

type UserService struct {
	users  map[int]models.User
	mu     sync.RWMutex
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (s *UserService) Create(user models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user
}

func (s *UserService) GetAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []models.User
	for _, user := range s.users {
		result = append(result, user)
	}
	return result
}

func (s *UserService) GetByID(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, exists := s.users[id]
	return user, exists
}

func (s *UserService) Update(id int, updatedUser models.User) (models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[id]; !exists {
		return models.User{}, false
	}
	updatedUser.ID = id
	s.users[id] = updatedUser
	return updatedUser, true
}

func (s *UserService) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.users[id]; !exists {
		return false
	}
	delete(s.users, id)
	return true
}
