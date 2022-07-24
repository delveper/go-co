package main

import (
	"fmt"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	slices := make(chan []int)
	var result int

	go split(slices, n)
	for s := range slices {
		result += sum(s)
	}
	fmt.Printf("result: %v\n", result)
}
func split(slices chan []int, n [][]int) {
	for _, s := range n {
		slices <- s
	}
	close(slices)
}

func sum(s []int) int {
	var E int
	for _, v := range s {
		E += v
	}
	return E
}
