/*
	Contains: Code for unparallelized "satellite"
*/

package main

import (
	"time"
	"log"
	"math/rand"
	"github.com/montanaflynn/stats"
)

// Simulates performing reading on hardware
func perform_reading(measurementType string) float64 {
	time.Sleep(1 * time.Millisecond)
	return float64(rand.Intn(10000) + 500)
}

//Simulates performing computation on reading
func normalize_reading(reading float64) float64{
	time.Sleep(10 * time.Millisecond)
	return reading * 2
}

func send_to_earth(data *[]float64, mean float64, median float64, p99 float64) {
	time.Sleep(500*time.Millisecond)
	log.Println(data);
	log.Println(mean);
	log.Println(median);
	log.Println(p99)
}


// List of reading "types" to perform
var reading_queue = []string{"temp", "dist"}

func main() {

	start := time.Now()

	for _,q := range(reading_queue) {
		log.Println("Performing Readings: ", q)
		data := make([]float64,0)
		numReadings := 1000

		for i := 0; i < numReadings; i+=1 {
			data = append(data, perform_reading(q))
		}

		for i:= 0; i < len(data); i+=1 {
			data[i] = normalize_reading(data[i])
		}


		mean, _ := stats.Mean(data) 
		median, _ := stats.Median(data) 
		p99th, _ := stats.Percentile(data, 99) 


		send_to_earth(&data, mean, median, p99th)
	}

	timeElapsed := time.Now().Sub(start)
	log.Println("Time Elapsed: ", timeElapsed.Seconds())
}
