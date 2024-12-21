package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	R, C int
}
type CheatPair struct {
	start Position
	end   Position
}
type Node struct {
	adjList     []Position
	cheatAdj    []Position
	wall        bool
	visited     bool
	finish      bool
	start       bool
	manDistance int
}

type BFSKey struct {
	pos          Position
	cheated      bool
	cheatedStart Position
	cheatedEnd   Position
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
	pos          Position
	distance     int
	cheated      bool
	cheatedStart Position
	cheatedEnd   Position
}

func convertRunesToNodes(runes []rune, rowCount int, nodeMap map[Position]*Node, startingPosition *Position, endPosition *Position) []Node {
	nodes := []Node{}
	for i, rune := range runes {
		newPosition := Position{R: rowCount, C: i}
		newNode := &Node{}

		switch rune {
		case '#':
			newNode = &Node{
				adjList: []Position{},
				wall:    true,
				visited: false,
				finish:  false,
			}
		case '.':
			newNode = &Node{
				adjList: []Position{},
				wall:    false,
				visited: false,
				finish:  false,
			}

		case 'E':
			newNode = &Node{
				adjList: []Position{},
				wall:    false,
				visited: false,
				finish:  true,
			}
			*endPosition = newPosition

		case 'S':
			newNode = &Node{
				adjList: []Position{},
				wall:    false,
				visited: false,
				finish:  false,
				start:   true,
			}
			*startingPosition = newPosition
		}

		nodeMap[newPosition] = newNode
	}
	return nodes
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func ReadInput() (map[Position]*Node, Position, Position, error) {
	nodeMap := map[Position]*Node{}
	start := Position{}
	end := Position{}
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, Position{}, Position{}, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		convertRunesToNodes(runes, rowCount, nodeMap, &start, &end)
		rowCount++
	}

	for pos, node := range nodeMap {
		for adjRow := -2; adjRow < 3; adjRow++ {
			for adjCol := -2; adjCol < 3; adjCol++ {
				if adjRow == 0 && adjCol == 0 || AbsInt(adjRow) == 2 && AbsInt(adjCol) == 2 || AbsInt(adjRow) == 2 && AbsInt(adjCol) == 1 || AbsInt(adjRow) == 1 && AbsInt(adjCol) == 2 {
					continue
				}

				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1 {
					if _, exists := nodeMap[neighborPos]; exists {
						if !nodeMap[neighborPos].wall {
							node.cheatAdj = append(node.cheatAdj, neighborPos)
						}
					}
				} else if AbsInt(adjCol) == 2 || AbsInt(adjRow) == 2 {
					if _, exists := nodeMap[neighborPos]; exists {
						if !nodeMap[neighborPos].wall {
							node.cheatAdj = append(node.cheatAdj, neighborPos)
						}
					}
				} else {
					if _, exists := nodeMap[neighborPos]; exists {
						if !nodeMap[neighborPos].wall {
							node.adjList = append(node.adjList, neighborPos)
						}
					}
				}
			}
		}
	}
	return nodeMap, start, end, nil
}

func ReadInput2() (map[Position]*Node, Position, Position, error) {
	nodeMap := map[Position]*Node{}
	start := Position{}
	end := Position{}
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, Position{}, Position{}, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		convertRunesToNodes(runes, rowCount, nodeMap, &start, &end)
		rowCount++
	}

	for pos, node := range nodeMap {
		for adjRow := -20; adjRow < 21; adjRow++ {
			for adjCol := -20; adjCol < 21; adjCol++ {

				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if adjRow == 0 && AbsInt(adjCol) == 1 || AbsInt(adjRow) == 1 && adjCol == 0 {
					if _, exists := nodeMap[neighborPos]; exists {
						if !nodeMap[neighborPos].wall {
							node.adjList = append(node.adjList, neighborPos)
						}
					}
				} else if AbsInt(adjRow)+AbsInt(adjCol) <= 20 {
					if _, exists := nodeMap[neighborPos]; exists {
						if !nodeMap[neighborPos].wall {
							node.cheatAdj = append(node.cheatAdj, neighborPos)
						}
					}
				}

			}
		}
	}
	return nodeMap, start, end, nil
}

func initBFS(sourcePosition Position, nodeMap map[Position]*Node) int {
	q := &Queue{&Item{pos: sourcePosition, distance: 0}}
	for _, val := range nodeMap {
		val.visited = false
	}

	for q.Len() != 0 {
		current := q.Dequeue()

		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			t_distance := current.(*Item).distance + 1
			if !nodeMap[neighbor].visited {
				nodeMap[neighbor].visited = true
				if nodeMap[neighbor].finish {
					return t_distance
				} else {
					q.Enqueue(&Item{distance: t_distance, pos: neighbor})
				}
			}
		}
	}
	return -1
}

func initManDist(endPosition Position, nodeMap map[Position]*Node) int {
	q := &Queue{&Item{pos: endPosition, distance: 0}}
	for _, val := range nodeMap {
		val.visited = false
	}

	nodeMap[endPosition].visited = true
	nodeMap[endPosition].manDistance = 0

	for q.Len() != 0 {
		current := q.Dequeue()

		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			t_distance := current.(*Item).distance + 1
			if !nodeMap[neighbor].visited {
				nodeMap[neighbor].visited = true
				nodeMap[neighbor].manDistance = t_distance
				q.Enqueue(&Item{pos: neighbor, distance: t_distance})
			}
		}
	}
	return -1
}

func BFS(sourcePosition Position, nodeMap map[Position]*Node, cheatMap map[CheatPair]bool, timeSave int, originalTime int) {
	for _, val := range nodeMap {
		val.visited = false
	}
	visited := make(map[BFSKey]bool)

	q := &Queue{&Item{pos: sourcePosition, distance: 0, cheated: false}}

	for q.Len() != 0 {
		current := q.Dequeue()

		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			newKey := BFSKey{}

			if current.(*Item).cheated {
				newKey = BFSKey{
					pos:          neighbor,
					cheated:      true,
					cheatedStart: current.(*Item).cheatedStart,
					cheatedEnd:   current.(*Item).cheatedEnd,
				}
			} else {
				newKey = BFSKey{
					pos:          neighbor,
					cheated:      false,
					cheatedStart: Position{},
					cheatedEnd:   Position{},
				}
			}

			if !visited[newKey] {
				visited[newKey] = true
				t_distance := current.(*Item).distance + 1
				if t_distance <= originalTime-timeSave {
					if nodeMap[neighbor].finish {
						if current.(*Item).cheated {
							cheatMap[CheatPair{start: current.(*Item).cheatedStart, end: current.(*Item).cheatedEnd}] = true
						}
					} else {
						if current.(*Item).cheated {
							q.Enqueue(&Item{distance: t_distance, pos: neighbor, cheated: current.(*Item).cheated, cheatedStart: current.(*Item).cheatedStart, cheatedEnd: current.(*Item).cheatedEnd})
						} else {
							q.Enqueue(&Item{distance: t_distance, pos: neighbor})
						}
					}
				}
			}
		}
		if !current.(*Item).cheated {
			for _, neighbor := range nodeMap[current.(*Item).pos].cheatAdj {

				newKey := BFSKey{
					pos:          neighbor,
					cheated:      true,
					cheatedStart: current.(*Item).pos,
					cheatedEnd:   neighbor,
				}
				if !visited[newKey] {
					visited[newKey] = true
					cheatCost := AbsInt(current.(*Item).pos.R-neighbor.R) + AbsInt(current.(*Item).pos.C-neighbor.C)

					t_distance := current.(*Item).distance + cheatCost
					if nodeMap[neighbor].manDistance+t_distance <= originalTime-timeSave {
						cheatMap[CheatPair{start: current.(*Item).pos, end: neighbor}] = true
					}
				}
			}
		}
	}
}

func Part1(timeSave int) (int, error) {
	cheats := 0
	nodeMap, start, end, err := ReadInput()
	if err != nil {
		return 0, err
	}
	noCheatTime := initBFS(start, nodeMap)
	initManDist(end, nodeMap)
	cheatMap := map[CheatPair]bool{}
	BFS(start, nodeMap, cheatMap, timeSave, noCheatTime)
	for _, val := range cheatMap {
		if val {
			cheats = cheats + 1
		}
	}
	return cheats, nil
}

func Part2(timeSave int) (int, error) {
	cheats := 0
	nodeMap, start, end, err := ReadInput2()
	if err != nil {
		return 0, err
	}
	noCheatTime := initBFS(start, nodeMap)
	initManDist(end, nodeMap)
	cheatMap := map[CheatPair]bool{}
	BFS(start, nodeMap, cheatMap, timeSave, noCheatTime)
	for _, val := range cheatMap {
		if val {
			cheats = cheats + 1
		}
	}
	return cheats, nil
}

func main() {
	result1, err := Part1(100)
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Number of cheats (2 pico): %v \n", result1)

	result2, err := Part2(100)
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Number of cheats (20 pico):  %v\n", result2)
}
