package main

import (
	"fmt"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	inputdigits := make(map[int][]int)
	for _, line := range lines {
		for i, char := range line {
			inputnumber, _ := strconv.Atoi(string(char))
			inputdigits[i] = append(inputdigits[i], inputnumber)
		}
	}
	gammaBinary := getGammaBinary(inputdigits)
	gammaDecimal, _ := strconv.ParseInt(gammaBinary, 2, 64)
	fmt.Printf("Gamma Binary: %s Decimal: %d\n", gammaBinary, gammaDecimal)
	epsilonBinary := getEpsilonBinary(inputdigits)
	epsilonDecimal, _ := strconv.ParseInt(epsilonBinary, 2, 64)
	fmt.Printf("Epsilon Binary: %s Decimal: %d\n", epsilonBinary, epsilonDecimal)
	fmt.Printf("Power: %d", epsilonDecimal*gammaDecimal)
}

func getGammaBinary(inputdigits map[int][]int) string {
	var gammaBinary string
	for i := 0; i < len(inputdigits); i++ {
		var count1, count0 int
		for _, digit := range inputdigits[i] {
			switch digit {
			case 0:
				count0++
			case 1:
				count1++
			}

		}
		if count0 > count1 {
			fmt.Printf("%d", 0)
			gammaBinary = gammaBinary + fmt.Sprintf("%d", 0)
		} else {
			fmt.Printf("%d", 1)
			gammaBinary = gammaBinary + fmt.Sprintf("%d", 1)
		}
	}
	fmt.Printf("\n")
	return gammaBinary
}

func getEpsilonBinary(inputdigits map[int][]int) string {
	var epsilonBinary string

	for i := 0; i < len(inputdigits); i++ {
		var count1, count0 int
		for _, digit := range inputdigits[i] {
			switch digit {
			case 0:
				count0++
			case 1:
				count1++
			}

		}
		if count0 < count1 {
			fmt.Printf("%d", 0)
			epsilonBinary = epsilonBinary + fmt.Sprintf("%d", 0)
		} else {
			fmt.Printf("%d", 1)
			epsilonBinary = epsilonBinary + fmt.Sprintf("%d", 1)
		}
	}
	fmt.Printf("\n")
	return epsilonBinary
}
