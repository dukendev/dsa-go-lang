package problems

import (
	"fmt"
	"strconv"
)

func Compress(chars []byte) int {
	write := 0
	count := 1
	n := len(chars)
	for read := 1; read < n; read++ {
		if chars[read] == chars[read-1] {
			count++
		} else {
			chars[write] = chars[read-1]
			write++
			if count > 1 {
				countStr := []byte(strconv.Itoa(count))
				for _, c := range countStr {
					chars[write] = byte(c)
					write++
				}
			}
			count = 1
		}
	}

	chars[write] = chars[n-1]
	write++
	if count > 1 {
		countStr := []byte(strconv.Itoa(count))
		for _, c := range countStr {
			chars[write] = byte(c)
			write++
		}
	}
	printByteArray(chars, write)
	return write
}

func printByteArray(chars []byte, length int) {
	var result string
	for i := 0; i < length; i++ {
		result += string(chars[i])
	}
	fmt.Println(result)
}
