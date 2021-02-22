package delivery

import (
	"net/http"

	"ghdataapi.htm/domain"
	"ghdataapi.htm/log"
	"ghdataapi.htm/repos/service"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// ReposHandler defines the interface for the handler of repository APIs
type ReposHandler interface {
	GetRepo(c *gin.Context)
	GetRepoContributors(c *gin.Context)
}

type reposHandler struct {
	reposService service.ReposService
}

var handler ReposHandler

// NewReposHandler returns an instance of ReposHandler
func NewReposHandler(db neo4j.Driver) ReposHandler {
	if handler == nil {
		handler = &reposHandler{
			reposService: service.NewReposService(db),
		}
	}

	return handler
}

// GetRepo retrieves a single repo identified by its name and owner
func (h *reposHandler) GetRepo(c *gin.Context) {
	userName := c.Param("username")
	repoName := c.Param("reponame")
	repo, err := h.reposService.GetRepo(repoName, userName)
	if err != nil {
		log.Error.Println(err.Error())
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, repo)
	return
}

// GetRepoContributors retrieves a repo and a list of its contributors
func (h *reposHandler) GetRepoContributors(c *gin.Context) {
	userName := c.Param("username")
	repoName := c.Param("reponame")
	repo, contributors, err := h.reposService.GetRepoContributors(repoName, userName)
	if err != nil || repo == nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	resp := struct {
		Repo             *domain.Repo   `json:"repo"`
		ContributorCount int            `json:"contributor_count"`
		Contributors     []*domain.User `json:"contributors"`
	}{
		Repo:             repo,
		ContributorCount: len(contributors),
		Contributors:     contributors,
	}
	c.JSON(http.StatusOK, resp)
	return
}
