package main

import "trox667.de/aoc/2024/days"

func main() {
	println("Advent of Code 2024")

	day1 := days.Day1{}
	data := day1.ReadInput(1)
	result, err := day1.Part1(data)
	if err != nil {
		panic("Day 1 Part 1 failed")
	}
	println(result)

	result, err = day1.Part2(data)
	if err != nil {
		panic("Day 1 Part 2 failed")
	}
	println(result)
}
