package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() string {
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {

	if len(s.items) == 0 {
		return true
	}

	return false
}

func reverseWords(s string) string {

	var st Stack

	strArr := strings.Fields(s)

	for i := 0; i < len(strArr); i++ {
		st.Push(strArr[i])
	}

	i := 0

	for !st.isEmpty() {
		strArr[i] = st.Top()
		st.Pop()
		i++
	}

	return strings.Join(strArr, " ")

}

func main() {

	s := "a good   example"

	fmt.Println(reverseWords(s))

}
