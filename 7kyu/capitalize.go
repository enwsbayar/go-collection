// Given a string of lowercase letters and an array of integer indices, capitalize all letters at the given indices. If an index is beyond the string, it should be ignored.

// Examples:
// "abcdef", [1,2,5]     ==> "aBCdeF"
// "abcdef", [1,2,5,100] ==> "aBCdeF" // There is no index 100.
// Good luck!

package main

import "strings"

func capitalize(s string, indices []int) string {
	runes := []rune(s)
	
	for _, idx := range indices {
		if idx >= 0 && idx < len(runes) {
			runes[idx] = []rune(strings.ToUpper(string(runes[idx])))[0]
		}
	}
	
	return string(runes)
}