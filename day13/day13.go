package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isInteger(num float64) bool {
	test := num - math.Round(num)
	return math.Abs(test) < .01
}

func TransposeMatrix(matrix [][]int) ([][]float64, bool) {
	transposed := [][]float64{{0, 0}, {0, 0}}
	a := float64(matrix[0][0])
	b := float64(matrix[0][1])
	c := float64(matrix[1][0])
	d := float64(matrix[1][1])
	if a*d-b*c == 0 {
		return nil, true
	}

	determinant := 1 / (a*d - b*c)
	if determinant == 0 {
		return nil, true
	}

	transposed[0][0] = d * determinant
	transposed[0][1] = -1 * (b * determinant)

	transposed[1][0] = -1 * (c * determinant)
	transposed[1][1] = a * determinant

	return transposed, false
}

func MultiplyMatrix(answers []int, buttons [][]float64) []float64 {
	buttonPresses := []float64{}
	var result float64 = 0
	for i := range answers {
		result = result + float64(answers[0])*buttons[0][i]
		result = result + float64(answers[1])*buttons[1][i]
		buttonPresses = append(buttonPresses, result)
		result = 0
	}
	return buttonPresses
}

func ReadInput(addend int) ([][][]int, [][]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		return nil, nil, err
	}

	buttonMatrices := [][][]int{}
	answerMatrices := [][]int{}

	buttonMatrix := [][]int{{0, 0}, {0, 0}}
	col := 0

	for scanner.Scan() {

		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)

		if strings.Contains(line, "Prize") {
			intMatch, err := strconv.Atoi(matches[0][0])
			intMatch2, err2 := strconv.Atoi(matches[1][0])
			if err != nil || err2 != nil {
				return nil, nil, err
			}
			buttonMatrices = append(buttonMatrices, buttonMatrix)
			answerMatrices = append(answerMatrices, []int{intMatch + addend, intMatch2 + addend})
			buttonMatrix = [][]int{{0, 0}, {0, 0}}

			col = 0

		} else if strings.Contains(line, "Button") {
			intMatch, err := strconv.Atoi(matches[0][0])
			intMatch2, err2 := strconv.Atoi(matches[1][0])
			if err != nil || err2 != nil {
				return nil, nil, err
			}
			buttonMatrix[col][0] = intMatch
			buttonMatrix[col][1] = intMatch2
			col++
		}

	}
	return buttonMatrices, answerMatrices, nil
}

func Part1() (int, error) {
	buttonMatrices, answerMatrices, err := ReadInput(0)
	if err != nil {
		return 0, err
	}
	totalTokens := 0
	for i, buttonMatrix := range buttonMatrices {
		transposed, none := TransposeMatrix(buttonMatrix)
		if none {
			continue
		}
		buttonPresses := MultiplyMatrix(answerMatrices[i], transposed)
		if buttonPresses[0] < 0 || buttonPresses[1] < 0 || !isInteger(buttonPresses[0]) || !isInteger(buttonPresses[1]) {
			continue
		} else {
			tokens := int(math.Round(buttonPresses[0])*3 + math.Round(buttonPresses[1])*1)
			totalTokens = totalTokens + tokens
		}

	}
	return totalTokens, nil
}

func Part2() (int, error) {
	buttonMatrices, answerMatrices, err := ReadInput(10000000000000)
	if err != nil {
		return 0, err
	}
	totalTokens := 0
	for i, buttonMatrix := range buttonMatrices {
		transposed, none := TransposeMatrix(buttonMatrix)
		if none {
			continue
		}
		buttonPresses := MultiplyMatrix(answerMatrices[i], transposed)
		if buttonPresses[0] < 0 || buttonPresses[1] < 0 || !isInteger(buttonPresses[0]) || !isInteger(buttonPresses[1]) {
			continue
		} else {
			tokens := int(math.Round(buttonPresses[0])*3 + math.Round(buttonPresses[1])*1)
			totalTokens = totalTokens + tokens
		}

	}
	return totalTokens, nil
}

func main() {
	answer1, err := Part1()
	if err != nil {
		fmt.Printf("Error occurred in part 1 %v \n \n", err)
	} else {
		fmt.Printf("Total Tokens (part 1) : %v \n\n", answer1)
	}
	answer2, err := Part2()
	if err != nil {
		fmt.Printf("Error occurred in part 2 %v \n \n", err)
	} else {
		fmt.Printf("Total Tokens (part 2) : %v \n\n", answer2)
	}
}
