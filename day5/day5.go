package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// struct of position defined as number and whether it is marked or not
type position struct {
	number int
	marked bool
}

// coordinate struct
type coordinate struct {
	x int
	y int
} 

// start and end coordinates struct as line defined by two coordinates
type line struct {
	start coordinate
	end coordinate
}

func main() {
	// create a map of coordinates to a count of how many times they have been visited
	coordinates := make(map[coordinate]int)

	// read the input file
	lines := readInput("input.txt")
	
	// for each line calculate the coordinates it covers, incrementing the count of each coordinate
	for _, line := range lines {
		// if the x coordinate is the same for both points, then we are dealing with a vertical line
		// if the y coordinate is the same for both points, then we are dealing with a horizontal line
		if line.start.x == line.end.x {
			// increment the count of each coordinate in the vertical line. if the end y is less than the start y then we need to increment the coordinates in reverse
			if line.end.y < line.start.y {
				for i := line.end.y; i <= line.start.y; i++ {
					coordinates[coordinate{line.start.x, i}]++
				}
			} else {
				for y := line.start.y; y <= line.end.y; y++ {
					coordinates[coordinate{line.start.x, y}]++
				}
			}
		} else if line.start.y == line.end.y {
			// increment the count of each coordinate in the horizontal line
			if line.end.x < line.start.x {
				for i := line.end.x; i <= line.start.x; i++ {
					coordinates[coordinate{i, line.start.y}]++
				}
			} else {
				for x := line.start.x; x <= line.end.x; x++ {
					coordinates[coordinate{x, line.start.y}]++
				}
			}
		} else {
			// if the line is not vertical or horizontal, then it is a diagonal line
			// calculate the slope of the line
			slope := float64(line.end.y - line.start.y) / float64(line.end.x - line.start.x)
			// if the slope is negative, then the line is going up
			if slope < 0 {
				// increment the count of each coordinate in the line

				// if the end x is less than the start x then we need to increment the coordinates in reverse
				if line.end.x < line.start.x {
					for i := line.end.x; i <= line.start.x; i++ {
						coordinates[coordinate{i, line.start.y + int(slope * float64(i - line.start.x))}]++
					}
				} else {
					for x := line.start.x; x <= line.end.x; x++ {
						coordinates[coordinate{x, line.start.y + int(slope * float64(x - line.start.x))}]++
					}
				}
			} else {
				// if the slope is positive, then the line is going down
				// increment the count of each coordinate in the line

				// if the end x is less than the start x then we need to increment the coordinates in reverse
				if line.end.x < line.start.x {
					for x := line.end.x; x <= line.start.x; x++ {
						coordinates[coordinate{x, line.start.y + int(slope * float64(x - line.start.x))}]++
					}
				} else {
					for x := line.start.x; x <= line.end.x; x++ {
						coordinates[coordinate{x, line.start.y + int(slope * float64(x - line.start.x))}]++
					}
				}
			}
		}
	}

	// print how many coordinates have at least 2 visits
	count := 0
	for _, value := range coordinates {
		if value >= 2 {
			count++
		}
	}

	// print the answer
	println(count)
}

// read the input file, returning a slice of lines
func readInput(filename string) []line {
	lines := make([]line, 0)
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read line as line text
		lineText := scanner.Text()
		// split the line into start and end coordinates, where start is the first coordinate and end is the second and the delimiter is " -> "
		lineSplit := strings.Split(lineText, " -> ")
		start := strings.Split(lineSplit[0], ",")
		end := strings.Split(lineSplit[1], ",")
		// convert the strings to ints
		startX, _ := strconv.Atoi(start[0])
		startY, _ := strconv.Atoi(start[1])
		endX, _ := strconv.Atoi(end[0])
		endY, _ := strconv.Atoi(end[1])
		// create a parsed line struct with the start and end coordinates
		parsedLine := line{coordinate{startX, startY}, coordinate{endX, endY}}
		// add the parsed line to the lines slice
		lines = append(lines, parsedLine)
	}
	return lines
}