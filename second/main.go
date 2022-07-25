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
	var result int
	for _, s := range n {
		r := make(chan int)
		go sum(s, r)
		result += <-r
	}

	fmt.Printf("result: %v\n", result)

}

func sum(s []int, gochan chan int) {
	var E int
	for _, v := range s {
		E += v
	}
	gochan <- E
}
