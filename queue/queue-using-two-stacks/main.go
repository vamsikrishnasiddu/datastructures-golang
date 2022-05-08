package main

import (
	"fmt"
	"log"
)

type Stack []int

type Queue struct {
	s1 Stack
	s2 Stack
}

func (q *Queue) Enqueue(data int) {

	for len(q.s1) != 0 {
		q.s2 = q.s2.Push(q.s1.Top())
		q.s1, _ = q.s1.Pop()

	}

	q.s1 = q.s1.Push(data)

	for len(q.s2) != 0 {
		q.s1 = q.s1.Push(q.s2.Top())
		q.s2, _ = q.s2.Pop()

	}

}

func (q *Queue) Dequeue() {

	q.s1, _ = q.s1.Pop()

}

func (q *Queue) Top() int {

	return q.s1.Top()

}

func (s Stack) Push(data int) Stack {
	s = append(s, data)
	return s
}

func (s Stack) Top() int {
	return s[len(s)-1]
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

	q := Queue{s1: Stack{}, s2: Stack{}}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q)

	top := q.Top()

	q.Dequeue()

	fmt.Println(top)

	fmt.Println(q)

	// q.Dequeue()

	// fmt.Println(q)

}
