package days

import (
	"fmt"
	"regexp"
	"trox667.de/aoc/2024/tools"
)

type Day4 struct {
	Day
}

func (day *Day4) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day4) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day4) Run() {
	data := day.ReadInput(4)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 4 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 4 Part 2 failed")
	}
	println(result)
}

func (day *Day4) RunSample() {
	data := day.ReadSample(4)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 4 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 4 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day4) Part1(input []string) (result string, err error) {
	puzzle := make([][]rune, 0)
	for _, line := range input {
		puzzle = append(puzzle, []rune(line))
	}

	r, err := regexp.Compile("XMAS")
	if err != nil {
		panic("Could not compile regexp")
	}
	matches := countMatches(r, puzzle)
	matches += countMatches(r, rotate(puzzle))
	matches += checkDiagonal(puzzle)

	return fmt.Sprintf("Day 4 Part 1: %d", matches), nil
}

func (day *Day4) Part2(input []string) (result string, err error) {
	puzzle := make([][]rune, 0)
	for _, line := range input {
		puzzle = append(puzzle, []rune(line))
	}

	matches := 0
	for y, row := range puzzle {
		for x, _ := range row {
			if checkXMas(puzzle, x , y) {
				matches++
			}
		}
	}

	return fmt.Sprintf("Day 4 Part 1: %d", matches), nil
}

func countMatches(r *regexp.Regexp, puzzle [][]rune) (count int) {
	for _, row := range puzzle {
		count += len(r.FindAllString(string(row), -1))
		reverse(row)
		count += len(r.FindAllString(string(row), -1))
		reverse(row)
	}
	return count
}

func printPuzzle(puzzle [][]rune) {
	for _, row := range puzzle {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func reverse(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(puzzle [][]rune) [][]rune {
	rotated := make([][]rune, len(puzzle))
	for i := 0; i < len(puzzle); i++ {
		rotated[i] = make([]rune, len(puzzle[i]))
	}

	for i, _ := range puzzle {
		for j, _ := range puzzle[i] {
			rotated[j][i] = puzzle[i][j]
		}
	}
	return rotated
}

func checkDiagonal(puzzle [][]rune) int {
	count := 0
	for y, row := range puzzle {
		for x, _ := range row {
			count += checkNeighborsForRune(puzzle, x, y, 0, 0)
		}
	}
	return count
}

var word = [4]rune{'X', 'M', 'A', 'S'}

func checkNeighborsForRune(puzzle [][]rune, x int, y int, level int, direction int) int {
	if y < 0 || y >= len(puzzle) || x < 0 || x >= len(puzzle[0]) {
		return 0
	}

	if puzzle[y][x] != word[level] {
		return 0
	}

	if puzzle[y][x] == 'S' && level == 3 {
		fmt.Printf("%d,%d\n", x, y)
		return 1
	}

	count := 0
	uy := y - 1
	by := y + 1
	lx := x - 1
	rx := x + 1

	// top left
	if y-1 >= 0 && x-1 >= 0 && level < 3 && (direction == 0 || direction == 1)  {
		count += checkNeighborsForRune(puzzle, lx, uy, level+1, 1)
	}

	// top right
	if y-1 >= 0 && x+1 <= len(puzzle[0]) && level < 3 && (direction == 0 || direction == 2) {
		count += checkNeighborsForRune(puzzle, rx, uy, level+1, 2)
	}

	// bottom left
	if y+1 <= len(puzzle) && x-1 >= 0 && level < 3 && (direction == 0 || direction == 3) {
		count += checkNeighborsForRune(puzzle, lx, by, level+1, 3)
	}

	// bottom right
	if y+1 <= len(puzzle) && x+1 <= len(puzzle[0]) && level < 3 && (direction == 0 || direction == 4) {
		count += checkNeighborsForRune(puzzle, rx, by, level+1, 4)
	}

	return count
}

func checkXMas(puzzle [][]rune, x int, y int) bool {

	if puzzle[y][x] != 'A' {
		return false
	}

	if x-1 < 0 || y-1 < 0 || y+1 >= len(puzzle) || x+1 >= len(puzzle[0]) {
		return false
	}

	ls := puzzle[y-1][x-1] == 'M' && puzzle[y+1][x+1] == 'S'
	rs := puzzle[y+1][x-1] == 'M' && puzzle[y-1][x+1] == 'S'
	if ls && rs {
		return true
	}

	ls = puzzle[y-1][x-1] == 'S' && puzzle[y+1][x+1] == 'M'
	rs = puzzle[y+1][x-1] == 'S' && puzzle[y-1][x+1] == 'M'
	if ls && rs {
		return true
	}

	ls = puzzle[y-1][x-1] == 'S' && puzzle[y+1][x+1] == 'M'
	rs = puzzle[y+1][x-1] == 'M' && puzzle[y-1][x+1] == 'S'
	if ls && rs {
		return true
	}

	ls = puzzle[y-1][x-1] == 'M' && puzzle[y+1][x+1] == 'S'
	rs = puzzle[y+1][x-1] == 'S' && puzzle[y-1][x+1] == 'M'
	if ls && rs {
		return true
	}

	return false
}