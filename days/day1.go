package days

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"trox667.de/aoc/2024/tools"
)

type Day1 struct {
	Day
}

func (day *Day1) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day1) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day1) Run() {
	data := day.ReadInput(1)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 1 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 1 Part 2 failed")
	}
	println(result)
}

func (day *Day1) RunSample() {
	data := day.ReadSample(1)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 1 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 1 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day1) Part1(input []string) (result string, err error) {
	firstList := make([]int, 0)
	secondList := make([]int, 0)
	for _, line := range input {
		tokens := strings.Split(line, "   ")
		if len(tokens) != 2 {
			panic("could not split input")
		}
		firstVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic("Could not convert input to int")
		}
		secondVal, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic("Could not convert input to int")
		}

		firstList = append(firstList, firstVal)
		secondList = append(secondList, secondVal)
	}

	slices.SortFunc(firstList, func(a, b int) int {
		return a - b
	})

	slices.SortFunc(secondList, func(a, b int) int {
		return a - b
	})

	distance := 0.0
	for i := 0; i < len(firstList); i++ {
		distance += math.Abs(float64(firstList[i] - secondList[i]))
	}

	return fmt.Sprintf("Day 1 Part 1: %d", int(distance)), nil
}

func (day *Day1) Part2(input []string) (result string, err error) {
	firstList := make([]int, 0)
	secondList := make([]int, 0)
	for _, line := range input {
		tokens := strings.Split(line, "   ")
		if len(tokens) != 2 {
			panic("could not split input")
		}
		firstVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic("Could not convert input to int")
		}
		secondVal, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic("Could not convert input to int")
		}

		firstList = append(firstList, firstVal)
		secondList = append(secondList, secondVal)
	}

	score := 0
	for _, firstVal := range firstList {
		count := 0
		for _, secondVal := range secondList {
			if firstVal == secondVal {
				count++
			}
		}
		score += firstVal * count
	}

	return fmt.Sprintf("Day 1 Part 2: %d", score), nil
}
