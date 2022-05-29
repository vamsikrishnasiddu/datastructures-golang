package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {

	length := 0
	max_length := math.MinInt32

	for i := 0; i < len(s); i++ {
		set := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if !set[s[j]] {
				set[s[j]] = true
				length = j - i + 1
				//fmt.Println("length", length)

				if length > max_length {
					max_length = length
				}

			} else {
				break
			}

		}
	}

	return max_length

}

func lengthOfLongestSubstringO2n(s string) int {
	length := 0
	max_length := 0
	l := 0
	set := make(map[byte]bool)

	for r := 0; r < len(s); r++ {
		if set[s[r]] {

			for l < r && set[s[r]] {
				set[s[l]] = false
				l++
			}
		}

		set[s[r]] = true
		length = r - l + 1

		if length > max_length {
			max_length = length
		}
	}

	return max_length

}

func lengthOfLongestSubstringOn(s string) int {
	length := 0
	max_length := 0
	l := 0
	set := make(map[byte]int)

	for r := 0; r < len(s); r++ {
		if val, ok := set[s[r]]; ok {

			if val+1 > l {
				l = val + 1
				//fmt.Println("l", l)
			}
		}

		set[s[r]] = r
		length = r - l + 1
		if length > max_length {
			max_length = length
		}
	}

	return max_length

}

func main() {
	s := "abcaabcdba"

	fmt.Println(lengthOfLongestSubstringOn(s))
}
