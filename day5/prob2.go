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

type interval struct {
	start int
	end   int
	size  int
}

func main() {
	content, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(".+:\n") // split on the mapping strings
	split := re.Split(string(content), -1)
	seeds_intervals := get_seeds(split[0]) // extract tehn seeds from the first string
	total_mappings := [][]mapping{}        // matrix of all mappings in all sections
	for _, mapping_data := range split[1:] {
		new_mapping_list := get_mapping(mapping_data)
		total_mappings = append(total_mappings, new_mapping_list)
	}
	final_intervals := []interval{}                 // prepare for final results
	for _, seed_interval := range seeds_intervals { // take all input intervals
		curr_intervals := []interval{seed_interval}    // copy on first pass
		for _, mapping_level := range total_mappings { // for each level of mapping
			updated_intervals := []interval{}                // prepare space for next level's intervals
			for _, single_interval := range curr_intervals { // for each interval
				new_intervals := compute_intervals(single_interval, mapping_level) // compute the various intervals
				updated_intervals = append(updated_intervals, new_intervals...)    // store the results
			}
			curr_intervals = updated_intervals // now that we move to the next layer
		}
		final_intervals = append(final_intervals, curr_intervals...) // save the results from each seed interval
	}

	// find the lowest
	sort.Slice(final_intervals, func(i, j int) bool { return final_intervals[i].start < final_intervals[j].start })
	fmt.Print(final_intervals[0])

}
func compute_intervals(input_interval interval, level []mapping) []interval {
	for _, possible_map := range level {
		// map the value according to the mapping
		if input_interval.start >= possible_map.start && input_interval.start < possible_map.end {
			diff := input_interval.start - possible_map.start
			if input_interval.end < possible_map.end { // then the curent mapping covers the entire input
				new_interval := interval{start: possible_map.out_index + diff, end: possible_map.out_index + diff + input_interval.size, size: input_interval.size}
				return []interval{new_interval}
			} // otherwise we need to split it
			new_start := possible_map.out_index + diff
			new_size := possible_map.end - input_interval.start
			new_interval := interval{start: new_start, end: new_start + new_size, size: new_size}
			remaining_input_interval := interval{start: input_interval.start + new_size, end: input_interval.end, size: input_interval.size - new_size}
			// we can use the function recursively
			extra_intervals := compute_intervals(remaining_input_interval, level)
			// and combine the results
			return append(extra_intervals, new_interval)
		}

	}
	// if no mapping applieas we cna return teh orignal interval, as it uses the identity mapping
	return []interval{input_interval}

}

func get_seeds(input string) []interval {
	// read the first line
	section_numbers := strings.Split(input, ":")[1]
	clean_section_numbers := strings.TrimSpace(section_numbers)
	// get the number list
	str_seeds := strings.Split(clean_section_numbers, " ")
	seeds := []int{}
	// convert to Ints
	for _, seed := range str_seeds {
		if seed == "" {
			continue
		}
		value, _ := strconv.Atoi(seed)
		seeds = append(seeds, value)
	}
	// now pair up into intervals
	new_seeds := []interval{}
	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		length := seeds[i+1]
		new_seeds = append(new_seeds, interval{start: start, end: start + length, size: length})
	}

	return new_seeds
}

func get_mapping(input string) []mapping {
	// read each line and convert into mappings
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
