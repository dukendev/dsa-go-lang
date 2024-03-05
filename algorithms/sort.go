package algorithms

import (
	"fmt"
)

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

func SelectionSort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i; j < n; j++ {
			if list[j] < list[k] {
				k = j
				list[i], list[k] = list[k], list[i]
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

func QuickSort(list []int) {
	if len(list) <= 1 {
		return
	}
	pivot := PartitionAlgo(list, 0, len(list)-1)
	QuickSort(list[0:pivot])
	QuickSort(list[pivot+1:])
}

func PartitionAlgo(list []int, low int, high int) int {
	pivot := list[high]
	i := low - 1
	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}

	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func Print(list []int) {
	fmt.Println()
	for _, i := range list {
		fmt.Print(i, " ")
	}
}

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) / 2

	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return MergeLists(left, right)
}

func MergeLists(a []int, b []int) []int {
	i := 0
	j := 0
	n := len(a)
	m := len(b)
	aux := make([]int, n+m)
	k := 0
	for i < n && j < m {
		if a[i] <= b[j] {
			aux[k] = a[i]
			i++
			k++
		} else {
			aux[k] = b[j]
			j++
			k++
		}
	}

	for i < n {
		aux[k] = a[i]
		i++
		k++
	}

	for j < m {
		aux[k] = b[j]
		j++
		k++
	}

	return aux

}
