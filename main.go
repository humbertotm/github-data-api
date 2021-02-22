package main

import (
	"log"

	rDelivery "ghdataapi.htm/repos/delivery"
	uDelivery "ghdataapi.htm/users/delivery"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	// if err := system.InitConfig(); err != nil {
	// 	log.Fatal.Fatal("Failed to set up config from environment")
	// }
	// log.InitLogger()

	db, err := InitDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	router := InitRouter(db)

	router.Run()
}

func InitDb() (neo4j.Driver, error) {
	return neo4j.NewDriver("neo4j://localhost:7687", neo4j.BasicAuth("", "", ""))
}

func InitRouter(db neo4j.Driver) *gin.Engine {
	r := gin.Default()

	usersHandler := uDelivery.NewUsersHandler(db)
	// r.GET("/user/:username", usersHandler.GetUser)
	users := r.Group("/users")
	{
		users.GET("/user/:username", usersHandler.GetUser)
		users.GET("/user/:username/followers", usersHandler.GetUserFollowers)
		users.GET("/user/:username/following", usersHandler.GetUserFollowing)
	}

	reposHandler := rDelivery.NewReposHandler(db)
	repos := r.Group("/repositories")
	{
		repos.GET("/repository/:username/:reponame", reposHandler.GetRepo)
		repos.GET("/repository/:username/:reponame/contributors")
	}

	return r
}
