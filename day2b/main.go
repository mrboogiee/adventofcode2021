package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	// lines, _ := readLines("example.txt")
	lines, _ := readLines("input.txt")
	horizontal := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		var re = regexp.MustCompile("(\\w*) (\\d*)")
		result := re.FindStringSubmatch(line)
		action := result[1]
		amount, _ := strconv.Atoi(result[2])
		switch action {
		case "forward":
			horizontal += amount
			aimchange := aim * amount
			depth += aimchange
			fmt.Printf("horizontal: %d depth %d aim: %d amount: %d aimchange: %d\n", horizontal, depth, aim, amount, aimchange)
		case "down":
			aim += amount
			fmt.Printf("horizontal: %d depth %d aim: %d amount: %d \n", horizontal, depth, aim, amount)
		case "up":
			aim -= amount
			fmt.Printf("horizontal: %d depth %d aim: %d amount: %d \n", horizontal, depth, aim, amount)
		default:
			log.Fatal(action)
		}
	}
	fmt.Println(horizontal * depth)
}
