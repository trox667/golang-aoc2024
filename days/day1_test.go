package days

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	day1 := Day1{}
	input := []string{"A", "B"}
	result, err := day1.Part1(input)
	if result != "Day 1 Part 1" || err != nil {
		t.Error("Part 1 does not match required result")
	}
}
