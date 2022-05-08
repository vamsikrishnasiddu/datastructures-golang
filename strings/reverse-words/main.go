package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	s = strings.Join(strings.Fields(s), " ")

	strArr := strings.Split(s, " ")

	start := 0
	end := len(strArr) - 1

	for start < end {

		strArr[start], strArr[end] = strArr[end], strArr[start]
		start++
		end--
	}

	s = strings.Join(strArr, " ")

	return s

}

func main() {

	var s string = "a good   example"

	fmt.Println(reverseWords(s))

}
