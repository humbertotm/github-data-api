package service

import (
	"ghdataapi.htm/domain"
	"ghdataapi.htm/repos/data"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// ReposService defines the interface for the repo service layer
type ReposService interface {
	GetRepo(name, owner string) (*domain.Repo, error)
	GetRepoContributors(name, owner string) (*domain.Repo, []*domain.User, error)
}

type reposService struct {
	repoData data.RepoData
}

var service ReposService

// NewReposService returns an instance of ReposService
func NewReposService(db neo4j.Driver) ReposService {
	if service == nil {
		service = &reposService{
			repoData: data.NewRepoData(db),
		}
	}

	return service
}

// GetRepo ...
func (s *reposService) GetRepo(name, owner string) (*domain.Repo, error) {
	return s.repoData.GetRepo(name, owner)
}

// GetRepoContributors ...
func (s *reposService) GetRepoContributors(name, owner string) (*domain.Repo, []*domain.User, error) {
	return s.repoData.GetRepoContributors(name, owner)
}
