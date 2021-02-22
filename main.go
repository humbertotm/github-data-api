package main

import (
	"log"

	"ghdataapi.htm/db"
	logger "ghdataapi.htm/log"
	"ghdataapi.htm/routes"
	"ghdataapi.htm/system"
)

func main() {
	if err := system.InitConfig(); err != nil {
		log.Fatal("Failed to set up config from environment")
	}
	logger.InitLogger()

	db, err := db.InitDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	router := routes.InitRouter(db)

	router.Run()
}
