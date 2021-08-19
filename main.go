package main

import (
	"github.com/R3mmurd/RESTful-API-go/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/people", services.GetPeople)
	router.GET("/people/:id", services.GetPersonById)
	router.POST("/people", services.PostPeople)

	router.GET("/movies", services.GetMovies)
	router.GET("/movies/:id", services.GetMovieById)
	router.POST("/movies", services.PostMovies)

	router.Run("localhost:8080")
}
