package main

import (
	"BasicAuth/Controllers"
	"BasicAuth/Middlewares"
	"BasicAuth/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase() //Required for other db func
	r := gin.Default()

	Authorized := r.Group("/", Middlewares.BasicAuth())

	Authorized.GET("/secret", Controllers.UserGet)
	r.Run()

}
