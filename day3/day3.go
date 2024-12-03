package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() (int, error) {
	sum := 0
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		return 0, err
	}
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {
		prod := 1
		for i := 1; i < 3; i++ {
			matchInt, err := strconv.Atoi(match[i])
			if err != nil {
				return 0, err
			}
			prod = prod * matchInt
		}
		sum = sum + prod
	}
	return sum, nil
}

func Part2() (int, error) {
	sum := 0
	enabled := true
	data, err := os.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	re, err := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	if err != nil {
		return 0, err
	}
	matches := re.FindAllStringSubmatch(string(data), -1)

	for _, match := range matches {
		if strings.Contains(match[0], "mul") && enabled {
			prod := 1
			for i := 1; i < 3; i++ {
				matchInt, err := strconv.Atoi(match[i])
				if err != nil {
					return 0, err
				}
				prod = prod * matchInt
			}
			sum = sum + prod
		} else if strings.Contains(match[0], "don't") {
			enabled = false
		} else if strings.Contains(match[0], "do") {
			enabled = true
		}
	}
	return sum, nil
}

func main() {
	answer1, err := Part1()
	if err != nil {
		fmt.Printf("Error occurred in part 1 %v \n \n", err)
	} else {
		fmt.Printf("sum of products (part 1) : %v \n\n", answer1)
	}
	answer2, err := Part2()
	if err != nil {
		fmt.Printf("Error occurred in part 2 %v \n \n", err)
	} else {
		fmt.Printf("sum of enabled products (part 2) : %v \n\n", answer2)
	}
}
