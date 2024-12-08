package days

import (
	"fmt"
	"trox667.de/aoc/2024/tools"
)

type AntennaMap struct {
	Antennas      map[rune][]tools.Position
	Width, Height int
}

func NewAntennaMap() AntennaMap {
	return AntennaMap{make(map[rune][]tools.Position), 0, 0}
}

type Day8 struct {
	Day
}

func (day *Day8) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day8) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day8) Run() {
	data := day.ReadInput(8)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 8 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 8 Part 2 failed")
	}
	println(result)
}

func (day *Day8) RunSample() {
	data := day.ReadSample(8)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 8 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 8 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day8) Part1(input []string) (result string, err error) {
	score := 0

	antinodeMap := make(map[tools.Position]bool)
	antennaMap := createAntennaMap(input)
	for k, _ := range antennaMap.Antennas {
		positions := calculateAntinodeForAntenna(k, &antennaMap)
		for _, pos := range positions {
			antinodeMap[pos] = true
		}
	}

	score = len(antinodeMap)

	return fmt.Sprintf("Day 8 Part 1: %d", score), nil
}

func (day *Day8) Part2(input []string) (result string, err error) {
	score := 0

	antinodeMap := make(map[tools.Position]bool)
	antennaMap := createAntennaMap(input)
	for k, _ := range antennaMap.Antennas {
		positions := calculateAntinodeForAntenna2(k, &antennaMap)
		for _, antennaPos := range antennaMap.Antennas[k] {
			antinodeMap[antennaPos] = true
		}
		for _, pos := range positions {
			antinodeMap[pos] = true
		}
	}

	score = len(antinodeMap)
	return fmt.Sprintf("Day 8 Part 2: %d", score), nil
}

func createAntennaMap(lines []string) AntennaMap {
	antennaMap := NewAntennaMap()

	for y, row := range lines {
		for x, col := range row {
			if col != '.' {
				pos := tools.NewPosition(x, y)
				antennaMap.Antennas[col] = append(antennaMap.Antennas[col], pos)
			}
		}
	}
	if len(lines) > 0 {
		antennaMap.Height = len(lines)
		antennaMap.Width = len(lines[0])
	}
	return antennaMap
}

func calculateAntinodeForAntenna(antenna rune, antennaMap *AntennaMap) []tools.Position {
	positions := antennaMap.Antennas[antenna]
	antinodePositions := make([]tools.Position, 0)

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions); j++ {
			if i == j {
				continue
			}

			a := antinode(positions[i], positions[j])
			if a.X < 0 || a.X >= antennaMap.Width || a.Y < 0 || a.Y >= antennaMap.Height {
				// foo
			} else {
				antinodePositions = append(antinodePositions, a)
			}
		}
	}
	return antinodePositions
}

func antinode(a, b tools.Position) tools.Position {
	c := tools.NewPosition(a.X, a.Y)
	c.X += -1 * (b.X - a.X)
	c.Y += -1 * (b.Y - a.Y)
	return c
}

func calculateAntinodeForAntenna2(antenna rune, antennaMap *AntennaMap) []tools.Position {
	positions := antennaMap.Antennas[antenna]
	antinodePositions := make([]tools.Position, 0)

	for i := 0; i < len(positions); i++ {
		for j := 0; j < len(positions); j++ {
			if i == j {
				continue
			}

			a := antinode2(positions[i], positions[j], antennaMap.Width, antennaMap.Height)
			antinodePositions = append(antinodePositions, a...)
		}
	}
	return antinodePositions
}

func antinode2(a, b tools.Position, width, height int) []tools.Position {
	positions := make([]tools.Position, 0)
	c := tools.NewPosition(a.X, a.Y)
	for {
		c.X += -1 * (b.X - a.X)
		c.Y += -1 * (b.Y - a.Y)

		if c.X < 0 || c.X >= width || c.Y < 0 || c.Y >= height {
			break
		} else {
			positions = append(positions, c)
		}
	}
	return positions
}