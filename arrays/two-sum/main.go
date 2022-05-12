package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			return []int{i, m[target-nums[i]]}
		}
	}

	return nil

}

func main() {

	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))

}
