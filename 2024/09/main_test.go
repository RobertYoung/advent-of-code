package main

import (
	"strings"
	"testing"
)

func TestConvertDiskMap(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2333133121414131402", "00...111...2...333.44.5555.6666.777.888899"},
	}

	for _, test := range tests {
		result, _ := ConvertDiskMap(test.input)
		expected := strings.Split(test.expected, "")

		for i, block := range result {
			if block.value != expected[i] {
				t.Errorf("TestConvertDiskMap(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestMoveBlocksPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2333133121414131402", "0099811188827773336446555566.............."},
	}

	for _, test := range tests {
		input, _ := ConvertDiskMap(test.input)
		result, _ := MoveBlocksPart1(input)
		expected := strings.Split(test.expected, "")

		for i, block := range result {
			if block.value != expected[i] {
				t.Errorf("TestMoveBlocksPart1(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestCalculateChecksumPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"2333133121414131402", 1928},
	}

	for _, test := range tests {
		input, _ := ConvertDiskMap(test.input)
		blocks, _ := MoveBlocksPart1(input)
		result, _ := CalculateChecksum(blocks)

		if result != test.expected {
			t.Errorf("TestCalculateChecksum(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestMoveBlocksPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2333133121414131402", "00992111777.44.333....5555.6666.....8888.."},
	}

	for _, test := range tests {
		input, _ := ConvertDiskMap(test.input)
		result, _ := MoveBlocksPart2(input)
		expected := strings.Split(test.expected, "")

		for i, block := range result {
			if block.value != expected[i] {
				t.Errorf("TestMoveBlocksPart2(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestCalculateChecksumPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"2333133121414131402", 2858},
	}

	for _, test := range tests {
		input, _ := ConvertDiskMap(test.input)
		blocks, _ := MoveBlocksPart2(input)
		result, _ := CalculateChecksum(blocks)

		if result != test.expected {
			t.Errorf("TestCalculateChecksum2(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
