package main

import "fmt"

func missingRepeating(nums []int) (int, int) {
	missNum := -1
	repeatNum := -1

	freq := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] == 0 {
			missNum = i
		} else if freq[i] == 2 {
			repeatNum = i
		}
	}

	return missNum, repeatNum

}

func missingRepeatingOptimal(nums []int) (int, int) {
	var xor1 int
	var set_bit_no int
	var i int
	var x, y int
	var n int = len(nums)

	xor1 = nums[0]

	// Get the xor of all the array elements

	for i = 1; i < len(nums); i++ {
		xor1 = xor1 ^ nums[i]
	}

	// xor the previous result with numbers from (1 to n)

	for i = 1; i <= n; i++ {
		xor1 = xor1 ^ i
	}

	//Get the rightmost set bit in set_bit_no

	set_bit_no = xor1 & ^(xor1 - 1)

	/* Now divide elements into two sets by comparing a rightmost set bit
	   of xor1 with the bit at the same position in each element.
	   Also, get XORs of two sets. The two XORs are the output elements.
	   The following two for loops serve the purpose */

	for i = 0; i < n; i++ {
		if nums[i]&set_bit_no == 1 {

			/* arr[i] belongs to first set */
			x = x ^ nums[i]

		} else {
			/* arr[i] belongs to second set */
			y = y ^ nums[i]

		}
	}

	for i = 1; i <= n; i++ {
		if i&set_bit_no == 1 {

			/* i belongs to first set */
			x = x ^ i

		} else {

			/* i belongs to second set */
			y = y ^ i

		}

	}

	// NB! numbers can be swapped, maybe do a check ?
	x_count := 0
	for i = 0; i < n; i++ {
		if nums[i] == x {

			x_count++

		}

	}

	if x_count == 0 {
		return x, y
	}

	return y, x

}

func main() {

	var arr = []int{4, 3, 6, 2, 1, 1}

	miss, repeat := missingRepeatingOptimal(arr)

	fmt.Println(miss, repeat)

}
