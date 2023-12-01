package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		// For each line prepare a slice to store the detected integers
		var integers []int

		for _, char := range scanner.Text() { // For each char
			if value, err := strconv.Atoi(string(char)); err == nil { // Check if they are redable as ints
				integers = append(integers, value) // if yes store
			}
		}
		// get forst and last and combine
		line_value := integers[0]*10 + integers[len(integers)-1]
		total += line_value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return result
	fmt.Printf("TOTAL: %d", total)
}
