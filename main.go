package main

import (
	"Algorithms/sort"
	"fmt"
)

func main() {
	arr := []int{64, 73, 82, 4, 2, 1, 6}
	sort.BubbleSort(arr)
	fmt.Print(arr)
}
