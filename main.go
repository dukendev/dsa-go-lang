package main

import (
	"com.github/dukendev/basics/algorithms"
)

func main() {
	l := []int{12, 13, 4, 9, 7}
	algorithms.Print(l)
	algorithms.BubbleSort(l)
	algorithms.Print(l)
}
