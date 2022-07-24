package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	var wg sync.WaitGroup
	wg.Add(len(n))

	for i, slice := range n {
		go run(i, slice, &wg)
	}
	wg.Wait()
}

func run(i int, slice []int, wg *sync.WaitGroup) {
	fmt.Printf("slice %d: ", i)
	sum(slice)
	wg.Done()
}

func sum(s []int) {
	var E int
	for _, k := range s {
		E += k
	}
	str := fmt.Sprintln(E)
	if _, err := fmt.Fprintf(os.Stdout, str); err != nil {
		fmt.Printf("error printing sum of slice: %v\n", err)
	}
}
