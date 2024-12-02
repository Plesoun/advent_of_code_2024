package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"main.go/day1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//day1
	dat, err := os.ReadFile("./day1/day1_input.txt")
	check(err)
	splitInput := strings.Fields(string(dat))

	arr1 := []int{}
	arr2 := []int{}
	for n := range len(splitInput) {
		value, err := strconv.Atoi(splitInput[n])
		if err != nil {
			panic(err)
		}
		if n%2 == 0 {
			arr1 = append(arr1, value)
		} else {
			arr2 = append(arr2, value)
		}
	}

	fmt.Println("Cumulative difference is: ", day1.CalculateCumulativeDifference(arr1, arr2))
	fmt.Println("Cumulative similarity score is: ", day1.CalculateSimilarityScore(arr1, arr2))
}
