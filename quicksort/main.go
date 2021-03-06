package main

import (
	"fmt"

	"github.com/joaonsantos/goscripts-learning/quicksort/sort"
)

func main() {
	s := []int{6, 3, 2, 5, 4, 1, 10, 9, 8, 7}
	fmt.Println("before quicksort:", s)

	ss := sort.Quicksort(s)
	fmt.Println("after quicksort:", ss)
}
