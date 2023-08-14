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

	highestCalories := 0
	currentElfTotalCalories := 0

	for scanner.Scan() {
		line := scanner.Text()
		currentCalories, err := strconv.Atoi(line)

		// An error in conversion indicates a blank line or a non-number line.
		if err != nil || line == "" { 
			if currentElfTotalCalories > highestCalories {
				highestCalories = currentElfTotalCalories
			}
			currentElfTotalCalories = 0 // Reset the sum for the next group.
		} else {
			currentElfTotalCalories += currentCalories
		}
	}

	// After reading the file, check the sum of the last group.
	if currentElfTotalCalories > highestCalories {
		highestCalories = currentElfTotalCalories
	}

	fmt.Println("highest value:", highestCalories)
}




