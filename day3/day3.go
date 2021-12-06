package main

import (
	"bufio"
	"os"
)

func main() {
	// read input file into array of strings
	lines := readFile("input.txt")

	// calculate power consumption
	powerConsumption := calculatePowerConsumption(lines)

	// calculate life support rating
	lifeSupportRating := calculateLifeSupportRating(lines)

	// print power consumption
	println(powerConsumption)

	// print life support rating
	println(lifeSupportRating)
}

// calculate power consumption
func calculatePowerConsumption(lines []string) int {
	// get the length of the first line
	lineLength := len(lines[0])

	// fore each index from 0 to line length, get the most common character at that index and append to a new string (gamma)
	var gamma string
	for i := 0; i < lineLength; i++ {
		// get the most common bit at that index
		bit := getMostCommonBit(lines, i)
		if bit == 0 {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}

	// convert string to a bit set (gamma set)
	var gammaSet []int
	for _, char := range gamma {
		gammaSet = append(gammaSet, int(char-'0'))
	}

	// invert bit set into a new bit set (epsilon set)
	var epsilonSet []int
	for _, bit := range gammaSet {
		if bit == 0 {
			epsilonSet = append(epsilonSet, 1)
		} else {
			epsilonSet = append(epsilonSet, 0)
		}
	}

	// convert bit set to decimal (gamma rate)
	gammaRate := bitSetToDec(gammaSet)

	// convert bit set to decimal (epsilon rate)
	epsilonRate := bitSetToDec(epsilonSet)

	// multiply gamma rate epsilon rate as (powerConsumption)
	return gammaRate * epsilonRate
}

// calculate life support rating
func calculateLifeSupportRating(lines []string) int {
	// get the length of the first line
	lineLength := len(lines[0])

	// copy lines to new slice
	var filteredLines []string
	filteredLines = append(filteredLines, lines...)

	// starting at index 0, filter the lines until only one line remains or index reaches line length, by removing all lines that do not have the most common bit at the current index, if bit is -1, keep values with a 1 in the position being considered
	for i := 0; i < lineLength; i++ {
		// get the most common bit at that index
		bit := getMostCommonBit(filteredLines, i)

		// if bit is 1 or -1, keep values with a 1 in the position being considered
		if bit == 1 || bit == -1 {
			filteredLines = filterLines(filteredLines, i, '1')
		} else {
			// if bit is 0, keep values with a 0 in the position being considered
			filteredLines = filterLines(filteredLines, i, '0')
		}

		// if there is only one line left, break
		if len(filteredLines) == 1 {
			break
		}
	}

	// for the remaining line, convert the line to a bit set and set the decimal value (oxygenGeneratorRating)
	var oxygenGeneratorRatingSet []int
	for _, char := range filteredLines[0] {
		oxygenGeneratorRatingSet = append(oxygenGeneratorRatingSet, int(char-'0'))
	}

	// reset filtered lines, copy lines to filtered lines
	filteredLines = []string{}
	filteredLines = append(filteredLines, lines...)

	// starting at index 0, filter the lines until only one line remains or index reaches line length, by removing all lines that have the most common bit at the current index, if bit is -1, keep values with a 1 in the position being considered
	for i := 0; i < lineLength; i++ {
		// get the most common bit at that index
		bit := getMostCommonBit(filteredLines, i)

		// invert bit
		if bit == 0 {
			bit = 1
		} else {
			bit = 0
		}

		// if bit is 1 or -1, keep values with a 1 in the position being considered
		if bit == 1 || bit == -1 {
			filteredLines = filterLines(filteredLines, i, '1')
		} else {
			// if bit is 0, keep values with a 0 in the position being considered
			filteredLines = filterLines(filteredLines, i, '0')
		}

		// if there is only one line left, break
		if len(filteredLines) == 1 {
			break
		}
	}

	// for the remaining line, convert the line to a bit set and set the decimal value (carbonScrubberRating)
	var carbonScrubberRatingSet []int
	for _, char := range filteredLines[0] {
		carbonScrubberRatingSet = append(carbonScrubberRatingSet, int(char-'0'))
	}

	// convert bit set to decimal (oxygenGeneratorRating)
	oxygenGeneratorRating := bitSetToDec(oxygenGeneratorRatingSet)

	// convert bit set to decimal (carbonScrubberRating)
	carbonScrubberRating := bitSetToDec(carbonScrubberRatingSet)

	// multiply oxygenGeneratorRating carbonScrubberRating and return
	return oxygenGeneratorRating * carbonScrubberRating
}

// filter lines at index by bit
func filterLines(lines []string, index int, bit byte) []string {
	var filteredLines []string
	for _, line := range lines {
		// if line at index matches bit, add to filtered lines
		if line[index] == bit {
			filteredLines = append(filteredLines, line)
		}
	}

	return filteredLines
}


// convert bit set to decimal
func bitSetToDec(bitSet []int) int {
	var dec int
	for _, bit := range bitSet {
		dec = dec*2 + bit
	}
	return dec
}

// get most common bit at index for a slice of strings. if there are equal or more 1s, return 1, if there are more 0's return 0, else return -1
func getMostCommonBit(lines []string, index int) int {
	var count1 int
	var count0 int

	for _, line := range lines {
		if line[index] == '1' {
			count1++
		} else {
			count0++
		}
	}

	if count1 > count0 {
		return 1
	} else if count0 > count1 {
		return 0
	} else {
		return -1
	}
}


func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}