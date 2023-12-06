package day1

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	regexFirstSimple  = "([0-9]).*"
	regexSecondSimple = ".*([0-9])"
	regexFirstHard    = "(one|two|three|four|five|six|seven|eight|nine|[0-9]).*"
	regexSecondHard   = ".*(one|two|three|four|five|six|seven|eight|nine|[0-9])"
)

func ParseDataSimple(data string) int {
	splitted := strings.Split(data, "\n")

	sum := 0
	for _, line := range splitted {
		firstNumber, lastNumber := findNumbers(line, regexFirstSimple, regexSecondSimple)
		if lastNumber == "" {
			lastNumber = firstNumber
		}
		number, _ := strconv.Atoi(firstNumber + lastNumber)
		sum += number
	}

	return sum
}

func ParseDataHard(data string) int {
	splitted := strings.Split(data, "\n")

	sum := 0
	for _, line := range splitted {
		firstNumber, lastNumber := findNumbers(line, regexFirstHard, regexSecondHard)
		if lastNumber == "" {
			lastNumber = firstNumber
		}
		number, _ := strconv.Atoi(firstNumber + lastNumber)
		sum += number
	}

	return sum
}

func findNumbers(line string, regex1 string, regex2 string) (string, string) {
	firstNumber := ""
	lastNumber := ""

	r1, _ := regexp.Compile(regex1)
	subMatch := r1.FindStringSubmatch(line)
	if len(subMatch) > 1 {
		firstNumber = subMatch[1]
	}

	r2, _ := regexp.Compile(regex2)
	subMatch = r2.FindStringSubmatch(line)
	if len(subMatch) > 1 {
		lastNumber = subMatch[1]
	}

	firstNumber = convertStringToNumber(firstNumber)
	lastNumber = convertStringToNumber(lastNumber)

	return firstNumber, lastNumber
}

func convertStringToNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}
