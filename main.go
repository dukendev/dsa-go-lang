package main

import (
	"com.github/dukendev/basics/algorithms"
)

func main() {
	l := []int{3, 6, 1, 7, 8, 4}
	algorithms.Print(l)
	algorithms.BubbleSort(l)
	algorithms.Print(l)
}
