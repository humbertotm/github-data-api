package delivery

import (
	"net/http"

	"ghdataapi.htm/repos/service"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type ReposHandler interface {
	GetRepo(c *gin.Context)
}

type reposHandler struct {
	reposService service.ReposService
}

var handler ReposHandler

func NewReposHandler(db neo4j.Driver) ReposHandler {
	if handler == nil {
		handler = &reposHandler{
			reposService: service.NewReposService(db),
		}
	}

	return handler
}

func (h *reposHandler) GetRepo(c *gin.Context) {
	repo, _ := h.reposService.GetRepo("screener")
	c.JSON(http.StatusOK, repo)

	return
}
