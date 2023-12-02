package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var total int = 0

	for scanner.Scan() {

		line_value := parseLine(scanner.Text())
		total += line_value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return result
	fmt.Printf("TOTAL: %d", total)
}

func parseLine(line string) int {
	limits := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	split := strings.Split(line, ":")                          // split[0] contains ID, [1] the results
	for _, revelation := range strings.Split(split[1], "; ") { // Split results for each extraction
		for _, dice := range strings.Split(revelation, ", ") { // split extraction for each color
			split_dice := strings.Split(strings.TrimSpace(dice), " ") // get number and color separatedly
			val, _ := strconv.Atoi(split_dice[0])                     // get the int
			if val > limits[split_dice[1]] {                          // compare to max values
				limits[split_dice[1]] = val
			}
		}
	}
	power := limits["red"] * limits["green"] * limits["blue"]

	return power
}
