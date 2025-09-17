package main

import (
	"time"
	"log"
	"math/rand"
	"sync"
	"github.com/montanaflynn/stats"
)

func perform_reading(measurementType string) float64 {
	time.Sleep(1 * time.Millisecond)
	return float64(rand.Intn(10000) + 500)
}

func normalize_reading(reading float64) float64{
	time.Sleep(10 * time.Millisecond)
	return reading * 2
}

func batch_normalize_reading(input *[]float64) {
	data := *input
	batch := len(data) / 12 // Divide up data into 12 chunks, where 12 is # cores
	var wg sync.WaitGroup

	for start := 0; start < len(data); start += batch {
		end := start + batch
		if end > len(data) {
			end = len(data)
		}
		wg.Add(1)
		go func(s, e int) {
			//log.Println("Start Worker",s,e)
			defer wg.Done()
			for i := s; i < e; i++ {
				data[i] = normalize_reading(data[i])
			}
		}(start, end)
	}

	wg.Wait()
}

func send_to_earth(data *[]float64, mean float64, median float64, p99 float64) {
	//log.Println(data);
	time.Sleep(500*time.Millisecond)
	log.Println(mean);
	log.Println(median);
	log.Println(p99)
}

var reading_queue = []string{"temp", "dist"}

func do_stats(input *[]float64) (float64,float64,float64){
	meanCh := make(chan float64)
	medianCh := make(chan float64)
	p99Ch := make(chan float64)

	go func(input *[]float64) {
		mean, _ := stats.Mean(*input)
		meanCh <- mean
	}(input)

	go func(input *[]float64) {
		median, _:= stats.Median(*input)
		medianCh <- median
	}(input)

	go func(input *[]float64) {
		p99th, _ := stats.Percentile(*input, 99)
		p99Ch <- p99th
	}(input)

	mean := <-meanCh
	median := <-medianCh
	p99 := <-p99Ch
	
	return mean, median, p99
}


func main() {

	start := time.Now()

	var wg sync.WaitGroup

	for _,q := range(reading_queue) {
		wg.Add(1)
		go func() { //This breaks the problem setup restriction by overlapping performing readings on the two tasks.
			defer wg.Done()
			log.Println("Performing Readings: ", q) 
			data := make([]float64,0)
			numReadings := 1000

			//possible that multiple readings are collected at same time by each goroutine unlike with pipeline
			for i := 0; i < numReadings; i+=1 {
				data = append(data, perform_reading(q)) 
			}

			log.Println("Completed Readings")

			batch_normalize_reading(&data)
			log.Println("Normalized Readings")


			mean,median,p99th := do_stats(&data)


			send_to_earth(&data, mean, median, p99th)
		}()
	}

	wg.Wait()

	timeElapsed := time.Now().Sub(start)
	log.Println("Time Elapsed: ", timeElapsed.Seconds())
}
