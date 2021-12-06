package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentX := 0
	currentY := 0
	aim := 0

	// open file input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file into a string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	// parse each line into a string and int
	for scanner.Scan() {
		line := scanner.Text()
		
		// split line into a slice of strings
		lineSlice := strings.Split(line, " ")
		
		// if line is 2 elements, then it is a move command
		if len(lineSlice) == 2 {
			direction := lineSlice[0]
			magnitude, _ := strconv.Atoi(lineSlice[1])

			x, y, a := getVertices(direction, magnitude, aim)
			currentX += x
			currentY += y
			aim += a
		}
	}

	// print final position
	println(currentX, currentY)
	// print the multiplier of the final position
	println(currentX * currentY)
}

func getVertices(direction string, magnitude int, aim int) (int, int, int) {
	if direction == "forward" {
		return magnitude, magnitude * aim, 0
	} else if direction == "up" {
		return 0, 0, -1 * magnitude
	} else if direction == "down" {
		return 0, 0, magnitude
	} else {
		return 0, 0, 0
	}
}