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

func PopulateTree(root *Node, operands []int) {
	if len(operands) == 0 {
		return
	} else {
		leftNode := &Node{
			Value: root.Value + operands[0],
			Left:  nil,
			Right: nil,
		}
		rightNode := &Node{
			Value: root.Value * operands[0],
			Left:  nil,
			Right: nil,
		}
		root.Left = leftNode
		root.Right = rightNode

		PopulateTree(leftNode, operands[1:])
		PopulateTree(rightNode, operands[1:])

	}
}

func FindInLeaves(root *Node, element int) bool {
	if root.Left == nil && root.Right == nil {
		if root.Value == element {
			return true
		} else {
			return false
		}
	} else {
		return FindInLeaves(root.Left, element) || FindInLeaves(root.Right, element)
	}
}

func PopulateTree2(root *Node2, operands []int) {
	if len(operands) == 0 {
		return
	} else {
		sumNode := &Node2{
			Value:  root.Value + operands[0],
			Sum:    nil,
			Prod:   nil,
			Concat: nil,
		}
		prodNode := &Node2{
			Value:  root.Value * operands[0],
			Sum:    nil,
			Prod:   nil,
			Concat: nil,
		}
		concatInts, err := strconv.Atoi(strconv.Itoa(root.Value) + strconv.Itoa(operands[0]))
		if err != nil {
			fmt.Printf("Error in string to integer conversion ")
		}
		concatNode := &Node2{
			Value:  concatInts,
			Sum:    nil,
			Prod:   nil,
			Concat: nil,
		}

		root.Sum = sumNode
		root.Prod = prodNode
		root.Concat = concatNode

		PopulateTree2(sumNode, operands[1:])
		PopulateTree2(prodNode, operands[1:])
		PopulateTree2(concatNode, operands[1:])

	}
}

func FindInLeaves2(root *Node2, element int) bool {
	if root.Concat == nil && root.Sum == nil && root.Prod == nil {
		if root.Value == element {
			return true
		} else {
			return false
		}
	} else {
		return FindInLeaves2(root.Concat, element) || FindInLeaves2(root.Sum, element) || FindInLeaves2(root.Prod, element)
	}
}

func Part1() (int, error) {
	sum := 0
	equations, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, equation := range equations {
		root := &Node{
			Value: equation.Operands[0],
			Right: nil,
			Left:  nil,
		}
		PopulateTree(root, equation.Operands[1:])
		if FindInLeaves(root, equation.Result) {
			sum = sum + equation.Result
		}

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
		root := &Node2{
			Value:  equation.Operands[0],
			Sum:    nil,
			Prod:   nil,
			Concat: nil,
		}
		PopulateTree2(root, equation.Operands[1:])
		if FindInLeaves2(root, equation.Result) {
			sum = sum + equation.Result
		}

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
