package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printMe(100, &wg)
	wg.Wait()
}

func printMe(inp int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("================= ", sum, " ====================")
}
