package delivery

import (
	"net/http"

	"ghdataapi.htm/users/service"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UsersHandler interface {
	GetUser(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollowing(c *gin.Context)
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
	userName := c.Param("username")
	user, err := h.usersService.GetUser(userName)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)

	return
}

func (h *usersHandler) GetUserFollowers(c *gin.Context) {
	userName := c.Param("username")
	maxCount := c.DefaultQuery("maxCount", "10")
	followers, err := h.usersService.GetUserFollowers(userName, maxCount)
	if err != nil || followers == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, followers)

	return
}

func (h *usersHandler) GetUserFollowing(c *gin.Context) {
	userName := c.Param("username")
	maxCount := c.DefaultQuery("maxCount", "10")
	following, err := h.usersService.GetUserFollowing(userName, maxCount)
	if err != nil || following == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, following)

	return
}

// func GetUsers(c *gin.Context) {
// 	c.DefaultQuery("", "none")
// 	c.DefaultQuery("follow", "")
// 	c.DefaultQuery("followedBy", "")
// 	c.DefaultQuery("contribute", "")
// }
