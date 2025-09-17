/*
	Contains: Code for "satellite", with pipeline parallelism
*/

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

type ReadingBatch struct {
	data *[]float64
	typeReading string
}

type StatsData struct {
	data *[]float64
	typeReading string
	mean float64
	median float64
	p99th float64
}

func main() {

	start := time.Now()

	readChannel := make(chan ReadingBatch)
	sendChannel := make(chan StatsData)

	go func() {
		for _,q := range(reading_queue) {
				log.Println("Performing Readings: ", q)
				data := make([]float64,0)
				numReadings := 1000

				for i := 0; i < numReadings; i+=1 {
					data = append(data, perform_reading(q))
				}
				log.Println("Completed Readings ", q)
				readChannel <- ReadingBatch{&data, q}
		}
		close(readChannel)
	}()

	go func() {
		for readData := range(readChannel) {
			batch_normalize_reading(readData.data)
			log.Println("Normalized Readings ", readData.typeReading)
			mean,median,p99th := do_stats(readData.data)
			sendChannel <- StatsData{readData.data,readData.typeReading, mean, median, p99th}
		}
		close(sendChannel)
	}()
	
	for sendData := range(sendChannel) {
		log.Println("Send to Earth: ", sendData.typeReading)
		send_to_earth(sendData.data, sendData.mean, sendData.median, sendData.p99th)
	}

	timeElapsed := time.Now().Sub(start)
	log.Println("Time Elapsed: ", timeElapsed.Seconds())
}
