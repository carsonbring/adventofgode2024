package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Right *Node
	Left  *Node
}
type Node2 struct {
	Value  int
	Sum    *Node2
	Prod   *Node2
	Concat *Node2
}
type Equation struct {
	Result   int
	Operands []int
}

func ConvertStringsToInts(strings []string) ([]int, error) {
	integers := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q to int: %v", s, err)
		}
		integers[i] = n
	}
	return integers, nil
}

func ReadInput() ([]Equation, error) {
	var equations []Equation
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentEquation := Equation{
			Result:   -1,
			Operands: make([]int, 0),
		}
		line := scanner.Text()
		if colonIndex := strings.Index(line, ":"); colonIndex != -1 {
			currentEquation.Result, err = strconv.Atoi(line[:colonIndex])
			if err != nil {
				return nil, err
			}
			currentEquation.Operands, err = ConvertStringsToInts(strings.Fields(line[colonIndex+1:]))
			if err != nil {
				return nil, err
			}
			equations = append(equations, currentEquation)

		}
	}

	return equations, nil
}

func RecursiveSearch(value int, result int, operands []int) int {
	if len(operands) == 0 {
		if result == value {
			return value
		} else {
			return 0
		}
	} else {
		if RecursiveSearch(value*operands[0], result, operands[1:]) != 0 || RecursiveSearch(value+operands[0], result, operands[1:]) != 0 {
			return result
		} else {
			return 0
		}
	}
}

func RecursiveSearch2(value int, result int, operands []int) int {
	if len(operands) == 0 {
		if result == value {
			return value
		} else {
			return 0
		}
	} else {
		concatInts, err := strconv.Atoi(strconv.Itoa(value) + strconv.Itoa(operands[0]))
		if err != nil {
			fmt.Printf("Error in string to integer conversion ")
		}

		if RecursiveSearch2(value*operands[0], result, operands[1:]) != 0 || RecursiveSearch2(value+operands[0], result, operands[1:]) != 0 || RecursiveSearch2(concatInts, result, operands[1:]) != 0 {
			return result
		} else {
			return 0
		}
	}
}

func Part1() (int, error) {
	sum := 0
	equations, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, equation := range equations {
		sum = sum + RecursiveSearch(equation.Operands[0], equation.Result, equation.Operands[1:])
	}
	return sum, nil
}

func Part2() (int, error) {
	sum := 0
	equations, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, equation := range equations {
		sum = sum + RecursiveSearch2(equation.Operands[0], equation.Result, equation.Operands[1:])
	}
	return sum, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error occured in part 1: %v \n", err)
	}
	fmt.Printf("Total calibration result: %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error occured in part 2: %v \n", err)
	}
	fmt.Printf("Total calibration result (with concat): %v \n", result2)
}
