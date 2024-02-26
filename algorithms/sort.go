package algorithms

import "fmt"

func BubbleSort(list []int) {
	n := len(list)
	swapFlag := false
	for i := 0; i < n-1; i++ {
		swapFlag = false
		for j := 0; j < n-1-i; j++ {
			if list[j+1] < list[j] {
				list[j], list[j+1] = list[j+1], list[j]
				swapFlag = true
			}
		}
		if !swapFlag {
			break
		}
	}
}

func InsertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		current := list[i]
		j := i - 1
		for j >= 0 && list[j] > current {
			list[j+1] = list[j]
			j--
		}
		list[j+1] = current
	}
}

func Print(list []int) {
	fmt.Println()
	for _, i := range list {
		fmt.Print(i, " ")
	}
}
