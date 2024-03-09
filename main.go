package main

import (
	"com.github/dukendev/basics/algorithms"
	"com.github/dukendev/basics/problems"
	"fmt"
)

func main() {
	l := []int{3, 6, 1, 7, 8, 4}
	algorithms.Print(l)
	l = algorithms.MergeSort(l)
	algorithms.Print(l)
	chars := []byte{'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c', 'c', 'd', 'e', 'e'}
	fmt.Print(problems.Compress(chars))
	fmt.Println(problems.TwoSum([]int{2, 7, 11, 15}, 18))
}
