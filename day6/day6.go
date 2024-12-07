package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	R, C int
}

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func PrintLayout(layout [][]rune) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, runes := range layout {
		line := string(runes)

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadInput() ([][]rune, Point, error) {
	var layout [][]rune
	guardIndex := Point{
		R: 0,
		C: 0,
	}
	rowCount := 0
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, guardIndex, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		if colIndex := slices.Index(runes, '^'); colIndex != -1 {
			guardIndex = Point{
				R: rowCount,
				C: colIndex,
			}
		}
		layout = append(layout, runes)
		rowCount++
	}

	return layout, guardIndex, nil
}

func CheckFinish(layout [][]rune, currentPos Point, dir int) bool {
	if dir == UP && currentPos.R == 0 {
		return true
	} else if dir == RIGHT && currentPos.C == len(layout[currentPos.R])-1 {
		return true
	} else if dir == DOWN && currentPos.R == len(layout)-1 {
		return true
	} else if dir == LEFT && currentPos.C == 0 {
		return true
	} else {
		return false
	}
}

func NextPoint(layout [][]rune, currentPos Point, dir int) Point {
	next := currentPos

	if dir == UP {
		next.R = next.R - 1
	} else if dir == RIGHT {
		next.C = next.C + 1
	} else if dir == DOWN {
		next.R = next.R + 1
	} else {
		next.C = next.C - 1
	}

	return next
}

func MoveForward(layout [][]rune, currentPos Point, dir int) (Point, bool, bool, int) {
	if CheckFinish(layout, currentPos, dir) {
		return currentPos, false, true, 0
	} else if nextPoint := NextPoint(layout, currentPos, dir); layout[nextPoint.R][nextPoint.C] == '#' {
		return currentPos, true, false, 0
	} else {
		if layout[currentPos.R][currentPos.C] == '.' || layout[currentPos.R][currentPos.C] == '^' {
			layout[currentPos.R][currentPos.C] = 'X'
			return nextPoint, false, false, 1
		} else {
			return nextPoint, false, false, 0
		}
	}
}

func Part1() (int, error) {
	layout, guardPoint, err := ReadInput()
	finish := false
	turn := false
	distinctTiles := 0
	direction := UP
	intTile := 0

	if err != nil {
		return 0, err
	}
	for {
		for i := direction; i < 4; i++ {
			for {
				guardPoint, turn, finish, intTile = MoveForward(layout, guardPoint, i)
				distinctTiles = distinctTiles + intTile
				if finish || turn {
					break
				}
			}
			if finish {
				break
			}
		}
		if finish {
			break
		}
	}
	file, err := os.Create("output.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	for _, runes := range layout {
		line := string(runes)

		_, err := file.WriteString(line + "\n")
		if err != nil {
			return 0, err
		}
	}

	return distinctTiles + 1, nil
}

func Part2() (int, error) {
	layout, guardPoint, err := ReadInput()
	numLoops := 0
	if err != nil {
		return 0, err
	}
	for rowIndex, row := range layout {
		for colIndex := range row {
			if layout[rowIndex][colIndex] == '.' {
				PrintLayout(layout)
				tempLayout := make([][]rune, len(layout))
				for i := range layout {
					tempLayout[i] = make([]rune, len(layout[i]))
					copy(tempLayout[i], layout[i])
				}

				tempLayout[rowIndex][colIndex] = '#'

				finish := false
				turn := false
				distinctTiles := 0
				direction := UP
				intTile := 0
				repeatCounter := 0
				loop := false
				tempGuardPoint := Point{
					R: guardPoint.R,
					C: guardPoint.C,
				}

				for {
					for i := direction; i < 4; i++ {
						for {
							tempGuardPoint, turn, finish, intTile = MoveForward(tempLayout, tempGuardPoint, i)
							if intTile == 0 {
								repeatCounter++
							} else {
								repeatCounter = 0
							}
							distinctTiles = distinctTiles + intTile
							if distinctTiles < repeatCounter {
								loop = true
							}
							if finish || turn || loop {
								break
							}
						}
						if finish || loop {
							break
						}
					}
					if finish || loop {
						break
					}
				}
				if loop {
					numLoops++
				}

			} else {
				continue
			}
		}
	}
	return numLoops - 1, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part1: %v", err)
	}
	fmt.Printf("Distance tiles (part 1): %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part2: %v", err)
	}
	fmt.Printf("Number of possible loops (part 2): %v \n", result2)
}
