package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContent := string(readFile("input.txt"))
	fmt.Println(fileContent)

	lines := strings.Split(fileContent, "\n")
	fmt.Println(returnPoints(lines))
}

func readFile(inputPath string) []byte {
	dat, _ := os.ReadFile(inputPath)

	return dat
}

func returnPoints(lines []string) int {

	var totalSum int

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		var firstDigit rune
		var lastDigit rune

		for _, char := range line {
			if isDigit(char) {
				firstDigit = char
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])

			if isDigit(char) {
				lastDigit = char
				break
			}
		}
		number := int(firstDigit-'0')*10 + int(lastDigit-'0')
		totalSum += number

	}

	return totalSum

}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
