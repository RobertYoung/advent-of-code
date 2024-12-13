package main

import (
	"testing"
)

func TestConvertInput(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`, 4},
	}

	for _, test := range tests {
		result, _ := ConvertInput(test.input, 0)

		if len(result) != test.expected {
			t.Errorf("TestConvertInput() = %v; want %v", len(result), test.expected)
		}
	}
}

func TestCalculateTokens(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`, 280},
		{`Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`, -1},
		{`Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450`, 200},
		{`Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`, -1},
	}

	for _, test := range tests {
		machines, _ := ConvertInput(test.input, 0)
		result, _ := machines[0].calculateTokens()

		if result != test.expected {
			t.Errorf("TestCalculateTokens() = %v; want %v", result, test.expected)
		}
	}
}
