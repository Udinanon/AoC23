package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time   int
	record int
}

func main() {
	content, err := os.ReadFile("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	str_times := strings.Fields(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	str_records := strings.Fields(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	races := []race{}
	for i := range str_times {
		time, _ := strconv.Atoi(str_times[i])
		record, _ := strconv.Atoi(str_records[i])
		races = append(races, race{time: time, record: record})
	}

	result := 1.
	// there is an analytical solution cause this is a quadratic formula
	for _, race := range races {
		interval_end := (float64(race.time) + math.Sqrt(float64(race.time)*float64(race.time)-(4*float64(race.record)))) / 2.
		interval_begin := (float64(race.time) - math.Sqrt(float64(race.time)*float64(race.time)-(4*float64(race.record)))) / 2.
		fmt.Println(interval_begin, interval_end)
		begin := math.Ceil(interval_begin + 0.0001)
		end := math.Floor(interval_end - 0.0001)
		size := (end - begin) + 1
		fmt.Println(begin, end, size)
		result *= size
	}

	fmt.Printf("%f\n", result)

}
