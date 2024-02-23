package algorithms

func BinarySearch(list []int, key int) int {
	for i, item := range list {
		if item == key {
			return i
		}
	}
	return -1
}
