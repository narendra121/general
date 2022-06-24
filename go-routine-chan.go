package main

import (
	"fmt"
)

func main() {

	//////// self calling func
	resChan := make(chan string)
	go func(inp string, resChan chan string) {
		res := "Hello " + inp + " !"
		resChan <- res
	}("Narendra", resChan)
	fmt.Println(<-resChan)

	///////////////////////////

}
