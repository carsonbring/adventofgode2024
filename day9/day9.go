package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Block struct {
	space       []int
	empty       bool
	emptySpaces int
	id          int
}

func AbsInt(num int) int {
	if num < 0 {
		num = -num
	}
	return num
}

func createSlice(element int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = element
	}
	return result
}

func ReadInput() ([]Block, error) {
	var blocks []Block

	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		for i, rune := range runes {
			if rune < '0' || rune > '9' {
				return nil, fmt.Errorf("rune is not a numeric digit: %c", rune)
			}
			num := AbsInt(int(rune - '0'))

			if i%2 == 0 {
				blocks = append(blocks, Block{
					empty:       false,
					emptySpaces: 0,
					space:       createSlice(i/2, num),
					id:          i,
				})
			} else {
				blocks = append(blocks, Block{
					empty:       true,
					emptySpaces: num,
					space:       createSlice(-1, num),
					id:          i,
				})
			}
		}
	}
	return blocks, nil
}

func fillSpace(end *int, block *Block, blocks *[]Block) {
	for i := 0; i < len(block.space); i++ {
		for j := len((*blocks)[*end].space) - 1; j >= 0; j-- {
			if j == -1 {
				break
			}
			if (*blocks)[*end].space[j] != -1 {
				block.space[i] = (*blocks)[*end].space[j]
				(*blocks)[*end].space[j] = -1
				break
			}
		}
		if block.space[i] == -1 {
			(*blocks)[*end].empty = true
			*end = *end - 1
			i--
		}
	}
}

func fillEntireSpace(end *int, blocks *[]Block, emptyBlockSizes map[int][]*Block, largestSize int) {
	originalLength := len((*blocks)[*end].space)
	for length := len((*blocks)[*end].space); length <= largestSize; length++ {
		if freeSpace, exists := emptyBlockSizes[length]; exists {
			if !exists || len(freeSpace) == 0 {
				continue
			}
			options := append([]*Block(nil), freeSpace...)
			for i := length + 1; i <= largestSize; i++ {
				if additionalSpace, exists := emptyBlockSizes[i]; exists {
					options = append(options, additionalSpace...)
				}
			}
			cleanedOptions := []*Block{}
			for _, option := range options {
				if (*blocks)[*end].id > option.id {
					cleanedOptions = append(cleanedOptions, option)
				}
			}
			if len(cleanedOptions) == 0 {
				break
			}
			sort.Slice(cleanedOptions, func(i, j int) bool {
				return cleanedOptions[i].id > cleanedOptions[j].id
			})
			best_option := cleanedOptions[len(cleanedOptions)-1]
			best_size := best_option.emptySpaces
			emptyBlockSizes[best_size] = emptyBlockSizes[best_size][:len(emptyBlockSizes[best_size])-1]

			for i := 0; i < originalLength; i++ {
				for j := 0; j < len(best_option.space); j++ {
					if best_option.space[j] == -1 {
						best_option.space[j] = (*blocks)[*end].space[i]
						(*blocks)[*end].space[i] = -1
						best_option.emptySpaces = best_option.emptySpaces - 1
						break
					}
				}
			}
			emptyBlockSizes[best_option.emptySpaces] = append(emptyBlockSizes[best_option.emptySpaces], best_option)

			sort.Slice(emptyBlockSizes[best_option.emptySpaces], func(i, j int) bool {
				return emptyBlockSizes[best_option.emptySpaces][i].id > emptyBlockSizes[best_option.emptySpaces][j].id
			})

			(*blocks)[*end].empty = true
			break
		}
	}
}

func Part1() (int, error) {
	blocks, err := ReadInput()
	checksum := 0
	done := false
	if err != nil {
		return 0, err
	}
	endPoint := len(blocks) - 1
	blockPosition := 0
	for i := 0; i != -1; i++ {
		if i%2 != 0 {
			fillSpace(&endPoint, &blocks[i], &blocks)
			if blocks[i+1].empty {
				break
			}
		}
		for _, element := range blocks[i].space {
			if element < 0 {
				done = true
				break
			}
			addition := element * blockPosition
			checksum = checksum + addition
			blockPosition++
		}
		if done {
			break
		}
	}
	return checksum, nil
}

func Part2() (int, error) {
	blocks, err := ReadInput()
	var sortedBlocks []*Block
	emptyBlockSizes := map[int][]*Block{}
	largestSize := 0
	for i := range blocks {
		block := &blocks[i]
		if !block.empty || block.emptySpaces == 0 {
			continue
		}
		if block.emptySpaces > largestSize {
			largestSize = block.emptySpaces
		}
		if sizelist, exists := emptyBlockSizes[block.emptySpaces]; exists {
			emptyBlockSizes[block.emptySpaces] = append(sizelist, block)
		} else {
			emptyBlockSizes[block.emptySpaces] = []*Block{block}
		}
	}

	for _, value := range emptyBlockSizes {
		sort.Slice(value, func(i, j int) bool {
			return value[i].id > value[j].id
		})
	}

	sort.Slice(sortedBlocks, func(i, j int) bool {
		return sortedBlocks[i].emptySpaces < sortedBlocks[j].emptySpaces
	})

	checksum := 0
	if err != nil {
		return 0, err
	}
	for endPoint := len(blocks) - 1; endPoint >= 0; endPoint-- {
		if !blocks[endPoint].empty {
			fillEntireSpace(&endPoint, &blocks, emptyBlockSizes, largestSize)
		}
	}
	blockPosition := 0

	for i := 0; i < len(blocks); i++ {
		for _, element := range blocks[i].space {
			if element < 0 {
				blockPosition++
				continue
			}
			addition := element * blockPosition
			checksum = checksum + addition
			blockPosition++
		}
	}
	return checksum, nil
}

func main() {
	result1, err := Part1()
	if err != nil {
		fmt.Printf("error in part1: %v \n", err)
	}
	fmt.Printf("Part1 checksum: %v \n", result1)

	result2, err := Part2()
	if err != nil {
		fmt.Printf("error in part2: %v \n", err)
	}
	fmt.Printf("Part2 checksum (whole files): %v \n", result2)
}
