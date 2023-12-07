package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	split_line := strings.FieldsFunc(line, split)
	fmt.Println(split_line)
	card_id, _ := strconv.Atoi(strings.Split(split_line[0], " ")[1])
	fmt.Println(card_id)
	str_win_numbers, str_got_numbers := strings.Split(strings.TrimSpace(split_line[1]), " "), strings.Split(strings.TrimSpace(split_line[2]), " ")
	win_numbers := []int{}
	for _, elem := range str_win_numbers {
		value, _ := strconv.Atoi(elem)
		if value != 0 {
			win_numbers = append(win_numbers, value)
		}
	}

	got_numbers := []int{}
	for _, elem := range str_got_numbers {
		value, _ := strconv.Atoi(elem)
		if value != 0 {
			got_numbers = append(got_numbers, value)
		}
	}
	fmt.Println(win_numbers)
	fmt.Println(got_numbers)
	matches := 0
	for _, win_num := range win_numbers {
		for _, got_num := range got_numbers {
			if win_num == got_num {
				matches++
			}
		}
	}
	fmt.Println(matches)
	if matches == 0 {
		return 0
	}
	value := int(math.Pow(2, float64(matches-1)))
	fmt.Println(value)
	return value
}

func split(r rune) bool {
	return r == '|' || r == ':'
}
