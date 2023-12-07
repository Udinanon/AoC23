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

	multipliers := [204]int{}
	for i := 0; i < 204; i++ {
		multipliers[i] = 1
	}
	for scanner.Scan() {
		parseLine2(scanner.Text(), &multipliers)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return result
	total := 0
	fmt.Println(multipliers)

	for _, mult := range multipliers {
		total += mult
	}
	fmt.Printf("TOTAL: %d", total)
}

func parseLine2(line string, multipliers *[204]int) {
	split_line := strings.FieldsFunc(line, split)
	fmt.Println(split_line)
	card_id, _ := strconv.Atoi(strings.Split(split_line[0], " ")[len(strings.Split(split_line[0], " "))-1])
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
	fmt.Println("matches: ", matches)
	if matches != 0 {

		//value := int(math.Pow(2, float64(matches-1)))
		curr_mult := (*multipliers)[card_id-1]
		fmt.Println("CURR MUL: ", curr_mult)
		for i := 0; i < matches; i++ {
			index := card_id + i
			if index < len(*multipliers) {
				(*multipliers)[index] += (curr_mult)
			}
		}
	}

}

func split(r rune) bool {
	return r == '|' || r == ':'
}
