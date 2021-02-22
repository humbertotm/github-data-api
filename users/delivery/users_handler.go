package delivery

import (
	"net/http"

	"ghdataapi.htm/domain"
	"ghdataapi.htm/users/service"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// UsersHandler defines the interface for the handler of user APIs
type UsersHandler interface {
	GetUser(c *gin.Context)
	GetUserFollowers(c *gin.Context)
	GetUserFollowing(c *gin.Context)
}

type usersHandler struct {
	usersService service.UsersService
}

var handler UsersHandler

// NewUsersHandler returns and instance of UsersHandler
func NewUsersHandler(db neo4j.Driver) UsersHandler {
	if handler == nil {
		handler = &usersHandler{
			usersService: service.NewUsersService(db),
		}
	}

	return handler
}

// GetUser retrieves a user identified by its username
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

// GetUserFollowers returns a user, and a list of its followers
func (h *usersHandler) GetUserFollowers(c *gin.Context) {
	userName := c.Param("username")
	user, followers, err := h.usersService.GetUserFollowers(userName)
	if err != nil || followers == nil {
		c.Status(http.StatusNotFound)
		return
	}

	resp := struct {
		User           *domain.User   `json:"user"`
		FollowersCount int            `json:"followers_count"`
		Followers      []*domain.User `json:"followers"`
	}{
		User:           user,
		FollowersCount: len(followers),
		Followers:      followers,
	}

	c.JSON(http.StatusOK, resp)

	return
}

// GetUserFollowing returns a user and a list of followed users
func (h *usersHandler) GetUserFollowing(c *gin.Context) {
	userName := c.Param("username")
	user, following, err := h.usersService.GetUserFollowing(userName)
	if err != nil || following == nil {
		c.Status(http.StatusNotFound)
		return
	}

	resp := struct {
		User           *domain.User   `json:"user"`
		FollowingCount int            `json:"following_count"`
		Following      []*domain.User `json:"following"`
	}{
		User:           user,
		FollowingCount: len(following),
		Following:      following,
	}

	c.JSON(http.StatusOK, resp)

	return
}
