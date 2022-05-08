package main

import (
	"fmt"
	"time"
)

func setZeros(arr [][]int) {

	var brr [][]int = make([][]int, len(arr))

	for i := 0; i < len(brr); i++ {
		brr[i] = make([]int, len(arr[0]))
	}

	for k := 0; k < len(brr); k++ {

		for l := 0; l < len(brr[0]); l++ {
			brr[k][l] = arr[k][l]
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {

			if arr[i][j] == 0 {

				for m := 0; m < len(arr); m++ {
					brr[m][j] = 0
				}

				for n := 0; n < len(arr[0]); n++ {
					brr[i][n] = 0
				}
			}
		}
	}

	copy(arr, brr)

}

// the solution does not work if arr[i][j] is a negative number
func setZerosBruteWithNoSpace(arr [][]int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {

			if arr[i][j] == 0 {

				for m := 0; m < len(arr); m++ {
					if arr[m][j] != 0 {
						arr[m][j] = -1
					}
				}

				for n := 0; n < len(arr[0]); n++ {
					if arr[i][n] != 0 {
						arr[i][n] = -1
					}
				}
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {

			if arr[i][j] == -1 {
				arr[i][j] = 0
			}
		}
	}

}

func setZerosWithTwoArrays(arr [][]int) {
	// take two arrays
	cols := make([]int, len(arr[0]))

	rows := make([]int, len(arr))

	for i := 0; i < len(rows); i++ {
		rows[i] = -1
	}

	for i := 0; i < len(cols); i++ {
		cols[i] = -1
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 0 {
				rows[i] = 0
				cols[j] = 0
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {

			if cols[j] == 0 || rows[i] == 0 {
				arr[i][j] = 0
			}
		}
	}

}

func setZerosOptimal(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	isCol := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if isCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func setZerosOptimalst(matrix [][]int) {
	col0 := 1
	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {

				matrix[i][0] = 0
				matrix[0][j] = 0

			}

		}

	}

	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}

		}

		if col0 == 0 {
			matrix[i][0] = 0
		}

	}
}

func main() {
	var arr = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	t := time.Now()

	//setZeros(arr)
	//setZerosBruteWithNoSpace(arr)
	setZerosOptimalst(arr)
	//setZerosOptimal(arr)

	d := time.Since(t)

	fmt.Println(arr)

	fmt.Println(d)
}
