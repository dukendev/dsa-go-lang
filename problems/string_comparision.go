package problems

import "fmt"

func Compress(chars []byte) int {
	fmt.Println()
	count := 0
	compressed := ""
	var previous byte
	for i := range chars {
		if count > 0 {
			if previous != chars[i] {

				count++
				compressed += fmt.Sprint(count)
				count = 0
			}
		} else {
			compressed += string(chars[i])
			count++
		}

		previous = chars[i]
	}
	fmt.Print(compressed)
	fmt.Println()
	return len(compressed)
}
