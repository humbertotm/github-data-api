package service

import (
	"ghdataapi.htm/domain"
	"ghdataapi.htm/users/data"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UsersService interface {
	GetUser(username string) (*domain.User, error)
	GetUserFollowers(username, maxCount string) ([]*domain.User, error)
	GetUserFollowing(username, maxCount string) ([]*domain.User, error)
}

type usersService struct {
	userData data.UserData
}

var service UsersService

func NewUsersService(db neo4j.Driver) UsersService {
	if service == nil {
		service = &usersService{
			userData: data.NewUserData(db),
		}
	}

	return service
}

func (s *usersService) GetUser(username string) (*domain.User, error) {
	return s.userData.GetUser(username)
}

func (s *usersService) GetUserFollowers(username, maxCount string) ([]*domain.User, error) {
	return s.userData.GetUserFollowers(username, maxCount)
}

func (s *usersService) GetUserFollowing(username, maxCount string) ([]*domain.User, error) {
	return s.userData.GetUserFollowing(username, maxCount)
}
