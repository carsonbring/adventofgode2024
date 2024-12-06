package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FindIndex(slice []int, element int) int {
	for index, value := range slice {
		if value == element {
			return index
		}
	}
	return -1
}

func MiddleIndex(slice []int) int {
	if len(slice) == 1 {
		return 0
	} else {
		return (len(slice) - 1) / 2
	}
}

func remove(slice []int, index int) []int {
	newList := make([]int, len(slice)-1)
	for i := 0; i < index; i++ {
		newList[i] = slice[i]
	}
	for i := index; i < len(slice)-1; i++ {
		newList[i] = slice[i+1]
	}
	return newList
}

func ReadInput() (map[int][]int, [][]int, error) {
	orderMap := make(map[int][]int)
	var pageUpdates [][]int

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			var ruleList []int

			line = strings.ReplaceAll(line, "|", " ")
			for _, s := range strings.Fields(line) {
				newInt, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, err
				}
				ruleList = append(ruleList, newInt)
			}
			if len(ruleList) != 2 {
				return nil, nil, fmt.Errorf("incorrect length of rules (must be 2)")
			}
			orderMap[ruleList[1]] = append(orderMap[ruleList[1]], ruleList[0])
		} else if strings.Contains(line, ",") {
			var pageList []int
			line = strings.ReplaceAll(line, ",", " ")
			for _, s := range strings.Fields(line) {
				newInt, err := strconv.Atoi(s)
				if err != nil {
					return nil, nil, err
				}
				pageList = append(pageList, newInt)
			}
			pageUpdates = append(pageUpdates, pageList)
		}
	}

	return orderMap, pageUpdates, nil
}

func Part1() (int, error) {
	sum := 0
	orderMap, pageUpdates, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, update := range pageUpdates {
		safeUpdate := true
		for i, num := range update {
			beforeNums := orderMap[num]
			for _, prereq := range beforeNums {
				if prereqIndex := FindIndex(update, prereq); prereqIndex != -1 && prereqIndex > i {
					safeUpdate = false
				}
			}

		}
		if safeUpdate {
			sum = sum + update[(len(update)-1)/2]
		}
	}
	return sum, nil
}

func Part2() (int, error) {
	sum := 0
	orderMap, pageUpdates, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, update := range pageUpdates {
		safeUpdate := true
		for {
			retry := false
			for i, num := range update {
				beforeNums := orderMap[num]
				for _, prereq := range beforeNums {
					if prereqIndex := FindIndex(update, prereq); prereqIndex != -1 && prereqIndex > i {
						safeUpdate = false
						update = remove(update, prereqIndex)
						update = append([]int{prereq}, update...)
						retry = true
					}
				}
			}
			if !retry {
				// fmt.Printf("%v \n", update)
				break
			}
		}
		if !safeUpdate {
			sum = sum + update[(len(update)-1)/2]
		}
	}
	return sum, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part 1: %v", err)
	}
	fmt.Printf("Sum of legit middle pages: %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part 2: %v", err)
	}
	fmt.Printf("Sum of illegit middle pages: %v \n", result2)
}
