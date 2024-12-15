package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	xpos, ypos int
}
type Robot struct {
	xpos, ypos int
	dX, dY     int
}

func ReadInput() ([]*Robot, error) {
	robots := []*Robot{}

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	re, err := regexp.Compile(`-?\d+`)
	if err != nil {
		return nil, err
	}
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		robotVals := []int{}
		for i := 0; i < 4; i++ {
			val, err := strconv.Atoi(matches[i][0])
			if err != nil {
				return nil, err
			}
			robotVals = append(robotVals, val)
		}
		robots = append(robots, &Robot{
			xpos: robotVals[0],
			ypos: robotVals[1],
			dX:   robotVals[2],
			dY:   robotVals[3],
		})
	}
	return robots, nil
}

func PrintMatrix(matrix [][]rune, number int) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", number)

	defer file.Close()
	for _, runes := range matrix {
		line := string(runes)
		fmt.Printf("%v\n", line)
	}
	return nil
}

func moveRobot(bot *Robot, seconds int, width int, height int) (int, int) {
	newX := (width + (bot.xpos + ((bot.dX * seconds) % width))) % width
	newY := (height + (bot.ypos + ((bot.dY * seconds) % height))) % height
	return newX, newY
}

func Part1(width int, height int) (int, error) {
	xBar := width / 2
	yBar := height / 2
	robots, err := ReadInput()
	if err != nil {
		return 0, err
	}
	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0

	for _, bot := range robots {
		x, y := moveRobot(bot, 100, width, height)
		if x < xBar && y > yBar {
			quad1 = quad1 + 1
		}
		if x > xBar && y > yBar {
			quad2 = quad2 + 1
		}
		if x < xBar && y < yBar {
			quad3 = quad3 + 1
		}
		if x > xBar && y < yBar {
			quad4 = quad4 + 1
		}
	}
	return quad1 * quad2 * quad3 * quad4, nil
}

func Part2(width int, height int) (int, error) {
	matrix := [][]rune{}
	for i := 0; i < height; i++ {
		row := []rune{}
		for j := 0; j < width; j++ {
			row = append(row, '.')
		}
		matrix = append(matrix, row)
	}

	robots, err := ReadInput()
	if err != nil {
		return 0, err
	}

	for i := 1; i < 10000; i++ {

		for _, bot := range robots {
			x, y := moveRobot(bot, i, width, height)
			matrix[y][x] = 'X'
		}

		for _, row := range matrix {
			line := string(row)
			if strings.Contains(line, "XXXXXXXXXXXX") {
				PrintMatrix(matrix, i)
				return i, nil
			}
		}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				matrix[i][j] = '.'
			}
		}

	}
	return 0, nil
}

func main() {
	answer1, err := Part1(101, 103)
	if err != nil {
		fmt.Printf("Error in Part1", err)
	}

	fmt.Printf("Safety Level (part1): %v \n", answer1)

	answer2, err := Part2(101, 103)
	if err != nil {
		fmt.Printf("Error in Part2", err)
	}

	fmt.Printf("Christmas tree easter egg second: %v \n", answer2)
}
