package main

import (
	"fmt"
	"log"
)

type Stack []int

func (s Stack) Push(data int) Stack {
	s = append(s, data)
	return s
}

func (s Stack) Pop() (Stack, int) {
	l := len(s)

	if l == 0 {
		log.Fatalf("Cannot pop from empty stack.")
	}
	data := s[len(s)-1]
	s = s[:len(s)-1]
	return s, data
}

func main() {

	var s Stack = []int{1, 2, 3, 4}

	s = s.Push(5)
	fmt.Println(s)
	s, data := s.Pop()
	fmt.Println(data)
	s, data = s.Pop()
	fmt.Println(data)
	s, data = s.Pop()
	fmt.Println(data)
	s, data = s.Pop()
	fmt.Println(data)
	s, data = s.Pop()
	fmt.Println(data)
	s, data = s.Pop()
	fmt.Println(data)

}
