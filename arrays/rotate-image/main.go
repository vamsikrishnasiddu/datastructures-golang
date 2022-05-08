package main

func rotate(matrix [][]int) {

	dummy := make([][]int, len(matrix))

	for i := 0; i < len(dummy); i++ {
		dummy[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix); j++ {

			dummy[j][len(matrix)-1-i] = matrix[i][j]
		}
	}

	copy(dummy, matrix)

}

func rotateOptimal(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < i; j++ {

			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix)/2; j++ {

			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}

	}

}

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	rotateOptimal(matrix)

}
