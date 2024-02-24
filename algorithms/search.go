package algorithms

import "fmt"

func LinearSearch(list []int, key int) int {
	for i, item := range list {
		if item == key {
			fmt.Println("found", item, "at", i)
			return i
		}
	}
	return -1
}

func BinarySearch(list []int, key int) bool {
	start := 0
	end := len(list)

	for start < end {
		mid := start + (end-start)/2
		switch {
		case list[mid] > key:
			end = mid - 1
		case list[mid] < key:
			start = mid + 1
		default:
			return true
		}
	}

	return false
}

func BinarySearchRec(list []int, key int) bool {
	return binarySearch(list, key, 0, len(list)-1)
}

func binarySearch(list []int, key int, s int, e int) bool {
	if s > e {
		return false
	}
	m := s + (e-s)/2
	switch {
	case list[m] > key:
		return binarySearch(list, key, s, m-1)
	case list[m] < key:
		return binarySearch(list, key, m+1, e)
	default:
		return true
	}
}
