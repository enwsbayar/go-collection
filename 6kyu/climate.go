// dataand data1 are two strings with rainfall records of a few cities for months from January to December. The records of towns are separated by \n. The name of each town is followed by :.

// data and towns can be seen in "Your Test Cases:".

// Task:
// function: mean(town, strng) should return the average of rainfall for the city town and the strng data or data1 (In R and Julia this function is called avg).
// function: variance(town, strng) should return the variance of rainfall for the city town and the strng data or data1.
// Examples:
// mean("London", data), 51.19(9999999999996)
// variance("London", data), 57.42(833333333374)
// Notes:
// if functions mean or variance have as parameter town a city which has no records return -1 or -1.0 (depending on the language)

// Don't truncate or round: the tests will pass if abs(your_result - test_result) <= 1e-2 or abs((your_result - test_result) / test_result) <= 1e-6 depending on the language.

// Shell

// Shell tests only variance.
// In "function "variance" $1 is "data", $2 is "town".
// A ref: http://www.mathsisfun.com/data/standard-deviation.html

// data and data1 (can be named d0 and d1 depending on the language; see "Sample Tests:") are adapted from: http://www.worldclimate.com

package main

import (
	"math"
	"strconv"
	"strings"
)

// Mean returns the average rainfall for the given town in the data string.
// If the town is not found, returns -1.0
func Mean(town, strng string) float64 {
	vals := parseTownValues(town, strng)
	if len(vals) == 0 {
		return -1.0
	}
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

// Variance returns the population variance of rainfall for the given town.
// If the town is not found, returns -1.0
func Variance(town, strng string) float64 {
	vals := parseTownValues(town, strng)
	if len(vals) == 0 {
		return -1.0
	}
	mu := Mean(town, strng)
	sumsq := 0.0
	for _, v := range vals {
		d := v - mu
		sumsq += d * d
	}
	return sumsq / float64(len(vals))
}

// parseTownValues extracts numerical rainfall values for a town from the data string.
func parseTownValues(town, strng string) []float64 {
	if town == "" || strng == "" {
		return nil
	}
	lines := strings.Split(strng, "\n")
	prefix := town + ":"
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			// after the colon, entries are like "Jan 81.2,Feb 63.2,..." or similar
			rest := strings.TrimSpace(line[len(prefix):])
			if rest == "" {
				return nil
			}
			parts := strings.Split(rest, ",")
			vals := make([]float64, 0, len(parts))
			for _, p := range parts {
				p = strings.TrimSpace(p)
				// split by space to get the numeric part at the end
				fields := strings.Fields(p)
				if len(fields) < 2 {
					continue
				}
				numStr := fields[len(fields)-1]
				if numStr == "" {
					continue
				}
				f, err := strconv.ParseFloat(numStr, 64)
				if err != nil {
					continue
				}
				if math.IsNaN(f) || math.IsInf(f, 0) {
					continue
				}
				vals = append(vals, f)
			}
			return vals
		}
	}
	return nil
}

