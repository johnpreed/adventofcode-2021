package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


func main() {
	// create a fish map of int to int, representing the number of lanternfish that spawn in a given day, where the key is the day and the value is the number of lanternfish that spawn in that day
	fishMap := readInput()

	// iterate over 256 days
	for i := 1; i <= 256; i++ {
		// set a counter for new fish
		newFish := 0

		// for each key from 0 to 8
		for key := 0; key < 9; key++ {
			// if the key is 0, set the new fish to the value of the key
			if key == 0 {
				// increment newF fish by the value at index 0
				newFish += fishMap[key]
				// set the value at index 0 to 0
				fishMap[0] = 0
			} else {
				// increment the key at the index of key - 1 by the value at index key and set the value at index key to 0
				fishMap[key - 1] += fishMap[key]
				fishMap[key] = 0
			}
		}

		// increment the value at index 8 by the new fish
		fishMap[8] += newFish
		// increment the value at index 6 by the new fish
		fishMap[6] += newFish
	}

	// sum the values in the fish map
	sum := 0
	for _, value := range fishMap {
		sum += value
	}
	
	// print the sum
	println(sum)
}

// read input file "input.txt" to populate the map adding one to each day that a lanternfish spawns
func readInput() map[int]int {
	// open the file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// create a scanner to read the file
	scanner := bufio.NewScanner(file)
	// create a map of int to int
	fishMap := make(map[int]int)
	// loop through the file
	for scanner.Scan() {
		// split the line by comma into array of integers
		line := strings.Split(scanner.Text(), ",")

		// loop through the array of integers
		for _, value := range line {
			// convert the string to an int
			fish, _ := strconv.Atoi(value)
			// increment the value at index fish by 1
			fishMap[fish]++
		}
	}
	// close the file
	file.Close()
	// return the map
	return fishMap
}