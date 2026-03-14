// Create a function add(n)/Add(n) which returns a function that always adds n to any number

// Note for Java: the return type and methods have not been provided to make it a bit more challenging.

// var addOne = Add(1)
// addOne(3) // 4

package main

func Add(n int) func(int) int {
	return func(x int) int {
		return n + x
	}
}

func add(n int) func(int) int {
	return Add(n)
}