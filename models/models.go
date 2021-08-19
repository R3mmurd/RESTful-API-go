// Package models provides some types.
package models

// Person is a type to represent an actor or an actress.
type Person struct {
	ID        int
	FirstName string
	LastName  string
	BirthYear int
}

// Movie type
type Movie struct {
	ID            int
	Title         string
	PublishedYear int
}

// Star represents a relationship between a person and a movie.
type Star struct {
	ID       int
	PersonId int
	MovieId  int
}
