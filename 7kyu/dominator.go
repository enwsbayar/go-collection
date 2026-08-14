// A zero-indexed array arr consisting of n integers is given. The dominator of array arr is the value that occurs in more than half of the elements of arr.
// For example, consider array arr such that arr = [3,4,3,2,3,1,3,3]
// The dominator of arr is 3 because it occurs in 5 out of 8 elements of arr and 5 is more than a half of 8.
// Write a function dominator(arr) that, given a zero-indexed array arr consisting of n integers, returns the dominator of arr. The function should return −1 if array does not have a dominator. All values in arr will be >=0.

package main
func Dominator(arr []int) int {
	var candidate int
	count := 0

	// Find candidate
	for _, num := range arr {
		if count == 0 {
			candidate = num
			count = 1
		} else if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Verify candidate
	occurrences := 0
	for _, num := range arr {
		if num == candidate {
			occurrences++
		}
	}

	if occurrences > len(arr)/2 {
		return candidate
	}

	return -1
}