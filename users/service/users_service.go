package service

import (
	"ghdataapi.htm/domain"
	"ghdataapi.htm/users/data"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// UsersService defines the interface for the users service layer
type UsersService interface {
	GetUser(username string) (*domain.User, error)
	GetUserFollowers(username string) (*domain.User, []*domain.User, error)
	GetUserFollowing(username string) (*domain.User, []*domain.User, error)
}

type usersService struct {
	userData data.UserData
}

var service UsersService

// NewUsersService returns an instance of UsersService
func NewUsersService(db neo4j.Driver) UsersService {
	if service == nil {
		service = &usersService{
			userData: data.NewUserData(db),
		}
	}

	return service
}

// GetUser ...
func (s *usersService) GetUser(username string) (*domain.User, error) {
	return s.userData.GetUser(username)
}

// GetUserFollowers ...
func (s *usersService) GetUserFollowers(username string) (*domain.User, []*domain.User, error) {
	return s.userData.GetUserFollowers(username)
}

// GetUserFollowing ...
func (s *usersService) GetUserFollowing(username string) (*domain.User, []*domain.User, error) {
	return s.userData.GetUserFollowing(username)
}
