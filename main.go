package main

// Step 1: Import the necessary packages
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Step 2: Create the movie and director struct
type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // * means pointer to Director
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Step 8: Create the movie slice
var movies []Movie // slice of Movie

// Step 9: Create a function to initialize the movies (2 parts)
// 9.1 Initialize the movies
func initializeMovies() {
	movies = append(movies, Movie{Id: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "Derrick", LastName: "Strong"}})
	movies = append(movies, Movie{Id: "2", Isbn: "847564", Title: "Movie Two", Director: &Director{FirstName: "Shavon Nicole", LastName: "Strong"}})
}

// Step 10: Create the 5 CRUD routes (5 parts)
// 10.1 Get all movies (GET)
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the header
	json.NewEncoder(w).Encode(movies)                  // encode the movies and write to the response
}

// 10.2 Get a single movie (GET)
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the header
	params := mux.Vars(r)                              // get the params
	// loop through the movies
	for _, item := range movies {
		if item.Id == params["id"] { // if the movie id matches the id in the params
			json.NewEncoder(w).Encode(item) // encode the movie and write to the response (movie found)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{}) // encode an empty movie and write to the response (no movie found)
}

// 10.3 Create a movie (POST)
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the header
	var movie Movie                                    // create a movie
	_ = json.NewDecoder(r.Body).Decode(&movie)         // decode the request body and assign to movie (movie created)
	movie.Id = strconv.Itoa(rand.Intn(1000000))        // generate an id for the movie and assign to movie
	movies = append(movies, movie)                     // append the movie to the movies slice
	json.NewEncoder(w).Encode(movie)                   // return and encode the movie and write to the response (movie created)
}

// 10.4 Update a movie (PUT)
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the header
	params := mux.Vars(r)                              // get the params
	for index, item := range movies {                  // loop through the movies
		if item.Id == params["id"] { // if the movie id matches the id in the params
			movies = append(movies[:index], movies[index+1:]...) // remove the movie from the movies
			var movie Movie                                      // create a movie
			_ = json.NewDecoder(r.Body).Decode(&movie)           // decode the request body and assign to movie (movie updated)
			movie.Id = params["id"]                              // set the id of the movie to the id in the params
			movies = append(movies, movie)                       // append the movie to the movies slice
			json.NewEncoder(w).Encode(movie)                     // encode the movie and write to the response (movie updated)
			return
		}
	}
	json.NewEncoder(w).Encode(movies) // encode the movies and write to the response (movie not found)
}

// 10.5 Delete a movie (DELETE)
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the header
	params := mux.Vars(r)                              // get the params
	for index, item := range movies {                  // loop through the movies
		if item.Id == params["id"] { // if the movie id matches the id in the params
			movies = append(movies[:index], movies[index+1:]...) // remove the movie from the movies
			break                                                // break the loop
		}
	}
	json.NewEncoder(w).Encode(movies) // encode the movies and write to the response (movies minus the deleted movie)
}

// Step 3: Create the main function
// Main function
func main() {
	// Step 4: Create a new router
	r := mux.NewRouter() // create a new router

	// 9.2: Initialize the movies
	initializeMovies()

	// Step 5: Create the routes
	r.HandleFunc("/movies", getMovies).Methods("GET")           // get all movies
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")       // get a single movie
	r.HandleFunc("/movies", createMovie).Methods("POST")        // create a movie
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")    // update a movie
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE") // delete a movie

	// Step 6: Add the router to the server
	http.Handle("/", r) // handle all requests with the router

	// Step 7: Start the server
	fmt.Println("Server running on port 8000")   // print a message
	log.Fatal(http.ListenAndServe(":8000", nil)) // start the server
}
