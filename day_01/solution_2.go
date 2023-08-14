package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input file
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var largest, secondLargest, thirdLargest int
	currentElfTotalCalories := 0

	for scanner.Scan() {
		line := scanner.Text()
		currentCalories, err := strconv.Atoi(line)

		// An error in conversion indicates a blank line or a non-number line.
		if err != nil || line == "" { // Added line == "" just in case there's a blank line without an error.
			if currentElfTotalCalories > largest {
				thirdLargest = secondLargest
				secondLargest = largest
				largest = currentElfTotalCalories
			} else if currentElfTotalCalories > secondLargest {
				thirdLargest = secondLargest
				secondLargest = currentElfTotalCalories
			} else if currentElfTotalCalories > thirdLargest {
				thirdLargest = currentElfTotalCalories
			}

			currentElfTotalCalories = 0 // Reset the sum for the next group.
		} else {
			currentElfTotalCalories += currentCalories
		}
	}

	// Check for the last group after finishing reading the file
	if currentElfTotalCalories > largest {
		thirdLargest = secondLargest
		secondLargest = largest
		largest = currentElfTotalCalories
	} else if currentElfTotalCalories > secondLargest {
		thirdLargest = secondLargest
		secondLargest = currentElfTotalCalories
	} else if currentElfTotalCalories > thirdLargest {
		thirdLargest = currentElfTotalCalories
	}

	// Sum the three largest values
	total := largest + secondLargest + thirdLargest

	fmt.Println(total)
}