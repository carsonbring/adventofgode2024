package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	adjList []*Node
	visited bool
	id      int
	value   int
	rating  int
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func convertRunesToNodes(runes []rune, idCounter *int) []Node {
	nodes := []Node{}
	for _, rune := range runes {
		nodes = append(nodes, Node{
			id:      *idCounter,
			adjList: []*Node{},
			value:   int(rune - '0'),
			rating:  0,
		})
		(*idCounter)++
	}
	return nodes
}

func checkBounds(row int, col int, matrix [][]Node) bool {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[row]) {
		return false
	} else {
		return true
	}
}

func ReadInput() ([]*Node, error) {
	matrix := [][]Node{}
	graph := []*Node{}

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	idCounter := 0

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		matrix = append(matrix, convertRunesToNodes(runes, &idCounter))
	}
	for i, row := range matrix {
		for j := range row {
			for adjRow := -1; adjRow < 2; adjRow++ {
				for adjCol := -1; adjCol < 2; adjCol++ {
					if (adjRow == 0 && adjCol == 0) || (AbsInt(adjRow) == 1 && AbsInt(adjCol) == 1) {
						continue
					}
					neighborRow := i + adjRow
					neighborCol := j + adjCol
					if checkBounds(neighborRow, neighborCol, matrix) {
						matrix[i][j].adjList = append(matrix[i][j].adjList, &matrix[neighborRow][neighborCol])
					}
				}
			}
			graph = append(graph, &matrix[i][j])
		}
	}
	return graph, nil
}

func DFSRec(sourceNode *Node, score *int) {
	if sourceNode.visited {
		return
	}
	sourceNode.visited = true
	for _, node := range sourceNode.adjList {
		if !node.visited && node.value == sourceNode.value+1 {
			if node.value == 9 {
				(*score)++
			}
			DFSRec(node, score)
		}
	}
}

func DFS(sourceNode *Node, graph []*Node) int {
	score := 0
	for i := range graph {
		graph[i].visited = false
	}
	DFSRec(sourceNode, &score)
	return score
}

func DFSRec2(sourceNode *Node, score *int) {
	if sourceNode.visited {
		return
	}
	sourceNode.visited = true
	defer func() { sourceNode.visited = false }()

	for _, node := range sourceNode.adjList {
		if !node.visited && node.value == sourceNode.value+1 {
			if node.value == 9 {
				(*score)++
			} else {
				DFSRec2(node, score)
			}
		}
	}
}

func DFS2(sourceNode *Node, graph []*Node) int {
	score := 0
	for i := range graph {
		graph[i].visited = false
	}
	DFSRec2(sourceNode, &score)
	return score
}

func Part1() (int, error) {
	score := 0
	graph, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for i := range graph {
		node := graph[i]
		if node.value == 0 {
			score = score + DFS(node, graph)
		}
	}
	return score, nil
}

func Part2() (int, error) {
	score := 0
	graph, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for i := range graph {
		node := graph[i]
		if node.value == 0 {
			score = score + DFS2(node, graph)
		}
	}

	return score, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in part1: %v \n", err)
	}

	fmt.Printf("Trailhead scores (part1): %v \n", result1)

	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in part2: %v \n", err)
	}

	fmt.Printf("Trailhead scores (part2): %v \n", result2)
}
