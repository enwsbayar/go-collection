// In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

// For example:

// dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

// dup(["kelless","keenness"]) = ["keles","kenes"].

// Strings will be lowercase only, no spaces. See test cases for more examples.

// Good luck!

// If you like this Kata, please try:

// Alternate capitalization

// Vowel consonant lexicon

package main

func dup(arr []string) []string {
	result := make([]string, len(arr))

	for i, text := range arr {
		if len(text) == 0 {
			result[i] = text
			continue
		}

		filtered := make([]byte, 0, len(text))
		filtered = append(filtered, text[0])

		for j := 1; j < len(text); j++ {
			if text[j] != text[j-1] {
				filtered = append(filtered, text[j])
			}
		}

		result[i] = string(filtered)
	}

	return result
}