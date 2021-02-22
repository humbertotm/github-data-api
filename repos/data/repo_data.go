package data

import (
	"fmt"

	"ghdataapi.htm/domain"
	"ghdataapi.htm/log"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const (
	userFields  = "u.username, u.external_id, u.user_url, u.followers_url, u.following_url, u.repos_url, u.type, u.site_admin"
	ownerFields = "o.username, o.external_id, o.user_url, o.followers_url, o.following_url, o.repos_url, o.type, o.site_admin"
	repoFields  = "r.external_id, r.name, r.full_name, r.html_url, r.url, r.contributors_url, r.issues_url, r.languages_url"
)

type RepoData interface {
	GetRepo(name, owner string) (*domain.Repo, error)
	GetRepoContributors(name, owner string) (*domain.Repo, []*domain.User, error)
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

func (d *repoData) GetRepo(name, owner string) (*domain.Repo, error) {
	log.Info.Printf("Retrieving repo %s owned by %s\n", name, owner)
	queryTemplate := `
               MATCH (r:Repo {name: $name})
               MATCH (o:User {username: $username})
               MATCH (o)-[:OWNS]->(r) RETURN %s, %s
        `
	query := fmt.Sprintf(queryTemplate, repoFields, ownerFields)

	session := d.db.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	data, err := session.Run(query, map[string]interface{}{
		"name":     name,
		"username": owner,
	})
	if err != nil {
		return nil, err
	}

	record, err := data.Single()
	if err != nil {
		return nil, err
	}

	return &domain.Repo{
		ExternalID:      int(record.Values[0].(float64)),
		Name:            record.Values[1].(string),
		FullName:        record.Values[2].(string),
		HTMLUrl:         record.Values[3].(string),
		URL:             record.Values[4].(string),
		ContributorsURL: record.Values[5].(string),
		IssuesURL:       record.Values[6].(string),
		LanguagesURL:    record.Values[7].(string),
		Owner: &domain.User{
			Username:     record.Values[8].(string),
			ExternalID:   int(record.Values[9].(float64)),
			UserURL:      record.Values[10].(string),
			FollowersURL: record.Values[11].(string),
			FollowingURL: record.Values[12].(string),
			ReposURL:     record.Values[13].(string),
			Type:         record.Values[14].(string),
			SiteAdmin:    record.Values[15].(bool),
		},
	}, nil
}

func (d *repoData) GetRepoContributors(name, owner string) (*domain.Repo, []*domain.User, error) {
	log.Info.Printf("Retrieving contributors for repo %s owned by %s\n", name, owner)
	queryTemplate := `
                MATCH (r:Repo {name: $repo_name})
                MATCH (o:User {username: $username})
                MATCH (u)-[:CONTRIBUTOR]->(r)
                RETURN %s, %s, %s
        `
	query := fmt.Sprintf(queryTemplate, repoFields, ownerFields, userFields)

	session := d.db.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	data, err := session.Run(query, map[string]interface{}{
		"repo_name": name,
		"username":  owner,
	})
	if err != nil {
		return nil, nil, err
	}

	var contributors []*domain.User
	var repo *domain.Repo
	for data.Next() {
		if repo == nil {
			repo = &domain.Repo{
				ExternalID:      int(data.Record().Values[0].(float64)),
				Name:            data.Record().Values[1].(string),
				FullName:        data.Record().Values[2].(string),
				HTMLUrl:         data.Record().Values[3].(string),
				URL:             data.Record().Values[4].(string),
				ContributorsURL: data.Record().Values[5].(string),
				IssuesURL:       data.Record().Values[6].(string),
				LanguagesURL:    data.Record().Values[7].(string),
				Owner: &domain.User{
					Username:     data.Record().Values[8].(string),
					ExternalID:   int(data.Record().Values[9].(float64)),
					UserURL:      data.Record().Values[10].(string),
					FollowersURL: data.Record().Values[11].(string),
					FollowingURL: data.Record().Values[12].(string),
					ReposURL:     data.Record().Values[13].(string),
					Type:         data.Record().Values[14].(string),
					SiteAdmin:    data.Record().Values[15].(bool),
				},
			}
		}

		contributor := &domain.User{
			Username:     data.Record().Values[16].(string),
			ExternalID:   int(data.Record().Values[17].(float64)),
			UserURL:      data.Record().Values[18].(string),
			FollowersURL: data.Record().Values[19].(string),
			FollowingURL: data.Record().Values[20].(string),
			ReposURL:     data.Record().Values[21].(string),
			Type:         data.Record().Values[22].(string),
			SiteAdmin:    data.Record().Values[23].(bool),
		}
		contributors = append(contributors, contributor)
	}

	return repo, contributors, nil
}
