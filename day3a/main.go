package main

import (
	"fmt"
	"strconv"
)

func main() {
	lines, _ := readLines("example.txt")
	// lines, _ := readLines("input.txt")
	inputdigits := make(map[int][]int)
	for _, line := range lines {
		for i, char := range line {
			inputnumber, _ := strconv.Atoi(string(char))
			inputdigits[i] = append(inputdigits[i], inputnumber)
		}
	}
	var result string
	for _, digits := range inputdigits {
		var count1, count0 int
		for _, digit := range digits {
			switch digit {
			case 0:
				count0++
			case 1:
				count1++
			}

		}
		if count0 > count1 {
			fmt.Printf("%d", 0)
			result = result + fmt.Sprintf("%d", 0)
		} else {
			fmt.Printf("%d", 1)
			result = result + fmt.Sprintf("%d", 1)
		}

	}
	output, _ := strconv.ParseInt(result, 2, 64)
	fmt.Println()
	fmt.Println(output)
	// fmt.Println(digits[0])
	// fmt.Println(digits[1])
}
