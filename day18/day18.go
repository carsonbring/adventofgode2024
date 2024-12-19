package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	distance int
}
type Position struct {
	R, C int
}
type Node struct {
	adjList   []Position
	corrupted bool
	visited   bool
	finish    bool
	distance  int
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func Dijkstra(nodeMap map[Position]*Node, source Position) {
	for _, node := range nodeMap {
		node.distance = 10000000
	}
	q := &PriorityQueue{&Item{pos: source, distance: 0}}
	heap.Init(q)
	for q.Len() != 0 {
		current := q.Pop()
		for _, neighbor := range nodeMap[current.(*Item).pos].adjList {
			t_distance := current.(*Item).distance + 1
			if t_distance < nodeMap[neighbor].distance {
				nodeMap[neighbor].distance = t_distance
				heap.Push(q, &Item{distance: t_distance, pos: neighbor})
			}

		}
	}
}

func ReadInput(bytes int, length int) (map[Position]*Node, []Position, error) {
	grid := map[Position]*Node{}

	restOfCorruption := []Position{}
	for i := 0; i <= length; i++ {
		for j := 0; j <= length; j++ {
			grid[Position{R: i, C: j}] = &Node{
				adjList:   []Position{},
				corrupted: false,
				distance:  100000000,
				finish:    false,
				visited:   false,
			}
		}
	}
	grid[Position{R: length, C: length}].finish = true
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {

		intSlice := []int{}
		line := scanner.Text()
		strs := strings.Split(line, ",")
		for _, str := range strs {
			intOp := -1
			intOp, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("error in recursion %v", err)
			}
			intSlice = append(intSlice, intOp)
		}
		if rowCount >= bytes {
			restOfCorruption = append(restOfCorruption, Position{R: intSlice[1], C: intSlice[0]})
		} else {
			grid[Position{R: intSlice[1], C: intSlice[0]}].corrupted = true
		}
		rowCount++
	}

	for pos, node := range grid {
		for adjRow := -1; adjRow < 2; adjRow++ {
			for adjCol := -1; adjCol < 2; adjCol++ {
				if (adjRow == 0 && adjCol == 0) || (AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1) {
					continue
				}
				neighborPos := Position{R: pos.R + adjRow, C: pos.C + adjCol}
				if _, exists := grid[neighborPos]; exists {
					if !grid[neighborPos].corrupted {
						node.adjList = append(node.adjList, neighborPos)
					}
				}
			}
		}
	}

	return grid, restOfCorruption, nil
}

func DFSRec(currentPos Position, endPos Position, nodeMap map[Position]*Node, finished *bool) {
	if *finished {
		return
	}
	nodeMap[currentPos].visited = true
	for _, node := range nodeMap[currentPos].adjList {
		if !nodeMap[node].visited && !nodeMap[node].corrupted {
			if currentPos.C == endPos.C && currentPos.R == endPos.R {
				*finished = true
				return
			} else {
				DFSRec(node, endPos, nodeMap, finished)
			}
		}
	}
}

func DFS(nodeMap map[Position]*Node, startPos Position, endPos Position) bool {
	for _, val := range nodeMap {
		val.visited = false
	}
	finished := false
	DFSRec(startPos, endPos, nodeMap, &finished)
	return finished
}

func Part1() (int, error) {
	length := 70
	grid, _, err := ReadInput(1024, length)
	if err != nil {
		return 0, err
	}
	Dijkstra(grid, Position{R: 0, C: 0})
	return grid[Position{R: length, C: length}].distance, nil
}

func Part2() (Position, error) {
	length := 70
	grid, restOfCorruption, err := ReadInput(1024, length)
	if err != nil {
		return Position{}, err
	}

	currentPos := Position{}
	for _, position := range restOfCorruption {
		currentPos = position
		grid[position].corrupted = true
		result := DFS(grid, Position{0, 0}, Position{length, length})
		if !result {
			break
		}

	}
	return currentPos, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("Tiles to exit: %v \n", result1)

	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}
	fmt.Printf("First impasse byte corruption: %v,%v\n", result2.C, result2.R)
}
