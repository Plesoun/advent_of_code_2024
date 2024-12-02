package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"main.go/day1"
	"main.go/day2"
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

	//day2
	readFileDay2, err := os.Open("./day2/day2_input.txt")
	if err != nil {
		panic(err)
	}

	day2FileScanner := bufio.NewScanner(readFileDay2)
	day2FileScanner.Split(bufio.ScanLines)
	day2InputArrayOfArrays := [][]int{}
	for day2FileScanner.Scan() {
		splitDay2 := strings.Fields(day2FileScanner.Text())
		lineArray := []int{}
		for n := range len(splitDay2) {
			value, err := strconv.Atoi(splitDay2[n])
			if err != nil {
				panic(err)
			}
			lineArray = append(lineArray, value)
		}
		day2InputArrayOfArrays = append(day2InputArrayOfArrays, lineArray)
	}

	defer readFileDay2.Close()

	fmt.Println("Safe measurement count is: ", day2.SafetyClassifier(day2InputArrayOfArrays))
}
