package main

import (
	"reflect"
	"testing"
)

func compareEquations(eq1, eq2 Equation) bool {
	if eq1.total != eq2.total {
		return false
	}
	return reflect.DeepEqual(eq1.values, eq2.values)
}

func TestConvertEquation(t *testing.T) {
	tests := []struct {
		input    string
		expected Equation
	}{
		{"190: 10 19", Equation{total: 190, values: []int{10, 19}}},
		{"3267: 81 40 27", Equation{total: 3267, values: []int{81, 40, 27}}},
	}

	for _, test := range tests {
		result, _ := ConvertEquation(test.input)

		if !compareEquations(result, test.expected) {
			t.Errorf("ConvertEquation(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindValidCalibrationsPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"190: 10 19", true},
		{"3267: 81 40 27", true},
		{"83: 17 5", false},
		{"156: 15 6", false},
		{"7290: 6 8 6 15", false},
		{"161011: 16 10 13", false},
		{"192: 17 8 14", false},
		{"21037: 9 7 18 13", false},
		{"292: 11 6 16 20", true},
	}

	for _, test := range tests {
		calibration, _ := ConvertEquation(test.input)
		result := calibration.IsValidPart1()

		if result != test.expected {
			t.Errorf("TestFindValidCalibrationsPart1(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindValidCalibrationsPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"190: 10 19", true},
		{"3267: 81 40 27", true},
		{"83: 17 5", false},
		{"156: 15 6", true},
		{"7290: 6 8 6 15", true},
		{"161011: 16 10 13", false},
		{"192: 17 8 14", true},
		{"21037: 9 7 18 13", false},
		{"292: 11 6 16 20", true},
	}

	for _, test := range tests {
		calibration, _ := ConvertEquation(test.input)
		result := calibration.IsValidPart2()

		if result != test.expected {
			t.Errorf("TestFindValidCalibrationsPart2(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
