package day4

import (
	_ "embed"
	"fmt"
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
	for _, line := range splitted {
		nbPoints := 0
		parts := strings.Split(line, ":")
		numbers := strings.Split(parts[1], "|")
		for _, winning := range strings.Split(numbers[0], " ") {
			if winning == " " || winning == "" {
				continue
			}
			for _, got := range strings.Split(numbers[1], " ") {
				if got == " " || got == "" {
					continue
				}
				if winning == got {
					if nbPoints == 0 {
						nbPoints = 1
					} else {
						nbPoints <<= 1
					}
				}
			}
		}
		sum += nbPoints
	}

	return sum
}

func parseHard() int {
	sum := 0

	multiplicatorMap := make(map[int]int)

	splitted := strings.Split(input, "\n")
	for index, line := range splitted {
		nbWon := 0
		parts := strings.Split(line, ":")
		numbers := strings.Split(parts[1], "|")
		for _, winning := range strings.Split(numbers[0], " ") {
			if winning == " " || winning == "" {
				continue
			}
			for _, got := range strings.Split(numbers[1], " ") {
				if got == " " || got == "" {
					continue
				}
				if winning == got {
					nbWon += 1
				}
			}
		}
		mult := multiplicatorMap[index]
		if mult != 0 {
			sum += mult + 1
		} else {
			sum += 1
		}
		for i := index + 1; i < index+nbWon+1; i++ {
			if mult != 0 {
				multiplicatorMap[i] += mult + 1
			} else {
				multiplicatorMap[i] += 1
			}
		}
	}

	return sum
}
