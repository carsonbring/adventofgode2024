package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (map[int]int, error) {
	stones := map[int]int{}
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		stonestr := strings.Fields(line)
		for _, str := range stonestr {
			currentInt, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			stones[currentInt] = 1

		}

	}
	return stones, nil
}

func BlinkRLE(stoneCounts map[int]int, blink int) {
	for i := 0; i < blink; i++ {
		changes := make(map[int]int)

		for key, value := range stoneCounts {
			BlinkModify(key, value, changes)
		}
		for key, value := range changes {
			stoneCounts[key] += value
		}
	}
}

func BlinkModify(stoneNum int, count int, changes map[int]int) {
	if count == 0 {
		return
	}
	if stoneNum == 0 {
		if prevCount, exists := changes[1]; exists {
			changes[1] = count + prevCount
		} else {
			changes[1] = count
		}
	} else if strNum := strconv.Itoa(stoneNum); len(strNum)%2 == 0 {
		middlePoint := len(strNum) / 2
		stone1, err := strconv.Atoi(strNum[:middlePoint])
		stone2, err2 := strconv.Atoi(strNum[middlePoint:])
		if err != nil && err2 != nil {
			fmt.Print("Error occurred in recursion")
		}
		if prevCount, exists := changes[stone1]; exists {
			changes[stone1] = count + prevCount
		} else {
			changes[stone1] = count
		}

		if prevCount, exists := changes[stone2]; exists {
			changes[stone2] = count + prevCount
		} else {
			changes[stone2] = count
		}

	} else {
		yearVal := stoneNum * 2024

		if prevCount, exists := changes[yearVal]; exists {
			changes[yearVal] = count + prevCount
		} else {
			changes[yearVal] = count
		}

	}
	changes[stoneNum] = changes[stoneNum] - count
}

func Part1() (int, error) {
	totalStones := 0
	counts, err := ReadInput()
	if err != nil {
		return 0, err
	}
	BlinkRLE(counts, 25)
	for _, val := range counts {
		totalStones = totalStones + val
	}
	return totalStones, nil
}

func Part2() (int, error) {
	totalStones := 0
	counts, err := ReadInput()
	if err != nil {
		return 0, err
	}
	BlinkRLE(counts, 75)
	for _, val := range counts {
		totalStones = totalStones + val
	}
	return totalStones, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part1: %v", err)
	}
	fmt.Printf("Stones afterblinking 25 times: %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part2: %v", err)
	}
	fmt.Printf("Stones after blinking 75 times: %v \n", result2)
}
