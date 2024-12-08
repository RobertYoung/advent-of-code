package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileAsArray(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return lines, nil
}

func ReadFileAsString(filename string) (string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()

	text := ""
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}

	text = text[:len(text)-1]

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return text, nil
}
