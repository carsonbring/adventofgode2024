package main

import (
	"bufio"
	"fmt"
	"os"
)

type BoxPosition struct {
	R      int
	leftC  int
	rightC int
}
type Tile struct {
	bot      bool
	occupied bool
	wall     bool
	leftbox  bool

	rightbox bool
}

type Box struct {
	Tile,
	left bool
	right bool
}

type Velocity struct {
	row int
	col int
}

func PrintWarehouse(warehouse [][]Tile) {
	for _, row := range warehouse {
		for _, tile := range row {
			if tile.bot {
				fmt.Printf("@")
			} else if tile.leftbox {
				fmt.Printf("[")
			} else if tile.rightbox {
				fmt.Printf("]")
			} else if tile.wall {
				fmt.Printf("#")
			} else if !tile.occupied {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func ReadInput() ([][]Tile, []Velocity, int, int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, 0, 0, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	botR := 0
	rowCount := 0
	botC := 0
	warehouse := [][]Tile{}
	directions := []Velocity{}

	dirs := false
	for scanner.Scan() {
		newRow := []Tile{}
		line := scanner.Text()
		if line == "\n" || len(line) == 0 {
			dirs = true
			continue
		}
		if !dirs {
			for i, r := range line {
				newTile := Tile{}
				switch r {
				case '#':
					newTile = Tile{
						bot:      false,
						occupied: true,
						wall:     true,
					}
				case 'O':
					newTile = Tile{
						bot:      false,
						occupied: true,
						wall:     false,
					}
				case '.':
					newTile = Tile{
						bot: false,

						occupied: false,
						wall:     false,
					}
				default:
					newTile = Tile{
						bot:      true,
						occupied: true,
						wall:     false,
					}
					botR = rowCount
					botC = i

				}
				newRow = append(newRow, newTile)
			}
			warehouse = append(warehouse, newRow)
			rowCount++
		} else {
			for _, r := range line {
				newVel := Velocity{}
				switch r {
				case '<':
					newVel = Velocity{
						row: 0,
						col: -1,
					}
				case '>':
					newVel = Velocity{
						row: 0,
						col: 1,
					}
				case '^':
					newVel = Velocity{
						row: -1,
						col: 0,
					}
				case 'v':
					newVel = Velocity{
						row: 1,
						col: 0,
					}

				}
				directions = append(directions, newVel)

			}
		}
	}
	return warehouse, directions, botR, botC, nil
}

func ReadInput2() ([][]Tile, []Velocity, int, int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, 0, 0, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	botR := 0
	rowCount := 0
	botC := 0
	warehouse := [][]Tile{}
	directions := []Velocity{}

	dirs := false
	for scanner.Scan() {
		newRow := []Tile{}
		line := scanner.Text()
		if line == "\n" || len(line) == 0 {
			dirs = true
			continue
		}
		if !dirs {
			for i, r := range line {
				switch r {
				case '#':
					newRow = append(newRow, Tile{occupied: true, wall: true}, Tile{occupied: true, wall: true})
				case 'O':
					newRow = append(newRow, Tile{leftbox: true, occupied: true}, Tile{rightbox: true, occupied: true})
				case '.':
					newRow = append(newRow, Tile{}, Tile{})
				default:
					botR = rowCount
					botC = i * 2
					newRow = append(newRow, Tile{bot: true, occupied: true}, Tile{})

				}
			}
			warehouse = append(warehouse, newRow)
			rowCount++
		} else {
			for _, r := range line {
				newVel := Velocity{}
				switch r {
				case '<':
					newVel = Velocity{
						row: 0,
						col: -1,
					}
				case '>':
					newVel = Velocity{
						row: 0,
						col: 1,
					}
				case '^':
					newVel = Velocity{
						row: -1,
						col: 0,
					}
				case 'v':
					newVel = Velocity{
						row: 1,
						col: 0,
					}

				}
				directions = append(directions, newVel)

			}
		}
	}
	return warehouse, directions, botR, botC, nil
}

func MoveBoxes(inverseVel Velocity, warehouse *[][]Tile, currentR int, currentC int, numBoxes int) {
	if numBoxes == 0 {
		(*warehouse)[currentR][currentC] = Tile{bot: true, occupied: true, wall: false}
		(*warehouse)[currentR+inverseVel.row][currentC+inverseVel.col] = Tile{bot: false, occupied: false, wall: false}
		return
	}
	(*warehouse)[currentR][currentC] = Tile{
		leftbox:  (*warehouse)[currentR+inverseVel.row][currentC+inverseVel.col].leftbox,
		rightbox: (*warehouse)[currentR+inverseVel.row][currentC+inverseVel.col].rightbox,
		bot:      false,

		occupied: true,
		wall:     false,
	}
	MoveBoxes(inverseVel, warehouse, currentR+inverseVel.row, currentC+inverseVel.col, numBoxes-1)
}

func MoveBot(botR, botC int, velocity Velocity, warehouse *[][]Tile) (int, int) {
	targetR := botR + velocity.row
	targetC := botC + velocity.col
	targetTile := (*warehouse)[targetR][targetC]
	boxesToMove := 0
	for {
		if targetTile.wall {
			return botR, botC
		} else if !targetTile.occupied {
			break
		} else {
			boxesToMove++
			targetR = targetR + velocity.row
			targetC = targetC + velocity.col
			targetTile = (*warehouse)[targetR][targetC]
		}
	}

	inverseVel := Velocity{
		row: -1 * velocity.row,
		col: -1 * velocity.col,
	}
	MoveBoxes(inverseVel, warehouse, targetR, targetC, boxesToMove)
	return botR + velocity.row, botC + velocity.col
}

func MoveBoxes2(velocity Velocity, warehouse *[][]Tile, targetR int, targetC int, boxPositions []BoxPosition) {
	if len(boxPositions) == 0 {
		(*warehouse)[targetR][targetC] = Tile{bot: true, occupied: true, wall: false}
		(*warehouse)[targetR-velocity.row][targetC-velocity.col] = Tile{bot: false, occupied: false, wall: false}
		return
	}
	(*warehouse)[boxPositions[0].R+velocity.row][boxPositions[0].leftC] = Tile{leftbox: true, occupied: true}
	(*warehouse)[boxPositions[0].R+velocity.row][boxPositions[0].rightC] = Tile{rightbox: true, occupied: true}

	MoveBoxes2(velocity, warehouse, targetR, targetC, boxPositions[1:])
}

func checkMove(boxPosition BoxPosition, velocity Velocity, warehouse *[][]Tile) ([]BoxPosition, bool) {
	targetTileL := (*warehouse)[boxPosition.R+velocity.row][boxPosition.leftC]
	targetTileR := (*warehouse)[boxPosition.R+velocity.row][boxPosition.rightC]
	if targetTileL.wall || targetTileR.wall {
		return nil, false
	} else if !targetTileL.occupied && !targetTileR.occupied {
		return []BoxPosition{boxPosition}, true
	} else {
		boxPositions := []BoxPosition{}
		if targetTileL.leftbox {
			newPosition := BoxPosition{R: boxPosition.R + velocity.row, leftC: boxPosition.leftC, rightC: boxPosition.rightC}
			positions, result := checkMove(newPosition, velocity, warehouse)
			if !result {
				return nil, false
			}

			boxPositions = append(boxPositions, newPosition)
			boxPositions = append(boxPositions, positions...)
		}
		if targetTileL.rightbox {
			newPosition := BoxPosition{R: boxPosition.R + velocity.row, leftC: boxPosition.leftC - 1, rightC: boxPosition.leftC}
			positions, result := checkMove(newPosition, velocity, warehouse)
			if !result {
				return nil, false
			}
			boxPositions = append(boxPositions, newPosition)
			boxPositions = append(boxPositions, positions...)

		}
		if targetTileR.leftbox {
			newPosition := BoxPosition{R: boxPosition.R + velocity.row, leftC: boxPosition.rightC, rightC: boxPosition.rightC + 1}
			positions, result := checkMove(newPosition, velocity, warehouse)
			if !result {
				return nil, false
			}

			boxPositions = append(boxPositions, newPosition)
			boxPositions = append(boxPositions, positions...)

		}
		return boxPositions, true

	}
}

func MoveBot2(botR, botC int, velocity Velocity, warehouse *[][]Tile) (int, int) {
	targetR := botR + velocity.row
	targetC := botC + velocity.col
	targetTile := (*warehouse)[targetR][targetC]
	boxesToMove := 0
	if velocity.row == 0 {
		for {
			if targetTile.wall {
				return botR, botC
			} else if !targetTile.occupied {
				break
			} else {
				boxesToMove++
				targetR = targetR + velocity.row
				targetC = targetC + velocity.col
				targetTile = (*warehouse)[targetR][targetC]
			}
		}
		inverseVel := Velocity{
			row: -1 * velocity.row,
			col: -1 * velocity.col,
		}
		MoveBoxes(inverseVel, warehouse, targetR, targetC, boxesToMove)

	} else {
		boxPositionsToMove, result := checkMove(BoxPosition{R: botR, leftC: botC, rightC: botC}, velocity, warehouse)
		if !result {
			return botR, botC
		}
		for _, pos := range boxPositionsToMove {
			(*warehouse)[pos.R][pos.leftC] = Tile{}
			(*warehouse)[pos.R][pos.rightC] = Tile{}
		}
		MoveBoxes2(velocity, warehouse, targetR, targetC, boxPositionsToMove)

	}

	return botR + velocity.row, botC + velocity.col
}

func Part1() (int, error) {
	warehouse, directions, botR, botC, err := ReadInput()
	sumOfBoxGPS := 0
	if err != nil {
		return 0, err
	}
	for _, vel := range directions {
		botR, botC = MoveBot(botR, botC, vel, &warehouse)
	}

	for i, row := range warehouse {
		for j, tile := range row {
			if tile.occupied && !tile.wall && !tile.bot {
				sumOfBoxGPS = sumOfBoxGPS + (100*i + j)
			}
		}
	}
	return sumOfBoxGPS, nil
}

func Part2() (int, error) {
	warehouse, directions, botR, botC, err := ReadInput2()

	sumOfBoxGPS := 0
	if err != nil {
		return 0, err
	}
	for _, vel := range directions {
		botR, botC = MoveBot2(botR, botC, vel, &warehouse)
	}

	for i, row := range warehouse {
		for j, tile := range row {
			if tile.occupied && tile.leftbox && !tile.wall && !tile.bot {
				sumOfBoxGPS = sumOfBoxGPS + (100*i + j)
			}
		}
	}
	return sumOfBoxGPS, nil
}

func main() {
	answer1, err := Part1()
	if err != nil {
		fmt.Printf("Error in Part1: %v \n", err)
	}

	fmt.Printf("Sum of all boxes' GPS coords (part1): %v \n", answer1)

	answer2, err := Part2()
	if err != nil {
		fmt.Printf("Error in Part2", err)
	}

	fmt.Printf("SUm of all boxes' GPS coords (part2): %v \n", answer2)
}
