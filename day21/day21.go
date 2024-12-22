package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Position struct {
	R, C int
}
type Path struct {
	start, end rune
}

type PathLevel struct {
	end   rune
	level int
}
type Node struct {
	adjList []Position
	val     rune
	visited bool
}

type Queue []*Item

func (q Queue) Len() int { return len(q) }

func (q *Queue) Enqueue(pos any) {
	*q = append(*q, pos.(*Item))
}

func (q *Queue) Dequeue() any {
	old := *q
	x := old[0]
	*q = old[1:]
	return x
}

type Item struct {
	pos        Position
	distance   int
	directions []rune
}

func extractInteger(s string) (int, error) {
	var digits string
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits += string(r)
		}
	}
	if digits == "" {
		return 0, fmt.Errorf("no digits found in string")
	}
	return strconv.Atoi(digits)
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func ReadInput() ([][]rune, error) {
	file, err := os.Open("test_input.txt")
	codes := [][]rune{}
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		currentcode := []rune{}
		for _, rune := range line {
			currentcode = append(currentcode, rune)
		}
		codes = append(codes, currentcode)
	}
	return codes, nil
}

func createNumPad() (map[Position]*Node, error) {
	graph := map[Position]*Node{}

	file, err := os.Open("numpad.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, rune := range line {
			if rune != 'X' {
				graph[Position{R: rowCount, C: col}] = &Node{
					adjList: []Position{},
					val:     rune,
				}
			}
		}
		rowCount++
	}

	for pos, node := range graph {
		for adjRow := -1; adjRow < 2; adjRow++ {
			for adjCol := -1; adjCol < 2; adjCol++ {
				if (adjRow == 0 && adjCol == 0) || (AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1) {
					continue
				}
				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if _, exists := graph[neighborPos]; exists {
					node.adjList = append(node.adjList, neighborPos)
				}
			}
		}
	}

	return graph, nil
}

func createArrowPad() (map[Position]*Node, error) {
	graph := map[Position]*Node{}

	file, err := os.Open("arrowpad.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for col, rune := range line {
			if rune != 'X' {
				graph[Position{R: rowCount, C: col}] = &Node{
					adjList: []Position{},
					val:     rune,
				}
			}
		}
		rowCount++
	}

	for pos, node := range graph {
		for adjRow := -1; adjRow < 2; adjRow++ {
			for adjCol := -1; adjCol < 2; adjCol++ {
				if (adjRow == 0 && adjCol == 0) || (AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1) {
					continue
				}
				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if _, exists := graph[neighborPos]; exists {
					node.adjList = append(node.adjList, neighborPos)
				}
			}
		}
	}

	return graph, nil
}

func findDirection(source Position, next Position) rune {
	rowDiff := next.R - source.R
	colDiff := next.C - source.C
	direction := '?'
	if rowDiff == -1 {
		direction = '^'
	} else if rowDiff == 1 {
		direction = 'v'
	} else if colDiff == -1 {
		direction = '<'
	} else if colDiff == 1 {
		direction = '>'
	}
	return direction
}

func BFS(sourcePosition Position, nodeMap map[Position]*Node, pathMap map[Path][]rune) {
	q := &Queue{&Item{pos: sourcePosition, distance: 0, directions: []rune{}}}
	for _, val := range nodeMap {
		val.visited = false
	}
	nodeMap[sourcePosition].visited = true

	for q.Len() != 0 {
		current := q.Dequeue().(*Item)

		for _, neighbor := range nodeMap[current.pos].adjList {
			if !nodeMap[neighbor].visited {

				nodeMap[neighbor].visited = true
				neighborPath := append([]rune{}, current.directions...)
				neighborPath = append(neighborPath, findDirection(current.pos, neighbor))
				pathMap[Path{start: nodeMap[sourcePosition].val, end: nodeMap[neighbor].val}] = append(neighborPath, 'A')
				q.Enqueue(&Item{directions: neighborPath, pos: neighbor})
			}
		}
	}
}

func getShortestPaths() (map[Path][]rune, map[Path][]rune, error) {
	numpad, err := createNumPad()
	if err != nil {
		return nil, nil, err
	}

	arrowpad, err := createArrowPad()
	if err != nil {
		return nil, nil, err
	}
	numPaths := map[Path][]rune{}
	arrowPaths := map[Path][]rune{}
	for pos := range numpad {
		BFS(pos, numpad, numPaths)
	}

	for pos := range arrowpad {
		BFS(pos, arrowpad, arrowPaths)
	}
	return numPaths, arrowPaths, nil
}

// func recResolver(start, end rune, numPaths, arrowPaths map[Path][]rune, memoMap map[PathLevel]int, level int) int {
// 	fmt.Printf("start: %v end: %v \n", string(start), string(end))
// 	sum := 0
// 	if val, exists := memoMap[PathLevel{level: level, start: start, end: end}]; exists {
// 		return val
// 	} else if level == 1 {
// 		sum = len(arrowPaths[Path{start: start, end: end}])
// 	} else if level == 3 {
// 		newPath := numPaths[Path{start: start, end: end}]
// 		for i := 0; i < len(newPath)-1; i++ {
// 			sum = sum + recResolver(newPath[i], newPath[i+1], numPaths, arrowPaths, memoMap, level-1)
// 		}
// 	} else {
// 		newPath := arrowPaths[Path{start: start, end: end}]
// 		fmt.Printf("newPath length: %v", len(newPath))
// 		for i := 0; i < len(newPath)-1; i++ {
// 			sum = sum + recResolver(newPath[i], newPath[i+1], numPaths, arrowPaths, memoMap, level-1)
// 		}
// 	}
// 	memoMap[PathLevel{level: level, start: start, end: end}] = sum
// 	return sum
// }

func recResolver(end rune, numPaths, arrowPaths map[Path][]rune, memoMap map[PathLevel]int, level int, start rune) int {
	// fmt.Printf("end: %v \n", string(end))
	sum := 0

	// if val, exists := memoMap[PathLevel{level: level, end: end}]; exists {
	// 	return val
	if level == 1 {

		if end == 'A' {
			fmt.Print("A")
			return 1
		}
		sum = len(arrowPaths[Path{start: start, end: end}])

		fmt.Printf("%v", string(arrowPaths[Path{start: start, end: end}]))
	} else if level == 3 {
		newPath := numPaths[Path{start: start, end: end}]

		for i := 0; i < len(newPath); i++ {
			sum = sum + recResolver(newPath[i], numPaths, arrowPaths, memoMap, level-1, start)
			start = newPath[i]
		}
	} else {

		newPath := arrowPaths[Path{start: start, end: end}]

		for i := 0; i < len(newPath); i++ {
			sum = sum + recResolver(newPath[i], numPaths, arrowPaths, memoMap, level-1, start)
			start = newPath[i]
		}
	}
	memoMap[PathLevel{level: level, end: end}] = sum
	return sum
}

func Part1() (int, error) {
	complexity := 0
	codes, err := ReadInput()
	if err != nil {
		return 0, err
	}
	numPaths, arrowPaths, err := getShortestPaths()
	memoMap := map[PathLevel]int{}
	if err != nil {
		return 0, err
	}
	start := 'A'
	for _, code := range codes {
		length := 0
		for i := 0; i < len(code); i++ {
			length = length + recResolver(code[i], numPaths, arrowPaths, memoMap, 3, start)
			start = code[i]
		}
		codeString := string(code)
		numericCode, err := extractInteger(codeString)
		if err != nil {
			return 0, err
		}
		complexity = complexity + (numericCode * length)
	}
	return complexity, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Sum of complexities: %v \n", result1)

	// result2, err := Part2(100)
	// if err != nil {
	// 	fmt.Printf("Error in Part 1: %v", err)
	// }
	// fmt.Printf("Number of cheats (20 pico):  %v\n", result2)
}
