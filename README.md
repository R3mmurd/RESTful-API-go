# RESTful-API with Go

A very simple example of how to do a RESTful-API with the Go Programming Language by 
using [gin](https://github.com/gin-gonic/gin).

To keep things simple, the data is stored in memory instead of a database.

There are three models: Person, Movie, and Star. They represent an actor or an actress, 
a movie, and a relationship between a person and a movie.

This project provides the following services: list people, get a person by id, create a
person, list movies, get a movie by id, and create a movie.