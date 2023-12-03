package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	value   int
	visited bool
}

type location struct {
	line int
	char int
}

func main() {

	var total int = 0

	content, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	n_lines := len(lines)
	size_line := len(lines[0])
	symbols_positions := []location{}

	values := make([][]*point, n_lines)
	for i := 0; i < n_lines; i++ {
		values[i] = make([]*point, size_line)
	}

	for line_index, line := range lines {
		var temp string
		var start int
		values_line := &values[line_index]
		for char_index, char := range line {
			is_number := strings.ContainsRune("0123456789", char)
			if is_number {
				if temp == "" {
					start = char_index
				}
				temp += string(char)
			} else {
				if char != rune('.') {
					position := location{line_index, char_index}
					symbols_positions = append(symbols_positions, position)
				}
				if temp != "" {
					//Start: start, End: i - 1
					saveTemp(temp, values_line, start, char_index)
					temp = ""
				}
				new_point := &point{
					value:   0,
					visited: false,
				}
				(*values_line)[char_index] = new_point
			}
		}
		if temp != "" {
			saveTemp(temp, values_line, start, len(line))
			temp = ""
		}
		//fmt.Println(values_line)
	}
	//	for _, line := range values {
	//		for _, point := range line {
	//			//fmt.Println(point.value)
	//		}
	//
	//	}
	//fmt.Println(symbols_positions)

	for _, position := range symbols_positions {
		for _, i := range []int{-1, 0, 1} {
			for _, j := range []int{-1, 0, 1} {
				curr_line := position.line + i
				curr_char := position.char + j
				valid_line := (curr_line) >= 0 && (curr_line) < n_lines
				valid_char := (curr_char) >= 0 && (curr_char) < size_line

				if valid_line && valid_char {
					curr_value := values[curr_line][curr_char]
					if !curr_value.visited {
						fmt.Println(curr_value.value)
						total += curr_value.value
						curr_value.visited = true
					}
				}
			}
		}
	}

	fmt.Printf("TOTAL: %d", total)
}

func saveTemp(temp string, array *[]*point, start int, end int) {
	temp_value, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println("ERROR IN CONVERSION: ", err)
	}
	new_point := &point{
		value:   temp_value,
		visited: false,
	}
	//Start: start, End: len(line) - 1
	for k := start; k < end; k++ {
		(*array)[k] = new_point
	}
	temp = ""
}
