package main

import "fmt"

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	return countPathsDp(m, n, 0, 0, dp)

}

func countPaths(m int, n int, i int, j int) int {

	if i == m-1 && j == n-1 {
		return 1
	}

	if i >= m || j >= n {
		return 0
	}

	return countPaths(m, n, i+1, j) + countPaths(m, n, i, j+1)

}

func countPathsDp(m int, n int, i int, j int, dp [][]int) int {

	if i == m-1 && j == n-1 {
		return 1
	}

	if i >= m || j >= n {
		return 0
	}

	if dp[i][j] != -1 {
		return dp[i][j]
	}

	dp[i][j] = countPathsDp(m, n, i+1, j, dp) + countPathsDp(m, n, i, j+1, dp)

	return dp[i][j]

}

func uniquePathsOptimal(m, n int) int {

	N := m + n - 2

	r := m - 1

	var res int = 1

	for i := 1; i <= r; i++ {
		res = res * (N - r + i) / i
	}

	return res
}

func main() {

	fmt.Println(uniquePathsOptimal(2, 3))

}
