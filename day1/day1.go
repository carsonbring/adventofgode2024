package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part1() {
	var group1 []int
	var group2 []int
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Couldn't read input file")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			fmt.Println(fmt.Errorf("Input in incorrect format"))
		} else {
			location1, err := strconv.Atoi(line[0])
			location2, err2 := strconv.Atoi(line[1])
			if err != nil || err2 != nil {
				fmt.Println(fmt.Errorf("Input in incorrect format, couldn't convert strings to ints"))
			} else {
				group1 = append(group1, location1)
				group2 = append(group2, location2)
			}
		}
	}

	totalDist := 0
	slices.Sort(group1)
	slices.Sort(group2)
	for i := range group1 {
		difference := absInt(group2[i] - group1[i])
		totalDist = totalDist + difference
	}
	fmt.Printf("The total distance is: %v \n", totalDist)
}

func part2() {
	var group1 []int
	scores := make(map[int]int)
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Couldn't read input file")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			fmt.Println(fmt.Errorf("Input in incorrect format"))
		} else {
			location1, err := strconv.Atoi(line[0])
			location2, err2 := strconv.Atoi(line[1])
			if err != nil || err2 != nil {
				fmt.Println(fmt.Errorf("Input in incorrect format, couldn't convert strings to ints"))
			} else {
				group1 = append(group1, location1)
				if _, exists := scores[location2]; exists {
					scores[location2] = scores[location2] + 1
				} else {
					scores[location2] = 1
				}
			}
		}
	}

	similarityScore := 0
	for _, id := range group1 {
		scoretoadd := id * scores[id]
		similarityScore = similarityScore + scoretoadd
	}
	fmt.Printf("The similarityscore is: %v \n", similarityScore)
}

func main() {
	part1()
	part2()
}
