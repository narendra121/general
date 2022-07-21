/*
What is Go routine?
It is a indipendent thread that runs along with the main thread.

Wait Groups are used to manage the Go routines and helpful to make sure that the all go routines completes their tasks.

Can we Use Wait groups and channels at a time?

Yes , We can use but that's not necessary. Bcz channels wont close until unless they recieves the data from go routines.
*/
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
