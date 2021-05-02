package main

import (
	"BasicAuth/Controllers"
	"BasicAuth/Middlewares"
	"BasicAuth/Models"

	"github.com/gin-gonic/gin"
)

func main() {

	Models.ConnectDatabase() //Initialize the DB connection
	r := gin.Default()
	Authorized := r.Group("/v1", Middlewares.BasicAuth()) //Set the routes which require authentication
	r.GET("/generate/:email", Middlewares.Gen)            //To generate API keys
	Authorized.GET("/secret", Controllers.UserGet)        //Test authentication
	r.Run()
}
