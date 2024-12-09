package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RobertYoung/advent-of-code/util"
)

type Block struct {
	id    int
	value string
}

func ConvertDiskMap(input string) ([]Block, error) {
	id := 0
	result := []Block{}

	for i, char := range strings.Split(input, "") {
		if (i % 2) == 0 {
			num, _ := strconv.Atoi(char)

			for j := 0; j < num; j++ {
				result = append(result, Block{
					id:    id,
					value: strconv.Itoa(id),
				})
			}

			id++
		} else {
			num, _ := strconv.Atoi(char)

			for j := 0; j < num; j++ {
				result = append(result, Block{
					id:    -1,
					value: ".",
				})
			}
		}
	}

	return result, nil
}

func MoveBlocksPart1(input []Block) ([]Block, error) {
	result := input

	for i := len(result) - 1; i >= 0; i-- {
		done := false

		if result[i].value != "." {
			for j := 0; j < len(result); j++ {
				if result[j].value == "." {
					if j >= i {
						done = true
						break
					}

					result[j], result[i] = result[i], result[j]

					break
				}
			}
		}

		if done {
			break
		}
	}

	return result, nil
}

func MoveBlocksPart2(input []Block) ([]Block, error) {
	result := input

	for i := len(result) - 1; i >= 0; i-- {
		done := false

		if result[i].value != "." {
			fileSize := 1

			for fileSizeIndex := i; fileSizeIndex > 0; fileSizeIndex-- {
				if result[fileSizeIndex].id == result[fileSizeIndex-1].id {
					fileSize++
				} else {
					break
				}
			}

			freeSpaceSize := 0

			for j := 0; j < len(result); j++ {
				if result[j].value == "." {
					freeSpaceSize++

					if j >= i {
						i -= (fileSize - 1)
						break
					}

					if freeSpaceSize >= fileSize {
						for k := 0; k < fileSize; k++ {
							result[j-k], result[i-k] = result[i-k], result[j-k]
						}

						break
					}
				} else {
					freeSpaceSize = 0
				}
			}
		}

		if done {
			break
		}
	}

	return result, nil
}

func CalculateChecksum(blocks []Block) (int, error) {
	result := 0
	for index, block := range blocks {
		if block.value == "." {
			continue
		}
		result += index * block.id
	}

	return result, nil
}

func main() {
	diskmap, _ := util.ReadFileAsString("input.txt")
	part1Converted, _ := ConvertDiskMap(diskmap)
	part1Moved, _ := MoveBlocksPart1(part1Converted)
	part1Checksum, _ := CalculateChecksum(part1Moved)

	fmt.Println("Part 1", part1Checksum)

	part2Converted, _ := ConvertDiskMap(diskmap)
	part2Moved, _ := MoveBlocksPart2(part2Converted)
	part2Checksum, _ := CalculateChecksum(part2Moved)

	fmt.Println("Part 2", part2Checksum)
}
