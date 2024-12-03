package days

import (
	"fmt"
	"regexp"
	"strconv"

	"trox667.de/aoc/2024/tools"
)

type Day3 struct {
	Day
}

func (day *Day3) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day3) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day3) Run() {
	data := day.ReadInput(3)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 3 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 3 Part 2 failed")
	}
	println(result)
}

func (day *Day3) RunSample() {
	data := day.ReadSample(3)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 3 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 3 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day3) Part1(input []string) (result string, err error) {
	score := 0
	r, err := regexp.Compile("mul\\((\\d*),(\\d*)\\)")
	if err != nil {
		panic("Could not create regex")
	}
	for _, line := range input {
		result := r.FindAllStringSubmatch(line, -1)
		for _, matches := range result {
			a, b := 0, 0
			if len(matches) == 3 {
				a, err = strconv.Atoi(matches[1])
				if err != nil {
					panic("Failed int conversion")
				}
				b, err = strconv.Atoi(matches[2])
				if err != nil {
					panic("Failed int conversion")
				}
			}
			score += a * b
		}
	}

	return fmt.Sprintf("Day 3 Part 1: %d", score), nil
}

func (day *Day3) Part2(input []string) (result string, err error) {
	score := 0
	r, err := regexp.Compile("(do\\(\\))|(don't\\(\\))|mul\\((\\d*),(\\d*)\\)")
	skip := false
	for _, line := range input {
		result := r.FindAllStringSubmatch(line, -1)
		for _, matches := range result {
			if len(matches) != 5 {
				panic("No valid regex matching")
			}
			if matches[0] == "don't()" {
				skip = true
			} else if matches[0] == "do()" {
				skip = false
			} else if !skip {
				a, b := 0, 0
				a, err = strconv.Atoi(matches[3])
				if err != nil {
					panic("Failed int conversion")
				}
				b, err = strconv.Atoi(matches[4])
				if err != nil {
					panic("Failed int conversion")
				}
				score += a * b
			}
		}
	}

	return fmt.Sprintf("Day 3 Part 2: %d", score), nil
}
