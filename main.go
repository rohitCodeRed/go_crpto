package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	server := &http.Server{
		Addr:         ":4000",
		Handler:      router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func router() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	// routes := e.Group("/api")
	// {
	// 	routes.POST("/login", loginEndpoint)
	// 	routes.POST("/submit", submitEndpoint)
	// 	routes.POST("/read", readEndpoint)
	// }
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}
