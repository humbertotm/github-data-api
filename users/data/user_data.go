package data

import (
	"fmt"

	"ghdataapi.htm/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

const userFields = "u.username, u.external_id, u.user_url, u.followers_url, u.following_url, u.repos_url, u.type, u.site_admin"

type UserData interface {
	GetUser(username string) (*domain.User, error)
	GetUserFollowers(username, maxCount string) ([]*domain.User, error)
	GetUserFollowing(username, maxCount string) ([]*domain.User, error)
}

type userData struct {
	db neo4j.Driver
}

var data UserData

func NewUserData(db neo4j.Driver) UserData {
	if data == nil {
		data = &userData{db}
	}

	return data
}

func (d *userData) GetUser(username string) (*domain.User, error) {
	query := fmt.Sprintf("MATCH (u:User {username: $username}) RETURN %s", userFields)

	session := d.db.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	data, err := session.Run(query, map[string]interface{}{"username": username})
	if err != nil {
		return nil, err
	}

	record, err := data.Single()
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Username:     record.Values[0].(string),
		ExternalID:   int(record.Values[1].(float64)),
		UserURL:      record.Values[2].(string),
		FollowersURL: record.Values[3].(string),
		FollowingURL: record.Values[4].(string),
		ReposURL:     record.Values[5].(string),
		Type:         record.Values[6].(string),
		SiteAdmin:    record.Values[7].(bool),
	}, nil
}

func (d *userData) GetUserFollowers(username, maxCount string) ([]*domain.User, error) {
	queryTemplate := `
                MATCH (f:User {username: $username})
                MATCH (u)-[:FOLLOWS]->(f)
                RETURN %s
                LIMIT %s
        `
	query := fmt.Sprintf(queryTemplate, userFields, maxCount)

	session := d.db.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	data, err := session.Run(query, map[string]interface{}{"username": username})
	if err != nil {
		return nil, err
	}

	var followers []*domain.User
	for data.Next() {
		follower := &domain.User{
			Username:     data.Record().Values[0].(string),
			ExternalID:   int(data.Record().Values[1].(float64)),
			UserURL:      data.Record().Values[2].(string),
			FollowersURL: data.Record().Values[3].(string),
			FollowingURL: data.Record().Values[4].(string),
			ReposURL:     data.Record().Values[5].(string),
			Type:         data.Record().Values[6].(string),
			SiteAdmin:    data.Record().Values[7].(bool),
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

func (d *userData) GetUserFollowing(username, maxCount string) ([]*domain.User, error) {
	queryTemplate := `
                MATCH (f:User {username: $username})
                MATCH (f)-[:FOLLOWS]->(u)
                RETURN %s
                LIMIT %s
        `
	query := fmt.Sprintf(queryTemplate, userFields, maxCount)

	session := d.db.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	data, err := session.Run(query, map[string]interface{}{"username": username})
	if err != nil {
		return nil, err
	}

	var followers []*domain.User
	for data.Next() {
		follower := &domain.User{
			Username:     data.Record().Values[0].(string),
			ExternalID:   int(data.Record().Values[1].(float64)),
			UserURL:      data.Record().Values[2].(string),
			FollowersURL: data.Record().Values[3].(string),
			FollowingURL: data.Record().Values[4].(string),
			ReposURL:     data.Record().Values[5].(string),
			Type:         data.Record().Values[6].(string),
			SiteAdmin:    data.Record().Values[7].(bool),
		}
		followers = append(followers, follower)
	}

	return followers, nil
}
