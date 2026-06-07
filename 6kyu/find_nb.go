// Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of
// n^3, the cube above will have volume of (n-1)^3 and so on until the top which will have a volume of 1^3.
// You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?
// The function should return n such that n^3 + (n-1)^3 + ... + 1^3 = m, or -1 if no such n exists.

package main

func findNb(m int) int {
    if m <= 0 {
        return -1
    }
    target := int64(m)
    sum := int64(0)
    n := 1

    for sum < target {
        sum += int64(n) * int64(n) * int64(n)
        if sum == target {
            return n
        }
        n++
    }

    return -1
}
