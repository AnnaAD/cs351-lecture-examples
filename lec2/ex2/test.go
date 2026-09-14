
package main

import (
	"log"
)

func test() {
	defer log.Println("1")
	defer log.Println("2")
	defer log.Println("3")
}

func main() {
	test()
}


