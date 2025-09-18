package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	max := map[string]int16{
		"red": 12, "green": 13, "blue": 14,
	}
	var sum int
	lines := readFile("input.txt")

	for i, line := range lines {
		if line == "" {
			continue
		}

		gameID := i + 1
		colonIndex := strings.Index(line, ":")
		if colonIndex == -1 {
			continue
		}
		lineBody := line[colonIndex+1:]

		if isGamePossible(lineBody, max) {
			sum += gameID
		}
	}
	fmt.Println(sum)
}

func isGamePossible(game string, max map[string]int16) bool {
	remainingRounds := game
	for len(remainingRounds) > 0 {
		semicolonIndex := strings.Index(remainingRounds, ";")
		var currentRound string
		if semicolonIndex == -1 {
			currentRound = remainingRounds
			remainingRounds = ""
		} else {
			currentRound = remainingRounds[:semicolonIndex]
			remainingRounds = remainingRounds[semicolonIndex+1:]
		}

		remainingDices := currentRound
		for len(remainingDices) > 0 {
			commaIndex := strings.Index(remainingDices, ",")
			var currentDice string
			if commaIndex == -1 {
				currentDice = remainingDices
				remainingDices = ""
			} else {
				currentDice = remainingDices[:commaIndex]
				remainingDices = remainingDices[commaIndex+1:]
			}

			currentDice = strings.TrimSpace(currentDice)
			spaceIndex := strings.Index(currentDice, " ")
			if spaceIndex == -1 {
				continue
			}

			quantityStr := currentDice[:spaceIndex]
			color := currentDice[spaceIndex+1:]

			quantity, err := strconv.Atoi(quantityStr)
			if err != nil {
				return false
			}

			maxVal, ok := max[color]
			if !ok || int16(quantity) > maxVal {
				return false
			}
		}
	}
	return true
}

func readFile(filePath string) []string {
	fileContent, _ := os.ReadFile(filePath)
	return strings.Split(string(fileContent), "\n")
}
