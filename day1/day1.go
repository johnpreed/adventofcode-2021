package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	// load contents of file input.txt into an array of integers
	values := loadInput()

	// part 1
	println(getNumberOfIncreasingValues(values))

	// part 2
	println(getNumberOfIncreasingWindows(values))
}

func getNumberOfIncreasingValues(values []int) int {
	// init a counter
	var counter int

	// for each value in the array
	// if there is a previous value and the previous value is less than the current value then increment the counter
	for i := 0; i < len(values); i++ {
		if i > 0 && values[i-1] < values[i] {
			counter++
		}
	}

	return counter
}

func getNumberOfIncreasingWindows(values []int) int {
	// init a counter
	var counter int

	// for each value in the array
	// if there are two more values after the current value sum them and the current value
	for i := 0; i < len(values); i++ {
		if i+2 < len(values) {
			sum := values[i] + values[i+1] + values[i+2]

			// if the sum of the previous value, the current value, and the next value are less than the sum then increment the counter
			if i > 0 && values[i-1] + values[i] + values[i+1] < sum {
				counter++
			}
		}
	}

	return counter
}

func loadInput() []int {
	// open file input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file into a string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values []int
	for scanner.Scan() {
		values = append(values, parseInt(scanner.Text()))
	}
	return values
}

func parseInt(s string) int {
	// convert string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}