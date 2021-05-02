package main

import (
	"BasicAuth/Controllers"
	"BasicAuth/Middlewares"
	"BasicAuth/models"
	"fmt"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase() //Required for other db function
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://a910c0cb213c41aebcfb967de81f5b6d@o596342.ingest.sentry.io/5742409",
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(sentrygin.New(sentrygin.Options{}))
	Authorized := r.Group("/", Middlewares.BasicAuth())

	Authorized.GET("/secret", Controllers.UserGet)
	r.Run()

}
