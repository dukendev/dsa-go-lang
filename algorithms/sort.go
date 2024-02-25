package algorithms

import "fmt"

// Todo : improve it
func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
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
