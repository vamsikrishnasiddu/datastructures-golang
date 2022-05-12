package main

import
(
    "fmt"
)

func isValid(s string) bool {

	var q []rune

    var dummy *int32
    dummy = nil

    fmt.Println(*dummy)

	for _, c := range s {

		switch c {
		case '(':
			q = append(q, ')')
		case '{':
			q = append(q, '}')
		case '[':
			q = append(q, ']')
		default:
			if len(q) == 0 || q[len(q)-1] != c {
				return false
			}

			q = q[:len(q)-1]
		}
	}

	return len(q) == 0
}

func main() {
	isValid("()[]{}")
}
