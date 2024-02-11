package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)


type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"` 
}

type Storage struct {
	users map[int]User
	nextID int 
	mutex sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[int]User),
		nextID: 1,
	}
}

func (s *Storage) GetUsers() []User {
	s.mutex.RLock()
	defet s.mutext.RUnlock()

	var result []User
	
	for _, user := range s.users {
		result = append(result, user)
	}
	return result 
}


func (s *Storage) GetUserById(id int) (User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, ok := s.users[id]
	return user, ok
}

func (s *Storage) CreateUser(name, email string) User {
	s.mutex.Lock()
	defer s.mutex.Unlock( )

}


