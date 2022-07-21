/*

Queue

Algorithm:- FIFO -First in First Out

Ops:- Queuing , Dequeing
*/
package main

import "fmt"

type mySlice []int

func (ml *mySlice) Queue(inp int) {
	*ml = append(*ml, inp)
}

func (ml *mySlice) DeQueue() {
	if ml.IsEmpty() {
		return
	}
	*ml = (*ml)[1:]
}
func (ml *mySlice) IsEmpty() bool {
	return len((*ml)) <= 0
}
func main() {
	mySlice := make(mySlice, 0)
	mySlice.DeQueue()
	fmt.Println(mySlice)
	mySlice.Queue(1)
	fmt.Println(mySlice)
	mySlice.DeQueue()
	fmt.Println(mySlice)
}
