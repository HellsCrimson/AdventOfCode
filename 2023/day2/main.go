package day2

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Main() {
	fmt.Println(parseDataSimple())
	fmt.Println(parseDataHard())
}

func parseDataSimple() int {
	splitted := strings.Split(input, "\n")

	sum := 0
	for _, line := range splitted {
		isValid := true

		linePart := strings.Split(line, ":")
		gameInfo := strings.Split(linePart[0], " ")
		id, _ := strconv.Atoi(gameInfo[1])

		games := strings.Split(linePart[1], ";")
		for _, game := range games {
			cubes := strings.Split(game, ",")
			for _, cube := range cubes {
				cubeInfo := strings.Split(cube, " ")
				number, _ := strconv.Atoi(cubeInfo[1])
				color := cubeInfo[2]

				switch color {
				case "red":
					if number > 12 {
						isValid = false
					}
				case "green":
					if number > 13 {
						isValid = false
					}
				case "blue":
					if number > 14 {
						isValid = false
					}
				}
			}
		}

		if isValid {
			sum += id
		}
	}
	return sum
}

func parseDataHard() int {
	splitted := strings.Split(input, "\n")

	sum := 0
	for _, line := range splitted {
		maxRedCubes := 0
		maxGreenCubes := 0
		maxBlueCubes := 0

		linePart := strings.Split(line, ":")

		games := strings.Split(linePart[1], ";")
		for _, game := range games {
			cubes := strings.Split(game, ",")
			for _, cube := range cubes {
				cubeInfo := strings.Split(cube, " ")
				number, _ := strconv.Atoi(cubeInfo[1])
				color := cubeInfo[2]

				switch color {
				case "red":
					if number > maxRedCubes {
						maxRedCubes = number
					}
				case "green":
					if number > maxGreenCubes {
						maxGreenCubes = number
					}
				case "blue":
					if number > maxBlueCubes {
						maxBlueCubes = number
					}
				}
			}
		}

		sum += (maxRedCubes * maxGreenCubes * maxBlueCubes)
	}
	return sum
}
