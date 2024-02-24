package main

import (
	"fmt"

	"com.github/dukendev/basics/algorithms"
)

func main() {
	l := []int{2, 3, 4, 5, 7}
	fmt.Print(algorithms.BinarySearchRec(l, 6))
}
