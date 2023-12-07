package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Helps to keep the reasoning clear
type mapping struct {
	start     int
	end       int
	out_index int
}

func main() {
	content, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(".+:\n") // split on the mapping strings
	split := re.Split(string(content), -1)
	seeds := get_seeds(split[0])    // extract tehn seeds from the first string
	total_mappings := [][]mapping{} // matrix of all mappings in all sections
	for _, mapping_data := range split[1:] {
		new_mapping_list := get_mapping(mapping_data)
		total_mappings = append(total_mappings, new_mapping_list)
	}
	final_values := []int{}
	for _, seed := range seeds {
		value := seed // we pass each seed through each layer of the mapping,
		for _, mapping_level := range total_mappings {
			for _, possible_map := range mapping_level {
				// map the value according to the mapping
				if value >= possible_map.start && value < possible_map.end {
					diff := value - possible_map.start
					value = possible_map.out_index + diff
					break
				}
			}
		}
		final_values = append(final_values, value)
	}
	// find the lowest
	sort.Ints(final_values)
	fmt.Print(final_values[0])

}
func get_seeds(input string) []int {
	section_numbers := strings.Split(input, ":")[1]
	clean_section_numbers := strings.TrimSpace(section_numbers)
	str_seeds := strings.Split(clean_section_numbers, " ")
	seeds := []int{}

	for _, seed := range str_seeds {
		if seed == "" {
			continue
		}
		value, _ := strconv.Atoi(seed)
		seeds = append(seeds, value)
	}
	return seeds
}

func get_mapping(input string) []mapping {
	lines := strings.Split(input, "\n")
	mappings := []mapping{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		line_elems := strings.Split(line, " ")
		start, _ := strconv.Atoi(line_elems[1])
		new_index, _ := strconv.Atoi(line_elems[0])
		size, _ := strconv.Atoi(line_elems[2])
		new_map := mapping{start: start, end: start + size, out_index: new_index}
		mappings = append(mappings, new_map)
	}
	return mappings
}
