package main

import (
	"fmt"

	"com.github/dukendev/basics/algorithms"
	"com.github/dukendev/basics/problems"
)

func main() {
	l := []int{3, 6, 1, 7, 8, 4}
	algorithms.Print(l)
	l = algorithms.MergeSort(l)
	algorithms.Print(l)
	chars := []byte{'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'd', 'e', 'e'}
	chars2 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Print(problems.Compress(chars))
	fmt.Print(problems.Compress(chars2))
}
