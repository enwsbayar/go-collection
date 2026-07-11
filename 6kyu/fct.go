// Given 
// u
// 0
// =
// 1
// ,
// u
// 1
// =
// 2
// u 
// 0
// ​
//  =1,u 
// 1
// ​
//  =2

// And the relation 
// 6
// u
// n
// u
// n
// +
// 1
// −
// 5
// u
// n
// u
// n
// +
// 2
// +
// u
// n
// +
// 1
// u
// n
// +
// 2
// =
// 0
// 6u 
// n
// ​
//  u 
// n+1
// ​
//  −5u 
// n
// ​
//  u 
// n+2
// ​
//  +u 
// n+1
// ​
//  u 
// n+2
// ​
//  =0

// Calculate 
// u
// n
// u 
// n
// ​
//   for any integer 
// n
// ≥
// 0
// n≥0.

// Examples:
// n
// =
// 17
// ,
// u
// 17
// =
// 131072
// n=17,u 
// 17
// ​
//  =131072

// n
// =
// 21
// ,
// u
// 21
// =
// 2097152
// n=21,u 
// 21
// ​
//  =2097152

// Remark:
// You can take two points of view to do this kata:


// the first one purely algorithmic from the definition of 
// u
// n
// u 
// n
// ​
 

// the second one - not at all mandatory, but as a complement - is to get a bit your head around and find which sequence is hidden behind 
// u
// n
// u 
// n
// ​
//  .


package main

// fct returns u_n for the sequence defined by u_0 = 1, u_1 = 2 and
// 6u_nu_{n+1} - 5u_nu_{n+2} + u_{n+1}u_{n+2} = 0.
// The hidden sequence is 2^n.
func fct(n int) int {
	return 1 << n
}