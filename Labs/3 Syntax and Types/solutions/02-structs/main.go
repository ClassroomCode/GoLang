package main

import "fmt"

// declare a variable called "movie" of type "Movie"
// add the following fields:
// - Title (string)
// - Released (bool)
// - Length (int)
type Movie struct {
	Title    string
	Released bool
	Length   int
}

func main() {
	// declare a variable called "movie"
	var movie Movie

	// Set the Title to "Wizard of Oz"
	// Set the Released variable to "true"
	// Set Length to 125
	movie.Title = "Wizard of Oz"
	movie.Released = true
	movie.Length = 125

	// Print the value of "movie" out
	// hint: you can use fmt.Println(movie)
	fmt.Println(movie)
}
