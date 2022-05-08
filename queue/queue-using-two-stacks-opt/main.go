package main

import (
	"fmt"
	"log"
)

type Stack []int

type Queue struct {
	input  Stack
	output Stack
}

func (q *Queue) Enqueue(data int) {
	q.input = q.input.Push(data)
}

func (s Stack) Empty() bool {
	if len(s) == 0 {
		return true
	}

	return false
}

func (q *Queue) Dequeue() {

	if !q.output.Empty() {
		q.output, _ = q.output.Pop()
	} else {

		for q.input.Empty() == false {
			q.output = q.output.Push(q.input.Top())
			q.input, _ = q.input.Pop()
		}

		q.output, _ = q.output.Pop()
	}

}

func (q *Queue) Top() int {

	if !q.output.Empty() {
		return q.output.Top()
	}

	for q.input.Empty() == false {
		q.output = q.output.Push(q.input.Top())
		q.input, _ = q.input.Pop()
	}

	return q.output.Top()

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

	q := Queue{input: Stack{}, output: Stack{}}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q)

	top := q.Top()

	fmt.Println(top)

	q.Dequeue()

	fmt.Println(q)

	top = q.Top()

	fmt.Println(top)

	q.Dequeue()

	fmt.Println(q)

	// q.Dequeue()

	// fmt.Println(q)

}
