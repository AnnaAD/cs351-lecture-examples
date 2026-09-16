/*
	Contains: Code for unparallelized "book rec. system"
*/

package main

import (
	"log"
	"math/rand"
	"time"
	"github.com/montanaflynn/stats"
	"strconv"
)

// Book is a single title in the catalog.
type Book struct {
	Title  string
	MinAge float64 // youngest recommended reader age
	Stars  float64 // average star rating (1-5)
	Length float64 // page count
}

// Ratings holds the raw readings pulled from the DB for a set of books.
type Ratings struct {
	Ages   []float64
	Stars  []float64
	Length []float64
}

// Profile summarizes a user's reading habits.
type Profile struct {
	mean   float64 // mean reader age
	median float64 // median star rating
	p99th  float64 // 99th percentile book length
}

// Database is a fake data store that returns dummy data.
type Database struct {
	all_books []Book
	history   map[string][]string // user -> titles they've read
}

func NewDatabase() *Database {

	books := []Book{
			{Title: "The Hobbit", MinAge: 10, Stars: 4.7, Length: 310},
			{Title: "Dune", MinAge: 14, Stars: 4.5, Length: 412},
			{Title: "Charlotte's Web", MinAge: 6, Stars: 4.4, Length: 184},
			{Title: "The Martian", MinAge: 13, Stars: 4.6, Length: 369},
			{Title: "Project Hail Mary", MinAge: 13, Stars: 4.7, Length: 476},
			{Title: "Pride and Prejudice", MinAge: 12, Stars: 4.3, Length: 432},
			{Title: "Gone Girl", MinAge: 18, Stars: 4.1, Length: 422},
			{Title: "The Road", MinAge: 16, Stars: 3.9, Length: 287},
			{Title: "War and Peace", MinAge: 16, Stars: 4.2, Length: 1225},
			{Title: "Twilight", MinAge: 13, Stars: 3.6, Length: 498},
	}

	for i := range 1000 {
		books = append(books, Book{
			Title:  "Book " + strconv.Itoa(i),
			MinAge: float64(rand.Intn(18)),       // 0-17
			Stars:  1 + rand.Float64()*4,         // 1.0-5.0
			Length: float64(50 + rand.Intn(1150)), // 50-1199 pages
		})
	}

	return &Database{
		all_books: books,
		history: map[string][]string{
			"anna": {"The Hobbit", "Dune", "The Martian"},
			"bob":  {"The Road", "Gone Girl"},
		},
	}
}


var DB = NewDatabase()

// get_books returns the books a user has read.
// Sleeps to simulate a slow DB query.
func (db *Database) get_books(user string) []Book {
	time.Sleep(50 * time.Millisecond)

	books := make([]Book, 0)
	for _, title := range db.history[user] {
		for _, b := range db.all_books {
			if b.Title == title {
				books = append(books, b)
			}
		}
	}
	return books
}

// get_all_ratings returns fake ratings for the given books.
// Sleeps to simulate a slow DB query.
func (db *Database) get_all_ratings(books []Book) Ratings {
	time.Sleep(100 * time.Millisecond)
	numReadings := 100000

	r := Ratings{
		Ages:   make([]float64, 0, numReadings),
		Stars:  make([]float64, 0, numReadings),
		Length: make([]float64, 0, numReadings),
	}
	if len(books) == 0 {
		return r
	}

	for i := 0; i < numReadings; i++ {
		b := books[rand.Intn(len(books))]

		age := b.MinAge + rand.Float64()*30         // reader somewhere above min age
		stars := b.Stars + (rand.Float64() - 0.5)   // +/- 0.5 stars of noise
		stars = min(max(stars, 1), 5)               // clamp to 1-5

		r.Ages = append(r.Ages, age)
		r.Stars = append(r.Stars, stars)
		r.Length = append(r.Length, b.Length)
	}
	return r
}

func is_good_match(book Book, profile Profile) bool {
	time.Sleep(time.Millisecond*10)
	return book.MinAge <= profile.mean &&
		book.Stars >= profile.median-0.5 &&
		book.Length <= profile.p99th*1.25
}

// List of users who made request
var request_queue = []string{"anna", "bob"}

func main() {

	start := time.Now()

	for _, user := range request_queue {
		log.Println("Getting Books: ", user)
		books := DB.get_books(user)
		ratings := DB.get_all_ratings(books)

		log.Println("Getting Stats Recommended:", user)

		var profile Profile
		profile.mean, _ = stats.Mean(ratings.Ages)
		profile.median, _ = stats.Median(ratings.Stars)
		profile.p99th, _ = stats.Percentile(ratings.Length, 99)
		
		log.Println("Getting Matches Recommended:", user)

		output := make([]string, 0)
		for i := 0; i < len(DB.all_books); i += 1 {
			book := DB.all_books[i]
			if is_good_match(book, profile) {
				output = append(output, book.Title)
			}
		}

		log.Println("Books Recommended:", user)
	}

	timeElapsed := time.Now().Sub(start)
	log.Println("Time Elapsed: ", timeElapsed.Seconds())
}