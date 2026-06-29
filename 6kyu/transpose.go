// Write a function that outputs the transpose of a matrix - a new matrix where the columns and rows of the original are swapped.

// For example, the transpose of:

// | 1 2 3 |
// | 4 5 6 |
// is

// | 1 4 |
// | 2 5 |
// | 3 6 |
// The input to your function will be an array of matrix rows. You can assume that each row has the same length, and that the height and width of the matrix are both positive.

package main

func Transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for col := 0; col < cols; col++ {
		transposed[col] = make([]int, rows)
		for row := 0; row < rows; row++ {
			transposed[col][row] = matrix[row][col]
		}
	}

	return transposed
}