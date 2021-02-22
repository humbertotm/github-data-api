package routes

import (
	rDelivery "ghdataapi.htm/repos/delivery"
	uDelivery "ghdataapi.htm/users/delivery"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func InitRouter(db neo4j.Driver) *gin.Engine {
	r := gin.Default()

	usersHandler := uDelivery.NewUsersHandler(db)
	users := r.Group("/users")
	{
		users.GET("/:username", usersHandler.GetUser)
		users.GET("/:username/followers", usersHandler.GetUserFollowers)
		users.GET("/:username/following", usersHandler.GetUserFollowing)
	}

	reposHandler := rDelivery.NewReposHandler(db)
	repos := r.Group("/repositories")
	{
		repos.GET("/:reponame/user/:username", reposHandler.GetRepo)
		repos.GET("/:reponame/user/:username/contributors", reposHandler.GetRepoContributors)
	}

	return r
}
