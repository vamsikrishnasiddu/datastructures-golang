package main

import "fmt"

func findDuplicateExtraSpace(nums []int) int {

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			return nums[i]
		}
	}

	return 0

}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	slow = nums[slow]
	fast = nums[nums[fast]]

	for slow != fast {

		slow = nums[slow]
		fast = nums[nums[fast]]

	}

	fast = nums[0]

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func main() {

	var arr = []int{1, 3, 4, 2, 2}

	fmt.Println(findDuplicate(arr))

}
