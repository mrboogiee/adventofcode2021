package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var increases int
	var inputvalues []int
	latest := 10000000

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		// fmt.Printf("%d %d\n", latest, current)
		if err != nil {
			log.Fatal(err)
		}
		inputvalues = append(inputvalues, current)
		// if latest < current {
		// 	increases++
		// }
		// latest = current
	}
	for i := 0; i < len(inputvalues)-2; i++ {
		slidingwindow := inputvalues[i] + inputvalues[i+1] + inputvalues[i+2]
		if slidingwindow > latest {
			increases++
		}
		latest = slidingwindow
	}
	fmt.Println(increases)
}
