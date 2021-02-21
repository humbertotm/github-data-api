package data

import (
	"ghdataapi.htm/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type RepoData interface {
	GetRepo(name string) (*domain.Repo, error)
}

type repoData struct {
	db neo4j.Driver
}

var data RepoData

func NewRepoData(db neo4j.Driver) RepoData {
	if data == nil {
		data = &repoData{db}
	}

	return data
}

func (d *repoData) GetRepo(name string) (*domain.Repo, error) {
	return &domain.Repo{Name: name}, nil
}
