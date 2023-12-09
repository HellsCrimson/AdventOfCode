package day3

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Main() {
	fmt.Println(parseSimple())
	fmt.Println(parseHard())
}

func parseSimple() int {
	sum := 0

	splitted := strings.Split(input, "\n")
	for indexY := 0; indexY < len(splitted); indexY++ {
		line := splitted[indexY]
		for indexX := 0; indexX < len(line); indexX++ {
			char := line[indexX]
			if char >= '0' && char <= '9' {
				if findAdjacent(splitted, indexX, indexY) {
					numberStr := getNumber(line, indexX)
					number, _ := strconv.Atoi(numberStr)
					sum += number
					indexX += len(numberStr)
				}
			}
		}
	}

	return sum
}

func parseHard() int {
	sum := 0

	splitted := strings.Split(input, "\n")
	for indexY := 0; indexY < len(splitted); indexY++ {
		line := splitted[indexY]
		for indexX := 0; indexX < len(line); indexX++ {
			char := line[indexX]
			if char == '*' {
				if found, numberStr, ratioStr := findAdjacentHard(splitted, indexX, indexY); found {
					fmt.Println("Found: ", numberStr, ratioStr)
					number, _ := strconv.Atoi(numberStr)
					ratio, _ := strconv.Atoi(ratioStr)
					sum += number * ratio
				}
			}
		}
	}

	return sum
}

func findAdjacent(arr []string, indexX int, indexY int) bool {
	for i := indexX; i < len(arr[indexY]) && arr[indexY][i] >= '0' && arr[indexY][i] <= '9'; i++ {
		startY := indexY - 1
		if startY < 0 {
			startY = 0
		}
		for y := startY; y < len(arr) && y <= indexY+1; y++ {
			startX := i - 1
			if startX < 0 {
				startX = 0
			}
			for x := startX; x < len(arr[y]) && x <= i+1; x++ {
				char := arr[y][x]
				if char != '.' && !(char >= '0' && char <= '9') {
					return true
				}
			}
		}
	}
	return false
}

func getNumber(line string, index int) string {
	number := ""
	for i := index; i < len(line); i++ {
		char := line[i]
		if char >= '0' && char <= '9' {
			number += string(char)
		} else {
			break
		}
	}
	return number
}

func findAdjacentHard(arr []string, indexX int, indexY int) (bool, string, string) {
	startY, startX := findNumber(arr, indexY, indexX, -1, -1)
	if startY == -1 || startX == -1 {
		return false, "", ""
	}
	number1 := getNumberHard(arr[startY], startX)

	startY, startX = findNumber(arr, indexY, indexX, startX, startY)
	if startY == -1 || startX == -1 {
		return false, "", ""
	}
	number2 := getNumberHard(arr[startY], startX)

	return true, number1, number2
}

func getNumberHard(line string, index int) string {
	startIndex := index
	for ; startIndex >= 0; startIndex-- {
		char := line[startIndex]
		if char >= '0' && char <= '9' {
			continue
		} else {
			break
		}
	}
	startIndex += 1

	number := ""
	for i := startIndex; i < len(line); i++ {
		char := line[i]
		if char >= '0' && char <= '9' {
			number += string(char)
		} else {
			break
		}
	}
	return number
}

func findNumber(arr []string, centerY int, centerX int, whitelistX int, whitelistY int) (int, int) {
	startY := centerY - 1
	if startY < 0 {
		startY = 0
	}

	for y := startY; y < len(arr) && y <= centerY+1; y++ {
		startX := centerX - 1
		if startX < 0 {
			startX = 0
		}

		for x := startX; x < len(arr[y]) && x <= centerX+1; x++ {
			char := arr[y][x]
			if char >= '0' && char <= '9' {
				if whitelistX == -1 {
					return y, x
				}
				if y != whitelistY || !isSameNumber(arr, y, x, whitelistX) {
					return y, x
				}
			}
		}
	}
	return -1, -1
}

func isSameNumber(arr []string, indexY int, indexX1 int, indexX2 int) bool {
	i1 := findStartIndex(arr, indexY, indexX1)
	i2 := findStartIndex(arr, indexY, indexX2)
	return i1 == i2
}

func findStartIndex(arr []string, indexY int, indexX int) int {
	startIndex := indexX
	for ; startIndex >= 0; startIndex-- {
		char := arr[indexY][startIndex]
		if char >= '0' && char <= '9' {
			continue
		} else {
			break
		}
	}
	return startIndex + 1
}
