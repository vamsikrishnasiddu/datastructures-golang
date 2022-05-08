package main

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	lock sync.Mutex
	s    []int
}

func NewStack() *Stack {
	return &Stack{sync.Mutex{}, make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, v)
}

func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)

	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]

	return res, nil

}

func main() {

	s := NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
