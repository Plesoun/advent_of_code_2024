package day1

import (
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
	for n := range len(arr1) {
		difference := arr1[n] - arr2[n]
		cumulative += AbsoluteInt(difference)
	}
	return cumulative
}

func CalculateSimilarityScore(arr1, arr2 []int) int {
	slices.Sort(arr1)
	slices.Sort(arr2)

	similarityMap := make(map[int]int)
	similarityScore := 0

	for n := range len(arr1) {
		currentArr1 := arr1[n]

		if _, ok := similarityMap[currentArr1]; !ok {
			similarityMap[currentArr1] = 0
		}
	}
	for m := range len(arr2) {
		currentArr2 := arr2[m]
		if _, ok := similarityMap[currentArr2]; ok {
			similarityMap[currentArr2] += 1
		}
	}
	for key, value := range similarityMap {
		similarityScore += (key * value)
	}

	return similarityScore
}
