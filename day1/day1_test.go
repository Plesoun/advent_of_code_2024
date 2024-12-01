package day1

import (
	"testing"
)

func TestAbsoluteInt(t *testing.T) {
	intpoz := 10
	intneg := -4
	result1 := AbsoluteInt(intpoz)
	result2 := AbsoluteInt(intneg)
	if result1 != 10 {
		t.Fatalf("Wanted 10 got %v", result1)
	}
	if result2 != 4 {
		t.Fatalf("Wanted 4 got %v", result2)
	}
}

func TestCalculateCumulativeDifference(t *testing.T) {
	arr1 := []int{2, 1, 4, 8}
	arr2 := []int{3, 5, 7, 9}
	result := CalculateCumulativeDifference(arr1, arr2)

	if result != 9 {
		t.Fatalf("Wanted 9 got %v", result)
	}
}
