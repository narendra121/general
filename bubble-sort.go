/*
Bubble Sort

Algorithm Used :- Checking the adjucent value
and proceed to shift the places
if the adjecent value is less than present value
*/
package main

import "fmt"

var mysl = []int{1, 5, 3, 6, 0, 7, 67, 75, 54}

func main() {
	fmt.Println(bubbleSort(mysl))
}

func bubbleSort(inp []int) []int {
	for i := 0; i < len(inp)-1; i++ {
		for j := 0; j < len(inp)-i-1; j++ {
			if inp[j] > inp[j+1] {
				inp[j], inp[j+1] = inp[j+1], inp[j]
			}
		}
	}
	return inp
}
