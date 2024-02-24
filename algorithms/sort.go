package algorithms

import "fmt"

func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func Print(list []int) {
	fmt.Println()
	for _, i := range list {
		fmt.Print(i, " ")
	}
}
