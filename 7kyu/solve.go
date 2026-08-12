// Remove the duplicates from a list of integers, keeping the last ( rightmost ) occurrence of each element.
// Example:
// For input: [3, 4, 4, 3, 6, 3]
// remove the 3 at index 0
// remove the 4 at index 1
// remove the 3 at index 3
// Expected output: [4, 6, 3]
// More examples can be found in the test cases.
// Good luck!

func solve(arr []int) []int {
	result := []int{}

	for i := 0; i < len(arr); i++ {
		found := false

		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				found = true
				break
			}
		}

	
		if !found {
			result = append(result, arr[i])
		}
	}

	return result
}