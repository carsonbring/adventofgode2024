package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	R, C int
}

type SignalPair struct {
	pointA, pointB *Point
	dC, dR         int
}

func ReadInput() ([][]rune, map[rune][]Point, error) {
	var runes [][]rune
	antennaMap := map[rune][]Point{}
	rowCount := 0
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tempRunes := []rune(line)
		runes = append(runes, tempRunes)
		for i, point := range tempRunes {
			if point != '.' {
				newPoint := Point{
					R: rowCount,
					C: i,
				}
				if pointList, exists := antennaMap[point]; exists {
					antennaMap[point] = append(pointList, newPoint)
				} else {
					antennaMap[point] = []Point{newPoint}
				}
			}
		}
		rowCount++
	}
	return runes, antennaMap, nil
}

func signalAntinodes(layout [][]rune, signal SignalPair) int {
	sum := 0
	antiNodeR := signal.pointA.R + signal.dR
	antiNodeC := signal.pointA.C + signal.dC
	if antiNodeR >= 0 && antiNodeR < len(layout) && antiNodeC >= 0 && antiNodeC < len(layout[signal.pointA.R]) && layout[antiNodeR][antiNodeC] != '#' {
		layout[antiNodeR][antiNodeC] = '#'
		sum++
	}
	antiNodeR = signal.pointB.R - signal.dR
	antiNodeC = signal.pointB.C - signal.dC
	if antiNodeR >= 0 && antiNodeR < len(layout) && antiNodeC >= 0 && antiNodeC < len(layout[signal.pointB.R]) && layout[antiNodeR][antiNodeC] != '#' {
		layout[antiNodeR][antiNodeC] = '#'
		sum++
	}
	return sum
}

func ResonanceApplier(layout [][]rune, antiNodeR int, antiNodeC int, dR int, dC int) int {
	sum := 0
	for {
		extendCondition := antiNodeR >= 0 && antiNodeR < len(layout) && antiNodeC >= 0 && antiNodeC < len(layout[0])
		if !extendCondition {
			break
		}
		if layout[antiNodeR][antiNodeC] != '#' {
			layout[antiNodeR][antiNodeC] = '#'
			sum++
		}
		antiNodeR = antiNodeR + dR
		antiNodeC = antiNodeC + dC

	}
	return sum
}

func signalAntinodesResonance(layout [][]rune, signal SignalPair) int {
	sum := 0
	if layout[signal.pointA.R][signal.pointA.C] != '#' {
		layout[signal.pointA.R][signal.pointA.C] = '#'
		sum++
	}
	if layout[signal.pointB.R][signal.pointB.C] != '#' {
		layout[signal.pointB.R][signal.pointB.C] = '#'

		sum++
	}
	antiNodeR := signal.pointA.R + signal.dR
	antiNodeC := signal.pointA.C + signal.dC
	sum = sum + ResonanceApplier(layout, antiNodeR, antiNodeC, signal.dR, signal.dC)
	antiNodeR = signal.pointB.R - signal.dR
	antiNodeC = signal.pointB.C - signal.dC
	sum = sum + ResonanceApplier(layout, antiNodeR, antiNodeC, -signal.dR, -signal.dC)

	return sum
}

func findAntinodes(layout [][]rune, points []Point, resonance bool) int {
	if len(points) == 0 {
		return 0
	} else {
		tempSum := 0
		for _, point := range points[1:] {

			tempSignalPair := SignalPair{
				pointA: &points[0],
				pointB: &point,
				dC:     points[0].C - point.C,
				dR:     points[0].R - point.R,
			}

			if resonance {
				tempSum = tempSum + signalAntinodesResonance(layout, tempSignalPair)
			} else {
				tempSum = tempSum + signalAntinodes(layout, tempSignalPair)
			}

		}
		return tempSum + findAntinodes(layout, points[1:], resonance)

	}
}

func Part1() (int, error) {
	uniqueAntinodes := 0
	layout, antennaMap, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, points := range antennaMap {
		uniqueAntinodes = uniqueAntinodes + findAntinodes(layout, points, false)
	}

	return uniqueAntinodes, nil
}

func Part2() (int, error) {
	uniqueAntinodes := 0
	layout, antennaMap, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, points := range antennaMap {
		uniqueAntinodes = uniqueAntinodes + findAntinodes(layout, points, true)
	}

	return uniqueAntinodes, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part1: %v \n", err)
	}
	fmt.Printf("Unique Antinodes (part 1): %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part2: %v \n", err)
	}
	fmt.Printf("Unique Antinodes accounting for resonance (part 2): %v \n", result2)
}
