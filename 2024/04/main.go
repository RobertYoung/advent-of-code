package main

import (
	"fmt"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getLine(lines [][]string, index int) ([]string, bool) {
	if index >= 0 && index < len(lines) {
		return lines[index], true
	}
	return nil, false
}

func getCharacter(characters []string, index int) (string, bool) {
	if index >= 0 && index < len(characters) {
		return characters[index], true
	}
	return "", false
}

func getIndex3D(lines []string) [][]string {
	var result [][]string

	for _, line := range lines {
		characters := strings.Split(line, "")
		result = append(result, characters)
	}

	return result
}

func FindWord(lines []string, word string) int {
	if len(lines) == 0 || len(word) == 0 {
		return 0
	}

	wordCharacters := strings.Split(word, "")
	found := 0
	lines3d := getIndex3D(lines)

	for lineIndex, line := range lines3d {
		for charIndex, character := range line {
			if character == wordCharacters[0] {
				// check forward
				foundForward := true

				for wordIndex, wordCharacter := range wordCharacters {
					if nextCharacter, ok := getCharacter(line, charIndex+wordIndex); !ok || nextCharacter != wordCharacter {
						foundForward = false
						break
					}
				}

				if foundForward {
					found++
				}

				// check backward
				foundBackward := true

				for wordIndex, wordCharacter := range wordCharacters {
					if nextCharacter, ok := getCharacter(line, charIndex-wordIndex); !ok || nextCharacter != wordCharacter {
						foundBackward = false
						break
					}
				}

				if foundBackward {
					found++
				}

				// check up
				foundUp := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex-wordIndex)

					if !ok {
						foundUp = false
						break
					}

					if nextCharacter, ok := getCharacter(foundLine, charIndex); !ok || nextCharacter != wordCharacter {
						foundUp = false
						break
					}
				}

				if foundUp {
					found++
				}

				// check down
				foundDown := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex+wordIndex)
					if !ok {
						foundDown = false
						break
					}

					if nextCharacter, ok := getCharacter(foundLine, charIndex); !ok || nextCharacter != wordCharacter {
						foundDown = false
						break
					}
				}

				if foundDown {
					found++
				}

				// check diagonal up left
				foundDiagonalUpLeft := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex-wordIndex)

					if !ok {
						foundDiagonalUpLeft = false
						break
					}

					if nextCharacter, ok := getCharacter(foundLine, charIndex-wordIndex); !ok || nextCharacter != wordCharacter {
						foundDiagonalUpLeft = false
						break
					}
				}

				if foundDiagonalUpLeft {
					found++
				}

				// check diagonal up right
				foundDiagonalUpRight := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex-wordIndex)

					if !ok {
						foundDiagonalUpRight = false
						break
					}

					if nextCharacter, ok := getCharacter(foundLine, charIndex+wordIndex); !ok || nextCharacter != wordCharacter {
						foundDiagonalUpRight = false
						break
					}
				}

				if foundDiagonalUpRight {
					found++
				}

				// check diagonal down left
				foundDiagonalDownLeft := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex+wordIndex)

					if !ok {
						foundDiagonalDownLeft = false
						break
					}

					if nextCharacter, ok := getCharacter(foundLine, charIndex-wordIndex); !ok || nextCharacter != wordCharacter {
						foundDiagonalDownLeft = false
						break
					}
				}

				if foundDiagonalDownLeft {
					found++
				}

				// check diagonal down right
				foundDiagonalDownRight := true

				for wordIndex, wordCharacter := range wordCharacters {
					foundLine, ok := getLine(lines3d, lineIndex+wordIndex)
					if !ok {
						foundDiagonalDownRight = false
						break
					}
					if nextCharacter, ok := getCharacter(foundLine, charIndex+wordIndex); !ok || nextCharacter != wordCharacter {
						foundDiagonalDownRight = false
						break
					}
				}

				if foundDiagonalDownRight {
					found++
				}
			}
		}
	}

	return found
}

func FindXPattern(lines []string, word string) int {
	if len(lines) == 0 || len(word) != 3 {
		return 0
	}

	found := 0
	lines3d := getIndex3D(lines)
	wordCharacters := strings.Split(word, "")

	for lineIndex, line := range lines3d {
		for charIndex, character := range line {
			if character != wordCharacters[1] {
				continue
			}

			aboveLine, aboveLineOk := getLine(lines3d, lineIndex-1)
			belowLine, belowLineOk := getLine(lines3d, lineIndex+1)
			topLeftCharacter, topLeftCharacterOk := getCharacter(aboveLine, charIndex-1)
			topRightCharacter, topRightCharacterOk := getCharacter(aboveLine, charIndex+1)
			bottomLeftCharacter, bottomLeftCharacterOk := getCharacter(belowLine, charIndex-1)
			bottomRightCharacter, bottomRightCharacterOk := getCharacter(belowLine, charIndex+1)

			if !aboveLineOk || !belowLineOk || !topLeftCharacterOk || !topRightCharacterOk || !bottomLeftCharacterOk || !bottomRightCharacterOk {
				continue
			}

			word1 := strings.Join([]string{topLeftCharacter, character, bottomRightCharacter}, "")
			word2 := strings.Join([]string{topRightCharacter, character, bottomLeftCharacter}, "")

			if (word1 == word || reverseString(word1) == word) && (word2 == word || reverseString(word2) == word) {
				found++
			}
		}
	}

	return found
}

func main() {
	file, _ := util.ReadFileAsString("input.txt")
	lines := strings.Split(file, "\n")
	part1Count := FindWord(lines, "XMAS")
	part2Count := FindXPattern(lines, "MAS")

	fmt.Println("Result Part 1:", part1Count)
	fmt.Println("Result Part 2:", part2Count)
}
