package day2

import "testing"

func TestSafetyClassifier(t *testing.T) {
	arr := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	result := SafetyClassifier(arr)
	if result != 2 {
		t.Fatalf("Expected result 2 got result %v", result)
	}
}
