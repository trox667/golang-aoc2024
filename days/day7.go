package days

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"trox667.de/aoc/2024/tools"
)

type Day7 struct {
	Day
}

type CalibrationEquation struct {
	Result int
	Values []int
}

type CalcResult struct {
	CalibrationEquation CalibrationEquation
	Result              bool
}

func (day *Day7) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day7) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day7) Run() {
	data := day.ReadInput(7)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 7 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 7 Part 2 failed")
	}
	println(result)
}

func (day *Day7) RunSample() {
	data := day.ReadSample(7)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 7 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 7 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day7) Part1(input []string) (result string, err error) {
	score := 0

	calibrationEquations := make([]CalibrationEquation, 0)
	for _, line := range input {
		calibrationEquations = append(calibrationEquations, createCalibrationEquation(line))
	}

	for _, calibrationEquation := range calibrationEquations {
		valid := calculateWithCombinations(calibrationEquation)
		if valid {
			score += calibrationEquation.Result
		}
	}

	return fmt.Sprintf("Day 7 Part 1: %d", score), nil
}

func (day *Day7) Part2(input []string) (result string, err error) {
	score := 0

	calibrationEquations := make([]CalibrationEquation, 0)
	for _, line := range input {
		calibrationEquations = append(calibrationEquations, createCalibrationEquation(line))
	}

	var wg sync.WaitGroup
	results := make(chan CalcResult, len(calibrationEquations))

	for _, calibrationEquation := range calibrationEquations {
		wg.Add(1)
		go calculateWithCombinations2(calibrationEquation, results, &wg)
	}
	wg.Wait()
	close(results)
	for result := range results {
		if result.Result {
			score += result.CalibrationEquation.Result
		}
	}
	return fmt.Sprintf("Day 7 Part 2: %d", score), nil
}

func createCalibrationEquation(line string) CalibrationEquation {
	tokens := strings.Split(line, ":")
	result, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic("Could not convert to int")
	}
	valueTokens := strings.Split(strings.Trim(tokens[1], " "), " ")
	values := make([]int, 0)
	for _, valueToken := range valueTokens {
		value, err := strconv.Atoi(valueToken)
		if err != nil {
			panic("Could not convert to int")
		}
		values = append(values, value)
	}

	return CalibrationEquation{
		Result: result,
		Values: values,
	}
}

func calculateWithCombinations(equation CalibrationEquation) bool {
	possibleOptions := make([][]int, 0)
	values := equation.Values
	result := equation.Result

	operators := []int{0, 1}

	applyOperator := func(op int, a int, b int) int {
		if op == 0 {
			return a + b
		} else {
			return a * b
		}
	}

	options := AllRepeat(operators, len(values)-1)
	//fmt.Printf("%#v\n", equation)
	//fmt.Printf("%#v\n", options)
	//fmt.Println()

	for _, option := range options {
		r := values[0]
		for i := 1; i < len(values); i++ {
			r = applyOperator(option[i-1], r, values[i])
		}

		if r == result {
			//fmt.Printf("Hit: %#v\n", option)
			possibleOptions = append(possibleOptions, option)
		}
	}
	if len(possibleOptions) > 0 {
		return true
	}
	return false
}

func calculateWithCombinations2(equation CalibrationEquation, results chan CalcResult, wg *sync.WaitGroup) {
	defer wg.Done()
	possibleOptions := make([][]int, 0)
	values := equation.Values
	result := equation.Result

	operators := []int{0, 1, 2}

	applyOperator := func(op int, a int, b int) int {
		if op == 0 {
			return a + b
		} else if op == 1 {
			return a * b
		} else {
			val, err := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
			if err != nil {
				panic("Could not combine numbers to single int")
			}
			return val
		}
	}

	options := AllRepeat(operators, len(values)-1)
	//fmt.Printf("%#v\n", equation)
	//fmt.Printf("%#v\n", options)
	//fmt.Println()

	for _, option := range options {
		r := values[0]
		for i := 1; i < len(values); i++ {
			r = applyOperator(option[i-1], r, values[i])
		}

		if r == result {
			//fmt.Printf("Hit: %#v\n", option)
			possibleOptions = append(possibleOptions, option)
		}
	}
	if len(possibleOptions) > 0 {
		results <- CalcResult{equation, true}
		return
	}
	results <- CalcResult{equation, false}
}

func AllRepeat[T any](set []T, m int) (subsets [][]T) {
	if m < 1 {
		return nil
	}

	var generateCombos func([]T, int)
	generateCombos = func(current []T, depth int) {
		if depth == 0 {
			subset := make([]T, len(current))
			copy(subset, current)
			subsets = append(subsets, subset)
			return
		}

		for _, item := range set {
			generateCombos(append(current, item), depth-1)
		}
	}

	//for length := 1; length <= m; length++ {
	generateCombos([]T{}, m)
	//}

	return subsets
}
