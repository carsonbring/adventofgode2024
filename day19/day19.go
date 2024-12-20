package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() ([]string, []string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	types := []string{}
	patterns := []string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if rowCount == 0 {
			types = strings.Split(line, ",")
		} else {
			if line != "" {
				patterns = append(patterns, line)
			}
		}
		rowCount++
	}
	return types, patterns, nil
}

// JUST USE RECURSION DUMMY!!!!!!!!!!!!!!!!!!!!!!!!!!!
func FindPattern(design string, typeMap map[int][]string, maxLength int, memoMap map[string]bool) bool {
	startLength := 0
	if found, exists := memoMap[design]; exists {
		return found
	}
	if design == "" {
		return true
	}

	if maxLength >= len(design) {
		startLength = len(design)
	} else {
		startLength = maxLength
	}

	section := design[:startLength]
	towelFound := false
	for i := startLength; i > 0; i-- {
		if towelFound {
			break
		}
		for _, t := range typeMap[i] {
			if section[:i] == t {
				towelFound = FindPattern(design[i:], typeMap, maxLength, memoMap)
				if towelFound {
					break
				}
			}
		}

	}
	memoMap[design] = towelFound

	return towelFound
}

func FindPattern2(design string, typeMap map[int][]string, maxLength int, memoMap map[string]int) int {
	startLength := 0
	arrangements := 0
	if found, exists := memoMap[design]; exists {
		return found
	}
	if design == "" {
		return 1
	}

	if maxLength >= len(design) {
		startLength = len(design)
	} else {
		startLength = maxLength
	}

	section := design[:startLength]
	towelFound := false
	for i := startLength; i > 0; i-- {
		if towelFound {
			break
		}
		for _, t := range typeMap[i] {
			if section[:i] == t {
				arrangements = arrangements + FindPattern2(design[i:], typeMap, maxLength, memoMap)
			}
		}

	}
	memoMap[design] = arrangements

	return arrangements
}

func Part1() (int, error) {
	possibleDesigns := 0
	memoMap := map[string]bool{}
	types, patterns, err := ReadInput()
	typeMap := map[int][]string{}
	if err != nil {
		return 0, err
	}
	maxLength := 0
	for _, t := range types {
		typeMap[len(strings.TrimSpace(t))] = append(typeMap[len(strings.TrimSpace(t))], strings.TrimSpace(t))
		if len(strings.TrimSpace(t)) > maxLength {
			maxLength = len(strings.TrimSpace(t))
		}
	}

	towelFound := false
	for _, str := range patterns {
		towelFound = FindPattern(str, typeMap, maxLength, memoMap)
		if towelFound {
			possibleDesigns = possibleDesigns + 1
		}
	}

	return possibleDesigns, nil
}

func Part2() (int, error) {
	arrangements := 0

	memoMap := map[string]int{}
	types, patterns, err := ReadInput()
	typeMap := map[int][]string{}
	if err != nil {
		return 0, err
	}
	maxLength := 0
	for _, t := range types {
		typeMap[len(strings.TrimSpace(t))] = append(typeMap[len(strings.TrimSpace(t))], strings.TrimSpace(t))
		if len(strings.TrimSpace(t)) > maxLength {
			maxLength = len(strings.TrimSpace(t))
		}
	}

	for _, str := range patterns {
		arrangements = arrangements + FindPattern2(str, typeMap, maxLength, memoMap)
	}

	return arrangements, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Possible Designs: %v \n", result1)

	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in Part 2: %v", err)
	}
	fmt.Printf("Number of arrangments: %v \n", result2)
}
