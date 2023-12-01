package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

	re1 := regexp.MustCompile("one")
	re2 := regexp.MustCompile("two")
	re3 := regexp.MustCompile("three")
	re4 := regexp.MustCompile("four")
	re5 := regexp.MustCompile("five")
	re6 := regexp.MustCompile("six")
	re7 := regexp.MustCompile("seven")
	re8 := regexp.MustCompile("eight")
	re9 := regexp.MustCompile("nine")

	for scanner.Scan() {
		// For each line prepare a map to store positons and values
		integersPositions := map[int]int{}

		text_line := scanner.Text()

		// Ge the number values themselves
		for index, char := range text_line { // For each char
			if value, err := strconv.Atoi(string(char)); err == nil { // Check if they are redable as ints
				integersPositions[index] = value
			}
		}

		// Find using regex all the matches with thes trings we compiled before and store them
		integersPositions = findMatch(re1, &integersPositions, text_line, 1)
		integersPositions = findMatch(re2, &integersPositions, text_line, 2)
		integersPositions = findMatch(re3, &integersPositions, text_line, 3)
		integersPositions = findMatch(re4, &integersPositions, text_line, 4)
		integersPositions = findMatch(re5, &integersPositions, text_line, 5)
		integersPositions = findMatch(re6, &integersPositions, text_line, 6)
		integersPositions = findMatch(re7, &integersPositions, text_line, 7)
		integersPositions = findMatch(re8, &integersPositions, text_line, 8)
		integersPositions = findMatch(re9, &integersPositions, text_line, 9)

		// prepare a slice for the keys
		keys := make([]int, 0, len(integersPositions))

		// Iterate over the map and append each key to the slice
		for k := range integersPositions {
			keys = append(keys, k)
		}

		// Sort the slice
		sort.Ints(keys)

		// get first and last and combine
		line_value := integersPositions[keys[0]]*10 + integersPositions[keys[len(integersPositions)-1]]
		total += line_value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return result
	fmt.Printf("TOTAL: %d", total)
}

// absolutely horrid thing
func findMatch(rex *regexp.Regexp, in_integersPositions *map[int]int, text_line string, value int) map[int]int {
	// dereference so i can access values
	integersPositions := *in_integersPositions
	// get matches
	matches := rex.FindAllStringIndex(text_line, -1)
	// iterate so we can get the start indices only and save them
	// likely possible more quickly using some functional method
	for _, match := range matches {
		// my genius also means i have to pass the value itself manually instead of having ti like saved somewhere or some shit
		integersPositions[match[0]] = value
	}
	// we can return the values themselves instead of another reference even though it's probably super inefficent and dumb
	return integersPositions
}
