package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	R, C int
}

func IsValidDir(number int) bool {
	if number >= -1 && number <= 1 {
		return true
	}
	return false
}

func ReadInput() ([][]rune, error) {
	var ceres [][]rune
	file, err := os.Open("input.txt")
	if err != nil {
		return ceres, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		ceres = append(ceres, runes)
	}

	return ceres, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func IncrementChecker(incrR, incrC int, pos Point, ceres [][]rune) int {
	word := "XMAS"

	for i, letter := range word {
		testR := pos.R + i*incrR
		testC := pos.C + i*incrC
		if (testR < 0 || testR >= len(ceres)) || (testC < 0 || testC >= len(ceres[testR])) {
			return 0
		} else if ceres[testR][testC] != letter {
			return 0
		}
	}
	return 1
}

func XChecker(pos Point, ceres [][]rune) int {
	word := "MAS"

	for i := -1; i < 3; i = i + 2 {
		tempR := pos.R + i
		var tempC int
		if i == 1 {
			tempC = pos.C - i
		} else {
			tempC = pos.C + i
		}
		if (tempR < 0 || tempR >= len(ceres)) || (tempC < 0 || tempC >= len(ceres[tempR])) {
			return 0
		} else if ceres[tempR][tempC] == rune(word[len(word)-1]) {
			word = reverseString(word)
		} else if ceres[tempR][tempC] != rune(word[0]) {
			return 0
		}
		incrR := -i
		var incrC int
		if i == 1 {
			incrC = i
		} else {
			incrC = -i
		}

		for j, letter := range word {

			testR := tempR + j*incrR
			testC := tempC + j*incrC

			if (testR < 0 || testR >= len(ceres)) || (testC < 0 || testC >= len(ceres[testR])) {
				return 0
			} else if ceres[testR][testC] != letter {
				return 0
			}
		}
	}
	return 1
}

func Part1() (int, error) {
	sum := 0
	ceres, err := ReadInput()
	if err != nil {
		fmt.Printf("Error in ReadInput: %v", err)
	}

	for i := 0; i < len(ceres); i++ {
		for j := 0; j < len(ceres[i]); j++ {
			currentPoint := Point{
				R: i,
				C: j,
			}
			for row := -1; row < 2; row++ {
				for col := -1; col < 2; col++ {
					if row == 0 && col == 0 {
						continue
					} else {
						sum = sum + IncrementChecker(row, col, currentPoint, ceres)
					}
				}
			}
		}
	}
	return sum, nil
}

func Part2() (int, error) {
	sum := 0
	ceres, err := ReadInput()
	if err != nil {
		fmt.Printf("Error in ReadInput: %v", err)
	}

	for i := 0; i < len(ceres); i++ {
		for j := 0; j < len(ceres[i]); j++ {
			currentPoint := Point{
				R: i,
				C: j,
			}
			sum = sum + XChecker(currentPoint, ceres)
		}
	}
	return sum, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part1: %v", err)
	}
	fmt.Printf("Number of XMAS: %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part2: %v \n", err)
	}
	fmt.Printf("Number of X's of MAS: %v \n", result2)
}
