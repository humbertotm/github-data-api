package service

import (
	"ghdataapi.htm/domain"
	"ghdataapi.htm/repos/data"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ReposService interface {
	GetRepo(name string) (*domain.Repo, error)
}

type reposService struct {
	repoData data.RepoData
}

var service ReposService

func NewReposService(db neo4j.Driver) ReposService {
	if service == nil {
		service = &reposService{
			repoData: data.NewRepoData(db),
		}
	}

	return service
}

func (s *reposService) GetRepo(name string) (*domain.Repo, error) {
	return s.repoData.GetRepo(name)
}
