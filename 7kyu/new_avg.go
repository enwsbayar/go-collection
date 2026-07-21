// The accounts of the "Fat to Fit Club (FFC)" association are supervised by John as a volunteered accountant. The association is funded through financial donations from generous benefactors. John has a list of the first n donations: [14, 30, 5, 7, 9, 11, 15] He wants to know how much the next benefactor should give to the association so that the average of the first n + 1 donations should reach an average of 30. After doing the math he found 149. He thinks that he could have made a mistake.

// Could you help him?

// Task
// The function should return the expected donation (rounded up to the next integer) that will permit to reach the average.

// Should the last donation be a non positive number (<= 0) John wants us either to throw an error or return such value so he will clearly see that his expectations are not great enough.

// In "Sample Tests" you can see what to return.

// Notes:
// all donations and new average are numbers (integers or floats), array of donations can be empty.
// See examples below and "Sample Tests" to see which return is to be done.
// array: [14, 30, 5, 7, 9, 11, 15], new average: 92 -> result: 645
// array: [14, 30, 5, 7, 9, 11, 15], new average: 2  -> result: variable, see "Sample Tests" 
package main

import "math"

func NewAvg(arr []float64, navg float64) int {
	if len(arr) == 0 {
		return int(math.Ceil(navg))
	}

	sum := 0.0
	for _, v := range arr {
		sum += v
	}

	required := (navg * float64(len(arr)+1)) - sum
	if required <= 0 {
		return -1
	}

	return int(math.Ceil(required))
} 
