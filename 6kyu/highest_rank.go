package main

// Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.

// Note: no empty arrays will be given.

// Examples
// [12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
// [12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
// [12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3

func highestRank(numbers []int) int {
    counts := make(map[int]int, len(numbers))
    best := numbers[0]
    maxCount := 0

    for _, n := range numbers {
        counts[n]++
        count := counts[n]

        if count > maxCount || (count == maxCount && n > best) {
            best = n
            maxCount = count
        }
    }

    return best
}
