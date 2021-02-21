package delivery

import (
	"net/http"

	"ghdataapi.htm/users/service"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UsersHandler interface {
	GetUser(c *gin.Context)
}

type usersHandler struct {
	usersService service.UsersService
}

var handler UsersHandler

func NewUsersHandler(db neo4j.Driver) UsersHandler {
	if handler == nil {
		handler = &usersHandler{
			usersService: service.NewUsersService(db),
		}
	}

	return handler
}

func (h *usersHandler) GetUser(c *gin.Context) {
	user, _ := h.usersService.GetUser("mojombo")
	// [wololo] Handle error
	c.JSON(http.StatusOK, user)

	return
}
