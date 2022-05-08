package main

import "fmt"

func myPow(x float64, n int) float64 {

	var sum float64 = 1

	var nn int = n

	if nn < 0 {
		nn = -nn
	}

	for nn > 0 {

		if nn%2 == 1 {
			sum = sum * x
			nn = nn - 1
		} else {
			x = x * x
			nn = nn / 2
		}
	}

	if n < 0 {
		return 1.0 / sum
	}

	return sum

}

func main() {

	fmt.Println(myPow(2, -2147483648))

}
