package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"trox667.de/aoc/2024/tools"
)

type Day2 struct {
	Day
}

func (day Day2) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day Day2) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day Day2) Run() {
	data := day.ReadInput(2)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 2 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 2 Part 2 failed")
	}
	println(result)
}

func (day Day2) RunSample() {
	data := day.ReadSample(2)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 2 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 2 Part 2 Sample failed")
	}
	println(result)
}

func (day Day2) Part1(input []string) (result string, err error) {
	safeReports := make([]bool, 0)
	for _, report := range input {
		tokens := strings.Split(report, " ")
		reportData := make([]int, 0)
		for _, level := range tokens {
			level, err := strconv.Atoi(level)
			if err != nil {
				panic("Could not convert to int")
			}
			reportData = append(reportData, level)
		}

		if checkSafety(reportData) {
			safeReports = append(safeReports, true)
		}
	}

	return fmt.Sprintf("Day 2 Part 1: %d", len(safeReports)), nil
}

func (day Day2) Part2(input []string) (result string, err error) {
	safeReports := make([]bool, 0)
	for _, report := range input {
		tokens := strings.Split(report, " ")
		reportData := make([]int, 0)
		for _, level := range tokens {
			level, err := strconv.Atoi(level)
			if err != nil {
				panic("Could not convert to int")
			}
			reportData = append(reportData, level)
		}

		for i := 0; i < len(reportData); i++ {
			if checkSafety(reportData) {
				safeReports = append(safeReports, true)
				break
			}
			temp := make([]int, 0)
			if i == 0 {
				temp = append(temp, reportData[0:i]...)
				temp = append(temp, reportData[i+1:]...)
			} else {
				temp = append(temp, reportData[:i]...)
				temp = append(temp, reportData[i+1:]...)
			}
			if checkSafety(temp) {
				safeReports = append(safeReports, true)
				break
			}
		}
	}

	return fmt.Sprintf("Day 2 Part 2: %d", len(safeReports)), nil
}

func checkSafety(reportData []int) bool {
	gradient := 0
	safeReport := true
	for i := 1; i < len(reportData); i++ {
		prev := reportData[i-1]
		curr := reportData[i]
		slope := curr - prev
		absSlope := math.Abs(float64(slope))

		if gradient == 0 {
			if slope > 0 {
				gradient = 1
			}
			if slope < 0 {
				gradient = -1
			}
		} else {
			if gradient > 0 && slope < 0 {
				safeReport = false
				break
			} else if gradient < 0 && slope > 0 {
				safeReport = false
				break
			}
		}

		if absSlope > 3 || slope == 0 {
			safeReport = false
			break
		}
	}
	return safeReport
}
