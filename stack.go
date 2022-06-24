package main

import "fmt"

type Mysl []int

func (ml *Mysl) Push(inp int) {
	*ml = append(*ml, inp)
}

func (ml *Mysl) Pop() {
	if ml.IsEmpty() {
		return
	}
	*ml = (*ml)[:len(*ml)-1]
}
func (ml *Mysl) IsEmpty() bool {
	return len((*ml)) <= 0
}
func main() {
	mySlice := make(Mysl, 0)
	mySlice.Pop()
	fmt.Println(mySlice)
	mySlice.Push(1)
	fmt.Println(mySlice)
	mySlice.Pop()
	fmt.Println(mySlice)

}
