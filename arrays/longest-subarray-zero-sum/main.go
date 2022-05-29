package main

import "fmt"

func maxSubArrayZeroSum(nums []int) int {
	count := 0
	max_count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]

			if sum == 0 {
				count = j - i + 1
			}

			if count > max_count {
				max_count = count
			}
		}
	}

	return max_count

}

func maxSubArrayZeroSumOn(nums []int) int {

	max := 0
	sum := 0

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]

		if sum == 0 {
			max = i + 1
		} else {

			if _, ok := m[sum]; ok {
				if i-m[sum] > max {
					max = i - m[sum]
				}
			} else {
				m[sum] = i
			}

		}

	}

	return max

}

func main() {

	var arr []int = []int{1, -1, 3, 2, -2, -8, 1, 7, 10, 23}

	fmt.Println(maxSubArrayZeroSumOn(arr))

}
