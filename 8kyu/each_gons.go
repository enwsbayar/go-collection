// Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n, like so:

// each_cons([1,2,3,4], 2)
//   #=> [[1,2], [2,3], [3,4]]

// each_cons([1,2,3,4], 3)
//   #=> [[1,2,3],[2,3,4]]

// As you can see, the lists are cascading; ie, they overlap, but never out of order.


package main

func each_cons(list []int, n int) [][]int {
    if n <= 0 || n > len(list) {
        return [][]int{}
    }

    result := make([][]int, 0, len(list)-n+1)
    for i := 0; i <= len(list)-n; i++ {
        window := make([]int, n)
        copy(window, list[i:i+n])
        result = append(result, window)
    }

    return result
}

func EachCons(list []int, n int) [][]int {
    return each_cons(list, n)
}

