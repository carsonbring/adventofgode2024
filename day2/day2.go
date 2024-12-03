package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func Part1() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	safeReportNum := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		sign := 0
		currentStatus := true

		for i := 0; i < len(line)-1; i = i + 1 {

			currentNum, err := strconv.Atoi(line[i])
			nextNum, err2 := strconv.Atoi(line[i+1])
			if err != nil || err2 != nil {
				return 0, fmt.Errorf("input couldn't be read. integer conversion failed")
			}
			difference := nextNum - currentNum

			if AbsInt(difference) > 3 || AbsInt(difference) < 1 {
				currentStatus = false
				break
			}

			if i == 0 {
				if AbsInt(difference) == difference {
					sign = 1
				} else {
					sign = -1
				}
			} else if Sign(difference) != sign {
				currentStatus = false
				break
			}
		}
		if currentStatus {
			safeReportNum = safeReportNum + 1
		}

	}
	return safeReportNum, nil
}

func SafetyCheck(line []string, strike bool) (int, error) {
	sign := 0
	for i := 0; i < len(line)-1; i = i + 1 {

		currentNum, err := strconv.Atoi(line[i])
		nextNum, err2 := strconv.Atoi(line[i+1])
		if err != nil || err2 != nil {
			return 0, fmt.Errorf("input couldn't be read. integer conversion failed")
		}
		difference := nextNum - currentNum

		if AbsInt(difference) > 3 || AbsInt(difference) < 1 {
			if !strike {
				saferLine := append(append([]string{}, line[:i]...), line[i+1:]...)

				strikeResponse, err := SafetyCheck(saferLine, true)
				if err != nil {
					return 0, err
				}

				if strikeResponse == 1 {
					return strikeResponse, nil
				} else {

					saferLine2 := append(append([]string{}, line[:i+1]...), line[i+2:]...)

					strikeResponse2, err := SafetyCheck(saferLine2, true)
					if err != nil {
						return 0, err
					}
					if strikeResponse2 == 1 {
						return strikeResponse2, nil
					} else if i-1 >= 0 {

						saferLine3 := append(append([]string{}, line[:i-1]...), line[i:]...)

						strikeResponse3, err := SafetyCheck(saferLine3, true)
						if err != nil {
							return 0, err
						}
						return strikeResponse3, nil
					} else {
						return 0, nil
					}
				}
			} else {
				return 0, nil
			}
		}

		if i == 0 {
			if AbsInt(difference) == difference {
				sign = 1
			} else {
				sign = -1
			}
		} else if Sign(difference) != sign {
			if !strike {
				saferLine := append(append([]string{}, line[:i]...), line[i+1:]...)

				strikeResponse, err := SafetyCheck(saferLine, true)
				if err != nil {
					return 0, err
				}

				if strikeResponse == 1 {
					return strikeResponse, nil
				} else {

					saferLine2 := append(append([]string{}, line[:i+1]...), line[i+2:]...)

					strikeResponse2, err := SafetyCheck(saferLine2, true)
					if err != nil {
						return 0, err
					}
					if strikeResponse2 == 1 {
						return strikeResponse2, nil
					} else if i-1 >= 0 {

						saferLine3 := append(append([]string{}, line[:i-1]...), line[i:]...)

						strikeResponse3, err := SafetyCheck(saferLine3, true)
						if err != nil {
							return 0, err
						}
						return strikeResponse3, nil
					} else {
						return 0, nil
					}

				}
			} else {
				return 0, nil
			}
		}
	}
	return 1, nil
}

func Part2() (int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return 0, fmt.Errorf("couldn't read input file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	safeReportNum := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		result, err := SafetyCheck(line, false)
		if err != nil {
			return 0, err
		}
		safeReportNum = safeReportNum + result
	}
	return safeReportNum, nil
}

func main() {
	answer1, err := Part1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of safe reports (part 1): %v\n\n", answer1)
	answer2, err := Part2()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of safe reports (part 2): %v\n", answer2)
}
