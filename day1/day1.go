package day1

import (
	"fmt"
	"slices"
)

func AbsoluteInt(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func CalculateCumulativeDifference(arr1, arr2 []int) int {
	slices.Sort(arr1)
	slices.Sort(arr2)
	var cumulative int = 0
	fmt.Println("Got inputs ", arr1, arr2)
	for n := range len(arr1) {
		difference := arr1[n] - arr2[n]
		cumulative += AbsoluteInt(difference)
	}
	return cumulative
}
