package main

import (
	"Sample/application/config"
	"Sample/application/http/router"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDB *mongo.Database

func main() {
	fmt.Println("starting sample application.....")

	conf := config.ParseConfig("./configuration")
	//MongoDB = database.Init(conf.AppConfig.DBConnString)

	port := conf.AppConfig.Port

	router.Run(port)
}