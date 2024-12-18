package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operand struct {
	combo   *int
	literal int
}
type Operator func(Operand, *int) (int, bool)

func InitOps() (map[int]Operator, map[int]Operand, map[rune]*int) {
	registers := map[rune]*int{}
	operandMap := map[int]Operand{}
	operatorMap := map[int]Operator{}
	regA := 0
	regB := 0
	regC := 0
	registers['A'] = &regA
	registers['B'] = &regB
	registers['C'] = &regC

	for i := 0; i < 4; i++ {
		v := i
		operandMap[i] = Operand{literal: i, combo: new(int)}
		*operandMap[i].combo = v
	}
	operandMap[4] = Operand{literal: 4, combo: registers['A']}
	operandMap[5] = Operand{literal: 5, combo: registers['B']}
	operandMap[6] = Operand{literal: 6, combo: registers['C']}
	operatorMap[0] = func(val Operand, iPointer *int) (int, bool) {
		exponent := *val.combo
		*registers['A'] = *registers['A'] >> exponent
		return 0, false
	}

	operatorMap[1] = func(val Operand, iPointer *int) (int, bool) {
		xor := *registers['B'] ^ val.literal
		*registers['B'] = xor
		return 0, false
	}

	operatorMap[2] = func(val Operand, iPointer *int) (int, bool) {
		lowest3bits := *val.combo & 7
		*registers['B'] = lowest3bits
		return 0, false
	}

	operatorMap[3] = func(val Operand, iPointer *int) (int, bool) {
		if *registers['A'] == 0 {
			return 0, false
		} else {
			(*iPointer) = val.literal
		}
		return 0, false
	}

	operatorMap[4] = func(val Operand, iPointer *int) (int, bool) {
		xor := *registers['B'] ^ *registers['C']
		*registers['B'] = xor
		return -1, false
	}

	operatorMap[5] = func(val Operand, iPointer *int) (int, bool) {
		lowest3bits := *val.combo & 7
		return lowest3bits, true
	}

	operatorMap[6] = func(val Operand, iPointer *int) (int, bool) {
		exponent := *val.combo
		*registers['B'] = *registers['A'] >> exponent
		return 0, false
	}

	operatorMap[7] = func(val Operand, iPointer *int) (int, bool) {
		exponent := *val.combo
		*registers['C'] = *registers['A'] >> exponent
		return 0, false
	}
	return operatorMap, operandMap, registers
}

func ReadInput(registers map[rune]*int) ([]int, error) {
	program := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if rowCount != 3 {
			line = strings.SplitN(line, ":", 2)[1]
			line = strings.TrimSpace(line)
		}
		if rowCount == 0 {
			*registers['A'], err = strconv.Atoi(line)
		} else if rowCount == 1 {
			*registers['B'], err = strconv.Atoi(line)
		} else if rowCount == 2 {
			*registers['C'], err = strconv.Atoi(line)
		} else if rowCount == 3 {
			rowCount++
			continue
		} else {
			strNumbers := strings.Split(line, ",")
			for _, str := range strNumbers {
				intOp := -1
				intOp, err = strconv.Atoi(str)
				program = append(program, intOp)
			}
		}
		if err != nil {
			return nil, fmt.Errorf("atoi conversion error")
		}
		rowCount++
	}
	return program, nil
}

func reverseSliceNew(s []int) []int {
	n := len(s)
	reversed := make([]int, n)
	for i, v := range s {
		reversed[n-1-i] = v
	}
	return reversed
}

func Part1() (string, error) {
	opcodeMap, operandMap, registers := InitOps()
	program, err := ReadInput(registers)
	if err != nil {
		return "", err
	}
	output := ""
	for i := 0; i < len(program)-1; i = i + 2 {
		instruction := i
		funOutput, returning := opcodeMap[program[i]](operandMap[program[i+1]], &i)
		if returning {
			output = output + strconv.Itoa(funOutput) + ","
		}
		if i != instruction {
			i = i - 2
		}

	}
	return output[:len(output)-1], nil
}

func recursiveSearch(A int, instructions []int, opcodeMap map[int]Operator, operandMap map[int]Operand, registers map[rune]*int, reverse []int, possibleAs *[]int) {
	if len(reverse) == 0 {
		*possibleAs = append(*possibleAs, A)
		return
	}
	for out := 0; out < 8; out++ {
		*registers['A'] = (A << 3) | out
		*registers['B'] = 0
		*registers['C'] = 0
		output := ""

		for i := 0; i < len(instructions)-1; i = i + 2 {
			instruction := i

			funOutput, returning := opcodeMap[instructions[i]](operandMap[instructions[i+1]], &i)
			if returning {
				output = output + strconv.Itoa(funOutput) + ","
			}
			if i != instruction {
				i = i - 2
			}
		}
		intSlice := []int{}
		strSlice := strings.Split(output[:len(output)-1], ",")
		for _, str := range strSlice {
			intOp := -1
			intOp, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("error in recursion %v", err)
			}
			intSlice = append(intSlice, intOp)
		}

		if intSlice[0] == reverse[0] {
			recursiveSearch((A<<3 | out), instructions, opcodeMap, operandMap, registers, reverse[1:], possibleAs)
		}

	}
}

func Part2() (int, error) {
	opcodeMap, operandMap, registers := InitOps()
	program, err := ReadInput(registers)
	reverse := reverseSliceNew(program)
	if err != nil {
		return -1, err
	}
	possibleAs := []int{}
	recursiveSearch(0, program, opcodeMap, operandMap, registers, reverse, &possibleAs)
	lowestA := 100000000000000000
	for _, a := range possibleAs {
		if a < lowestA {
			lowestA = a
		}
	}

	return lowestA, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Program output: %v \n", result1)

	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in Part 2: %v", err)
	}
	fmt.Printf("Uncorrupted A: %v \n", result2)
}
