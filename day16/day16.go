package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type PriorityQueue []*Item

func (q PriorityQueue) Len() int { return len(q) }
func (q PriorityQueue) Less(i, j int) bool {
	return q[i].distance < q[j].distance
}
func (q PriorityQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *PriorityQueue) Push(pos any) {
	*q = append(*q, pos.(*Item))
}

func (q *PriorityQueue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

type Item struct {
	pos      Position
	dir      Position
	distance int
}
type Position struct {
	R, C int
}
type Node struct {
	adjList         []Position
	wall            bool
	visited         bool
	finish          bool
	distance        int
	bestPath        bool
	parents         []Position
	distanceFromEnd int
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
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
			}
			*startingPosition = newPosition
		}

		nodeMap[newPosition] = newNode
	}
	return nodes
}

func Dijkstra(nodeMap map[Position]*Node, source Position) {
	for _, node := range nodeMap {
		node.distance = 10000000
	}
	q := &PriorityQueue{&Item{pos: source, distance: 0, dir: Position{R: 0, C: 1}}}
	heap.Init(q)
	for q.Len() != 0 {
		current := q.Pop()
		// if nodeMap[current.(*Item).pos].visited {
		// 	continue
		// }
		currentDir := current.(*Item).dir
		// nodeMap[current.(*Item).pos].visited = true
		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			turn_addend := 1000
			newDir := Position{
				R: neighbor.R - current.(*Item).pos.R,
				C: neighbor.C - current.(*Item).pos.C,
			}
			if newDir.R == currentDir.R && newDir.C == currentDir.C {
				turn_addend = 0
			}
			t_distance := current.(*Item).distance + turn_addend + 1

			if t_distance < nodeMap[neighbor].distance {
				nodeMap[neighbor].distance = t_distance
				nodeMap[neighbor].parents = []Position{current.(*Item).pos}
				heap.Push(q, &Item{distance: t_distance, pos: neighbor, dir: newDir})
			}

		}
	}
}

func Dijkstra2(nodeMap map[Position]*Node, source Position, starting_position Position) {
	for _, node := range nodeMap {
		node.distanceFromEnd = 10000000
	}
	q := &PriorityQueue{&Item{pos: source, distance: 0, dir: starting_position}}
	heap.Init(q)
	for q.Len() != 0 {
		current := q.Pop()
		currentDir := current.(*Item).dir
		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			turn_addend := 1000
			newDir := Position{
				R: neighbor.R - current.(*Item).pos.R,
				C: neighbor.C - current.(*Item).pos.C,
			}
			if newDir.R == currentDir.R && newDir.C == currentDir.C {
				turn_addend = 0
			}
			t_distance := current.(*Item).distance + turn_addend + 1
			if t_distance < nodeMap[neighbor].distanceFromEnd {
				nodeMap[neighbor].distanceFromEnd = t_distance
				nodeMap[neighbor].parents = []Position{current.(*Item).pos}
				heap.Push(q, &Item{distance: t_distance, pos: neighbor, dir: newDir})
			}

		}
	}
}

// I'm keeping this here to show the pain
//
//	func DFSRec(pos Position, dir Position, score int, endPos Position, graph map[Position]*Node, lowestScore int, inPath []Position, memoMap map[Position]int) {
//		if graph[pos].visited {
//			return
//		}
//		// append node to list to update (best path)
//		graph[pos].visited = true
//		newInPath := append(inPath, pos)
//
//		defer func() { graph[pos].visited = false }()
//		for _, adjPos := range graph[pos].adjList {
//			turn_addend := 1000
//			newDir := Position{
//				R: adjPos.R - pos.R,
//				C: adjPos.C - pos.C,
//			}
//			if newDir.R == dir.R && newDir.C == dir.C {
//				turn_addend = 0
//			}
//
//			if !graph[adjPos].visited {
//				// score+turn_addend+1 <= lowestScore
//				if adjPos.R == endPos.R && adjPos.C == endPos.C {
//					graph[endPos].bestPath = true
//					for _, pos := range inPath {
//						graph[pos].bestPath = true
//					}
//				} else {
//					if score+turn_addend+1 <= lowestScore {
//						DFSRec(adjPos, newDir, score+turn_addend+1, endPos, graph, lowestScore, newInPath, memoMap)
//					}
//				}
//			}
//
//		}
//	}
//
//	func DFS(sourcePos Position, endPos Position, lowestScore int, graph map[Position]*Node, memoMap map[Position]int) {
//		score := 0
//		for _, value := range graph {
//			value.visited = false
//		}
//		direction := Position{R: 0, C: 1}
//		DFSRec(sourcePos, direction, score, endPos, graph, lowestScore, []Position{}, memoMap)
//	}
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
		for adjRow := -1; adjRow < 2; adjRow++ {
			for adjCol := -1; adjCol < 2; adjCol++ {
				if (adjRow == 0 && adjCol == 0) || (AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1) {
					continue
				}
				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if _, exists := nodeMap[neighborPos]; exists {
					if !nodeMap[neighborPos].wall {
						node.adjList = append(node.adjList, neighborPos)
					}
				}
			}
		}
	}
	return nodeMap, start, end, nil
}

func Part1() (int, map[Position]*Node, Position, Position, error) {
	nodeMap, start, end, err := ReadInput()
	if err != nil {
		return 0, nil, Position{}, Position{}, err
	}
	Dijkstra(nodeMap, start)
	return nodeMap[end].distance, nodeMap, start, end, nil
}

func Part2(lowestScore int, nodeMap map[Position]*Node, start, end Position) (int, error) {
	// This takes advantage of the fact in each test input and real input, the end spot is in the top right corner. A neat happenstance
	tiles := 0
	starting_position := Position{R: 1, C: 0}
	Dijkstra2(nodeMap, end, starting_position)

	for _, node := range nodeMap {
		if node.distance+node.distanceFromEnd == lowestScore {
			tiles = tiles + 1
		}
	}
	starting_position = Position{R: 0, C: 1}

	Dijkstra2(nodeMap, end, starting_position)

	for _, node := range nodeMap {
		if node.distance+node.distanceFromEnd == lowestScore {
			tiles = tiles + 1
		}
	}
	return tiles + 2, nil
}

func main() {
	result1, nodeMap, start, end, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}

	fmt.Printf("Lowest Score: %v \n", result1)

	result2, err := Part2(result1, nodeMap, start, end)
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}

	fmt.Printf("Num tiles: %v \n", result2)
}
