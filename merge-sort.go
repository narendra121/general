package main

import "fmt"

var mysl = []int{1, 5, 3, 6, 0, 7, 67, 75, 54}

func main() {
	fmt.Println(performMergeSort(mysl))
}
func performMergeSort(sl []int) []int {
	if len(sl) <= 2 {
		return sl
	}
	left := performMergeSort(sl[:len(sl)/2])
	right := performMergeSort(sl[len(sl)/2:])
	return mergeSort(left, right)
}

func mergeSort(left, right []int) []int {
	resp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			resp = append(resp, left[i])
			i++
		} else {
			resp = append(resp, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		resp = append(resp, left[i])
	}
	for ; j < len(left); j++ {
		resp = append(resp, right[j])
	}
	return resp
}
