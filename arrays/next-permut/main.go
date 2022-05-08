package main

import "fmt"

func nextPermutation(nums []int) {

	var index1, index2 int = -1, -1

	// we need to find the break point index which
	// will break the increasing order.
	for i := len(nums) - 1; i > 0; i-- {

		if nums[i] > nums[i-1] {
			index1 = i - 1
			break
		}
	}

	// we need to find the index2 such that it is
	// greater than index1

	if index1 != -1 {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] > nums[index1] {
				index2 = i
				break
			}
		}

		// swap index1 and index2

		nums[index1], nums[index2] = nums[index2], nums[index1]

	}

	start := index1 + 1
	end := len(nums) - 1

	// we reverse the elements from index1+1 to last

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {

	var arr []int = make([]int, 5)

	arr = []int{3, 2, 1}

	nextPermutation(arr)

	fmt.Println(arr)

}
