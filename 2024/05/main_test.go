package main

import (
	"testing"
)

var rules = []PageRule{
	{x: 47, y: 53},
	{x: 97, y: 13},
	{x: 97, y: 61},
	{x: 97, y: 47},
	{x: 75, y: 29},
	{x: 61, y: 13},
	{x: 75, y: 53},
	{x: 29, y: 13},
	{x: 97, y: 29},
	{x: 53, y: 29},
	{x: 61, y: 53},
	{x: 97, y: 53},
	{x: 61, y: 29},
	{x: 47, y: 13},
	{x: 75, y: 47},
	{x: 97, y: 75},
	{x: 47, y: 61},
	{x: 75, y: 61},
	{x: 47, y: 29},
	{x: 75, y: 13},
	{x: 53, y: 13},
}

func TestReorderPages(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		// {[]int{61, 13, 29}, []int{61, 29, 13}},
		// {[]int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}

	for _, test := range tests {
		result := ReorderPages(rules, test.input)

		for i, page := range result {
			if page != test.expected[i] {
				t.Errorf("ReorderPages(%v) = %v; want %v", test.input, result, test.expected)
			}
		}
	}
}

func TestIsValidPage(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{75, 47, 61, 53, 29}, true},
		{[]int{97, 61, 53, 29, 13}, true},
		{[]int{75, 29, 13}, true},
		{[]int{75, 97, 47, 61, 53}, false},
		{[]int{61, 13, 29}, false},
		{[]int{97, 13, 75, 29, 47}, false},
	}

	for _, test := range tests {
		result := IsValidPages(rules, test.input)
		if result != test.expected {
			t.Errorf("IsValidPage(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
