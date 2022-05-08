package main

import (
	"fmt"
	"strings"
)

func reversewords(s string) {

	// b := []byte(s)

	str := ""
	temp := ""

	var sarr []string = []string{}

	for i := 0; i < len(s); i++ {

		fmt.Println(s[i])

		if s[i] == 32 {

			if s[i-1] == 32 {
				continue
			}

			temp = temp + str
			sarr = append(sarr, temp)

			str = ""
			temp = ""

		}

		str = str + string(s[i])

	}

	sarr = append(sarr, str)

	fmt.Println(sarr)

	s = strings.Join(sarr, " ")

	fmt.Println("==================================")

	for i := 0; i < len(s); i++ {

		fmt.Println(s[i])

	}

	fmt.Println(sarr)
}

func main() {
	var s string = "a good   example"

	reversewords(s)
}
