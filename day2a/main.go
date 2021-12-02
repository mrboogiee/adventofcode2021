package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	// _ = lines
	horizontal := 0
	depth := 0
	for _, line := range lines {
		fmt.Println(line)
		var re = regexp.MustCompile("(\\w*) (\\d)")
		result := re.FindStringSubmatch(line)
		action := result[1]
		amount, _ := strconv.Atoi(result[2])
		switch action {
		case "forward":
			horizontal += amount
		case "backward":
			horizontal -= amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}
	fmt.Println(horizontal * depth)
}
