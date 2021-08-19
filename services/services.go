// Package services provides handlers to retrieve and create data.
package services

import (
	"net/http"
	"strconv"

	"github.com/R3mmurd/RESTful-API-go/models"
	"github.com/gin-gonic/gin"
)

var people = []models.Person{
	{ID: 1, FirstName: "Daniel", LastName: "Radcliffe", BirthYear: 1989},
	{ID: 2, FirstName: "Emma", LastName: "Watson", BirthYear: 1990},
	{ID: 3, FirstName: "Rupert", LastName: "Grint", BirthYear: 1988},
	{ID: 4, FirstName: "Russell", LastName: "Crowe", BirthYear: 1964},
}

var movies = []models.Movie{
	{ID: 1, Title: "Harry Potter and the Philosopher's Stone", PublishedYear: 2001},
	{ID: 2, Title: "The F Word", PublishedYear: 2013},
	{ID: 3, Title: "Noah", PublishedYear: 2014},
}

var stars = []models.Star{
	{ID: 1, PersonId: 1, MovieId: 1},
	{ID: 2, PersonId: 3, MovieId: 1},
	{ID: 3, PersonId: 3, MovieId: 1},
	{ID: 4, PersonId: 1, MovieId: 2},
	{ID: 5, PersonId: 2, MovieId: 3},
	{ID: 6, PersonId: 4, MovieId: 3},
}

// GetPeople responds with the list of all people.
func GetPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

// GetPersonById retrieves the person whose ID matches with the id sent as a parameter
func GetPersonById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, v := range people {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

// PostPeople adds a new person from JSON received in the request body.
func PostPeople(c *gin.Context) {
	var NewPerson models.Person

	err := c.BindJSON(&NewPerson)

	if err != nil {
		return
	}

	people = append(people, NewPerson)
	c.IndentedJSON(http.StatusCreated, NewPerson)
}

// GetMovies responds with the list of all movies.
func GetMovies(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, movies)
}

// GetMovieById retrieves the movie whose ID matches with the id sent as a parameter
func GetMovieById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, v := range movies {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
}

// PostMovies adds a new person from JSON received in the request body.
func PostMovies(c *gin.Context) {
	var NewMovie models.Movie

	err := c.BindJSON(&NewMovie)

	if err != nil {
		return
	}

	movies = append(movies, NewMovie)
	c.IndentedJSON(http.StatusCreated, NewMovie)
}
