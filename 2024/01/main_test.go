package main

import (
	"testing"
)

func TestCalculateSimilarity(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	expected := 31
	result := CalculateSimilarity(left, right)
	if result != expected {
		t.Errorf("CalculateSimilarity failed, expected %d, got %d", expected, result)
	}
}

func TestCalculateDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	expected := 11
	result := CalculateDistance(left, right)
	if result != expected {
		t.Errorf("CalculateDistance failed, expected %d, got %d", expected, result)
	}
}
