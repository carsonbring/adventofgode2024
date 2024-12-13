// I know this is really sloppy its 2:30 AM and i have work tomorrow
// at least it works
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	adjList []*Node
	visited bool
	grouped bool
	value   rune
	fences  int
	row     int
	col     int
	ogrow   int
	ogcol   int
}

func PrintLayout(layout [][]Node) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, row := range layout {
		for _, element := range row {
			_, err := file.WriteString(string(element.value) + " ")
			if err != nil {
				return err
			}
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}

	}
	return nil
}

func checkBounds(row int, col int, matrix [][]Node) bool {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[row]) {
		return false
	} else {
		return true
	}
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func convertRunesToNodes(runes []rune, rowCount int) []Node {
	nodes := []Node{}
	for i, rune := range runes {
		nodes = append(nodes, Node{
			adjList: []*Node{},
			value:   rune,
			visited: false,
			fences:  0,
			row:     rowCount,
			col:     i,
		})
	}
	return nodes
}

func numFences(sourceNode *Node) int {
	sum := 0
	sum = 4 - len(sourceNode.adjList)
	for _, node := range sourceNode.adjList {
		if node.value != sourceNode.value {
			sum = sum + 1
		}
	}
	return sum
}

func numSame(sourceNode *Node) int {
	sum := 0
	for _, node := range sourceNode.adjList {
		if node.value == sourceNode.value {
			sum = sum + 1
		}
	}
	return sum
}

func DFSRec(sourceNode *Node, perimeter *int, area *int) {
	if sourceNode.visited {
		return
	}
	value := sourceNode.value
	sourceNode.visited = true
	sourceNode.grouped = true
	fences := numFences(sourceNode)
	*perimeter = *perimeter + fences
	*area = *area + 1
	for _, node := range sourceNode.adjList {
		if !node.visited && node.value == value {
			DFSRec(node, perimeter, area)
		}
	}
}

func DFS(sourceNode *Node, graph []*Node) (int, int) {
	area := 0
	perimeter := 0
	for i := range graph {
		graph[i].visited = false
	}
	DFSRec(sourceNode, &perimeter, &area)
	return area, perimeter
}

func isInsideCorner(sourceNode *Node, ogMatrix [][]Node) int {
	sourceVal := sourceNode.value
	corners := 0

	horizontalCols := []int{}
	verticalRows := []int{}

	for _, adj := range sourceNode.adjList {
		if adj.value == sourceVal {
			if adj.row == sourceNode.row {
				horizontalCols = append(horizontalCols, adj.col)
			} else if adj.col == sourceNode.col {
				verticalRows = append(verticalRows, adj.row)
			}
		}
	}

	for _, vertRow := range verticalRows {
		for _, horizCol := range horizontalCols {
			if checkBounds(vertRow, horizCol, ogMatrix) {
				diagonalVal := ogMatrix[vertRow][horizCol].value
				if diagonalVal != sourceVal {
					corners++
				}
			}
		}
	}

	return corners
}

func DFSRec2(sourceNode *Node, fenceside *int, inside *int, area *int, matrix [][]Node, ogMatrix [][]Node) {
	if sourceNode.visited {
		return
	}
	value := sourceNode.value
	sourceNode.visited = true
	sourceNode.grouped = true
	fences := numFences(sourceNode)
	if fences == 2 {
		*fenceside = *fenceside + 1
	}
	*area = *area + 1
	for _, node := range sourceNode.adjList {
		if !node.visited && node.value == value {
			DFSRec2(node, fenceside, inside, area, matrix, ogMatrix)
		}
	}
}

func DFSRecInside2(sourceNode *Node, inside *int, matrix [][]Node) {
	if sourceNode.visited {
		return
	}
	value := sourceNode.value
	sourceNode.visited = true
	sourceNode.grouped = true
	same := numSame(sourceNode)
	if same >= 2 {
		*inside = *inside + isInsideCorner(sourceNode, matrix)
	}
	for _, node := range sourceNode.adjList {
		if !node.visited && node.value == value {
			DFSRecInside2(node, inside, matrix)
		}
	}
}

func DFS2(sourceNode *Node, graph []*Node, matrix [][]Node, ogMatrix [][]Node) (int, int) {
	area := 0
	fenceside := 0
	inside := 0
	for i := range graph {
		graph[i].visited = false
	}
	DFSRec2(sourceNode, &fenceside, &inside, &area, matrix, ogMatrix)
	DFSRecInside2(&ogMatrix[sourceNode.ogrow][sourceNode.ogcol], &inside, ogMatrix)

	sides := fenceside + inside
	return sides, area
}

func ReadInput() ([]*Node, [][]Node, error) {
	matrix := [][]Node{}
	graph := []*Node{}

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		matrix = append(matrix, convertRunesToNodes(runes, rowCount))
		rowCount++
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
	return graph, matrix, nil
}

func ReadInput2() ([]*Node, [][]Node, error) {
	matrix := [][]Node{}
	graph := []*Node{}

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		matrix = append(matrix, convertRunesToNodes(runes, rowCount))
		rowCount++
	}
	matrix = InflateMatrix(matrix)
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
	return graph, matrix, nil
}

func InflateMatrix(matrix [][]Node) [][]Node {
	newMatrix := [][]Node{}

	for i, row := range matrix {
		dupedRow := []Node{}
		for j, node := range row {
			copy1 := Node{
				adjList: []*Node{},
				value:   node.value,
				visited: false,
				fences:  node.fences,
				row:     i * 2,
				col:     j * 2,
				ogrow:   i,
				ogcol:   j,
			}
			copy2 := Node{
				adjList: []*Node{},
				value:   node.value,
				visited: false,
				fences:  node.fences,
				row:     i * 2,
				col:     j*2 + 1,
				ogrow:   i,
				ogcol:   j,
			}
			dupedRow = append(dupedRow, copy1, copy2)
		}

		newMatrix = append(newMatrix, dupedRow)

		duplicatedRow := make([]Node, len(dupedRow))
		for k, node := range dupedRow {
			duplicatedRow[k] = Node{
				adjList: []*Node{},
				value:   node.value,
				visited: false,
				fences:  node.fences,
				row:     node.row + 1,
				col:     node.col,
				ogrow:   node.ogrow,
				ogcol:   node.ogcol,
			}
		}
		newMatrix = append(newMatrix, duplicatedRow)
	}

	return newMatrix
}

func Part1() (int, error) {
	price := 0
	graph, _, err := ReadInput()
	if err != nil {
		return 0, err
	}
	for _, node := range graph {
		if !node.grouped {
			area, perimeter := DFS(node, graph)
			price = price + (area * perimeter)

		}
	}
	return price, nil
}

func Part2() (int, error) {
	price := 0
	_, ogMatrix, err := ReadInput()
	if err != nil {
		return 0, err
	}
	graph, matrix, err := ReadInput2()
	if err != nil {
		return 0, err
	}
	for _, node := range graph {
		if !node.grouped {
			sides, area := DFS2(node, graph, matrix, ogMatrix)
			price = price + (sides * (area / 4))
		}
	}
	return price, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part 1: %v", err)
	}

	fmt.Printf("Total Price: %v \n", result1)
	result2, err := Part2()
	if err != nil {
		fmt.Printf("Error in Part 2: %v", err)
	}

	fmt.Printf("Total Price (with discount): %v \n", result2)
}
