package day2

import (
	"main.go/day1"
)

func SafetyClassifier(arr [][]int) int {
	var safeCounter int = 0
	for n := range arr {
		currentMeasurements := arr[n]
		var slope string = "unknown"

		for m := range currentMeasurements {
			if m > 0 {
				if currentMeasurements[m] == currentMeasurements[m-1] || day1.AbsoluteInt(currentMeasurements[m]-currentMeasurements[m-1]) > 3 {
					break
				}
				if currentMeasurements[m] < currentMeasurements[m-1] && (slope == "unknown" || slope == "decreasing") {
					slope = "decreasing"
				} else if currentMeasurements[m] > currentMeasurements[m-1] && (slope == "unknown" || slope == "increasing") {
					slope = "increasing"
				} else {
					break
				}
			}
			if m == len(currentMeasurements)-1 {
				safeCounter += 1
			}
		}

	}
	return safeCounter
}
