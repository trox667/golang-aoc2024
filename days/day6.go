package days

import (
	"fmt"
	"trox667.de/aoc/2024/tools"
)

type Day6 struct {
	Day
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type GuardMap struct {
	start     tools.Position
	path      map[tools.Position]bool
	pathList  []tools.Position
	obstacles map[tools.Position]bool
	width     int
	height    int
}

func EmptyGuardMap() GuardMap {
	return GuardMap{
		start:     tools.NewPosition(0, 0),
		path:      make(map[tools.Position]bool),
		pathList:  make([]tools.Position, 0),
		obstacles: make(map[tools.Position]bool),
		width:     0,
		height:    0,
	}
}

func (day *Day6) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day6) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day6) Run() {
	data := day.ReadInput(6)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 6 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 6 Part 2 failed")
	}
	println(result)
}

func (day *Day6) RunSample() {
	data := day.ReadSample(6)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 6 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 6 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day6) Part1(input []string) (result string, err error) {
	guardMap := createGuardMap(input)
	//printGuardMap(guardMap)
	runGuard(guardMap)

	return fmt.Sprintf("Day 6 Part 1: %d", len(guardMap.path)), nil
}

func (day *Day6) Part2(input []string) (result string, err error) {
	guardMap := createGuardMap(input)

	guardMap = runGuard(guardMap)
	options := runGuard2(guardMap)
	return fmt.Sprintf("Day 6 Part 2: %d", options), nil
}

func runGuard(guardMap GuardMap) GuardMap {
	position := guardMap.start
	direction := Up

	for {
		if position.X < 0 || position.X >= guardMap.width || position.Y < 0 || position.Y >= guardMap.height {
			break
		}

		if !guardMap.path[position] {
			guardMap.pathList = append(guardMap.pathList, position)
		}
		guardMap.path[position] = true

		if direction == Up && guardMap.obstacles[tools.NewPosition(position.X, position.Y-1)] {
			direction = Right
		} else if direction == Right && guardMap.obstacles[tools.NewPosition(position.X+1, position.Y)] {
			direction = Down
		} else if direction == Down && guardMap.obstacles[tools.NewPosition(position.X, position.Y+1)] {
			direction = Left
		} else if direction == Left && guardMap.obstacles[tools.NewPosition(position.X-1, position.Y)] {
			direction = Up
		}

		if direction == Up {
			position.Y -= 1
		} else if direction == Right {
			position.X += 1
		} else if direction == Down {
			position.Y += 1
		} else {
			position.X -= 1
		}
	}

	return guardMap
}

func runGuard2(guardMap GuardMap) int {
	options := make(map[tools.Position]bool)

	for _, p := range guardMap.pathList {
		position := tools.NewPosition(guardMap.start.X, guardMap.start.Y)
		direction := Up
		path := make(map[tools.Position]int)
		obstacles := make(map[tools.Position]bool)
		for k, v := range guardMap.obstacles {
			obstacles[k] = v
		}
		obstacles[p] = true

		for {
			// already used this obstacle
			if options[p] {
				break
			}

			// guard moved out of bounds
			if position.X < 0 || position.X >= guardMap.width || position.Y < 0 || position.Y >= guardMap.height {
				break
			}

			path[position] += 1

			if path[position] > 10 {
				options[p] = true
			}

			if direction == Up && obstacles[tools.NewPosition(position.X, position.Y-1)] {
				direction = Right
			} else if direction == Right && obstacles[tools.NewPosition(position.X+1, position.Y)] {
				direction = Down
			} else if direction == Down && obstacles[tools.NewPosition(position.X, position.Y+1)] {
				direction = Left
			} else if direction == Left && obstacles[tools.NewPosition(position.X-1, position.Y)] {
				direction = Up
			}

			if direction == Up {
				position.Y -= 1
			} else if direction == Right {
				position.X += 1
			} else if direction == Down {
				position.Y += 1
			} else {
				position.X -= 1
			}
		}
	}
	//fmt.Printf("%#v", options)
	return len(options)
}

func createGuardMap(input []string) GuardMap {
	guardMap := EmptyGuardMap()
	guardMap.height = len(input)
	guardMap.width = len(input[0])
	for y, row := range input {
		for x, cell := range row {
			if cell == '#' {
				guardMap.obstacles[tools.NewPosition(x, y)] = true
			} else if cell == '^' {
				guardMap.start = tools.NewPosition(x, y)
			}
		}
	}
	return guardMap
}

func printGuardMap(guardMap GuardMap) {
	for y := 0; y < guardMap.height; y++ {
		for x := 0; x < guardMap.width; x++ {
			value := guardMap.obstacles[tools.NewPosition(x, y)]
			if value {
				fmt.Print("#")
			} else if guardMap.start.X == x && guardMap.start.Y == y {
				fmt.Print("^")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
