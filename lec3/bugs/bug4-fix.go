package main

import (
	"log"
)

func add(slice *[]int, val int) {
	*slice = append(*slice, val)
	log.Println(slice)
}

func main() {
	slice := []int{1,2,3,4}
	log.Println(slice)
	add(&slice, 5)
	log.Println(slice)
}
