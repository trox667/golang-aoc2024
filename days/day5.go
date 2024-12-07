package days

import (
	"fmt"
	"strconv"
	"strings"
	"trox667.de/aoc/2024/tools"
)

type PageRules struct {
	parents, children map[int][]int
}

type PageUpdates struct {
	pages []int
}

func NewPageRules() PageRules {
	childrenMap := make(map[int][]int)
	parentMap := make(map[int][]int)
	return PageRules{parentMap, childrenMap}
}

type Day5 struct {
	Day
}

func (day *Day5) ReadSample(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/sample%d", dayIndex))
}

func (day *Day5) ReadInput(dayIndex int8) []string {
	return tools.ReadInput(fmt.Sprintf("./inputs/input%d", dayIndex))
}

func (day *Day5) Run() {
	data := day.ReadInput(5)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 5 Part 1 failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 5 Part 2 failed")
	}
	println(result)
}

func (day *Day5) RunSample() {
	data := day.ReadSample(5)
	result, err := day.Part1(data)
	if err != nil {
		panic("Day 5 Part 1 Sample failed")
	}
	println(result)

	result, err = day.Part2(data)
	if err != nil {
		panic("Day 5 Part 2 Sample failed")
	}
	println(result)
}

func (day *Day5) Part1(input []string) (result string, err error) {
	orderingRules := NewPageRules()
	pageUpdates := make([]PageUpdates, 0)
	firstSection := true
	for _, line := range input {
		if len(line) == 0 {
			firstSection = false
			continue
		}

		if firstSection {
			tokens := strings.Split(line, "|")
			a, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic("Could not parse page rule")
			}
			b, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("Could not parse page rule")
			}
			orderingRules.parents[b] = append(orderingRules.parents[b], a)
			orderingRules.children[a] = append(orderingRules.children[a], b)
		}
		if !firstSection {
			pageUpdate := PageUpdates{}
			pageUpdate.pages = make([]int, 0)
			for _, token := range strings.Split(line, ",") {
				val, err := strconv.Atoi(token)
				if err != nil {
					panic("Could not parse page update")
				}
				pageUpdate.pages = append(pageUpdate.pages, val)
			}
			pageUpdates = append(pageUpdates, pageUpdate)
		}
	}

	//fmt.Printf("%#v\n", orderingRules)
	//fmt.Printf("%#v\n", pageUpdates)

	notMatchingUpdates := make([]PageUpdates, 0)

	for _, pageUpdate := range pageUpdates {
		match := validateUpdates(pageUpdate, orderingRules)
		if match {
			notMatchingUpdates = append(notMatchingUpdates, pageUpdate)
		}
		//fmt.Printf("match %t, %#v\n", match, pageUpdate)
	}

	sumMiddleNumber := 0
	for _, matchingUpdate := range notMatchingUpdates {

		sumMiddleNumber += middleNumber(matchingUpdate)
	}

	return fmt.Sprintf("Day 5 Part 1: %d", sumMiddleNumber), nil
}

func (day *Day5) Part2(input []string) (result string, err error) {
	orderingRules := NewPageRules()
	pageUpdates := make([]PageUpdates, 0)
	firstSection := true
	for _, line := range input {
		if len(line) == 0 {
			firstSection = false
			continue
		}

		if firstSection {
			tokens := strings.Split(line, "|")
			a, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic("Could not parse page rule")
			}
			b, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("Could not parse page rule")
			}
			orderingRules.parents[b] = append(orderingRules.parents[b], a)
			orderingRules.children[a] = append(orderingRules.children[a], b)
		}
		if !firstSection {
			pageUpdate := PageUpdates{}
			pageUpdate.pages = make([]int, 0)
			for _, token := range strings.Split(line, ",") {
				val, err := strconv.Atoi(token)
				if err != nil {
					panic("Could not parse page update")
				}
				pageUpdate.pages = append(pageUpdate.pages, val)
			}
			pageUpdates = append(pageUpdates, pageUpdate)
		}
	}

	//fmt.Printf("%#v\n", orderingRules)
	//fmt.Printf("%#v\n", pageUpdates)

	matchingUpdates := make([]PageUpdates, 0)

	for _, pageUpdate := range pageUpdates {
		match := validateUpdates(pageUpdate, orderingRules)
		if !match {
			matchingUpdates = append(matchingUpdates, pageUpdate)
		}
	}

	sumMiddleNumber := 0
	for _, matchingUpdate := range matchingUpdates {
		sort(matchingUpdate, orderingRules)
		sumMiddleNumber += middleNumber(matchingUpdate)
	}

	return fmt.Sprintf("Day 5 Part 2: %d", sumMiddleNumber), nil
}

func middleNumber(pageUpdates PageUpdates) int {
	i := len(pageUpdates.pages) / 2
	return pageUpdates.pages[i]
}

func validateUpdates(pageUpdates PageUpdates, rules PageRules) bool {
	for i, _ := range pageUpdates.pages {
		match := validateUpdate(i, pageUpdates, rules)
		if !match {
			return false
		}
	}
	return true
}

func validateUpdate(pageIndex int, pageUpdates PageUpdates, rules PageRules) bool {
	left := pageUpdates.pages[:pageIndex]
	right := pageUpdates.pages[pageIndex+1:]
	curr := pageUpdates.pages[pageIndex]

	hasPrev, hasNext := true, true
	if len(left) > 0 {
		prev := left[len(left)-1]
		parents := rules.parents[curr]
		prevMatch := false
		for _, parent := range parents {
			if prev == parent {
				prevMatch = true
				break
			}
		}
		hasPrev = prevMatch
	}

	if len(right) > 0 {
		next := right[0]
		children := rules.children[curr]
		nextMatch := false
		for _, child := range children {
			if next == child {
				nextMatch = true
				break
			}
		}
		hasNext = nextMatch
	}
	return hasPrev && hasNext
}

func sort(pageUpdates PageUpdates, rules PageRules) {
	for i := len(pageUpdates.pages) - 1; i > 0; i-- {
		for pageIndex, _ := range pageUpdates.pages {
			//left := pageUpdates.pages[:pageIndex]
			right := pageUpdates.pages[pageIndex+1:]
			curr := pageUpdates.pages[pageIndex]

			//if len(left) > 0 {
			//	prev := left[len(left)-1]
			//	parents := rules.parents[curr]
			//	for _, parent := range parents {
			//		if prev == parent {
			//			break
			//		}
			//	}
			//}

			if len(right) > 0 {
				next := right[0]
				parents := rules.parents[curr]
				for _, parent := range parents {
					if next == parent {
						pageUpdates.pages[pageIndex], pageUpdates.pages[pageIndex+1] = pageUpdates.pages[pageIndex+1], pageUpdates.pages[pageIndex]
						break
					}
				}
			}
		}
	}
}