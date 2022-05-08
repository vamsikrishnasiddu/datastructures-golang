package main

import (
	"fmt"
	"log"
)

type Queue []int

type Stack struct {
	q1 Queue
	q2 Queue
}

func (s *Stack) Push(data int) {
	s.q2 = s.q2.Enqueue(data)

	for len(s.q1) != 0 {
		s.q2 = s.q2.Enqueue(s.q1.Top())
		_, s.q1 = s.q1.Dequeue()
	}

	s.q1, s.q2 = s.q2, s.q1
}

func (s *Stack) Pop() {

	if len(s.q1) == 0 {
		log.Fatalf("Stack is empty.")
	}
	_, s.q1 = s.q1.Dequeue()
}

func (q Queue) Top() int {

	return q[0]

}

func (q Queue) Enqueue(data int) Queue {
	q = append(q, data)

	return q
}

func (q Queue) Dequeue() (int, Queue) {
	l := len(q)

	if l == 0 {
		log.Fatal("Queue is empty")
	}
	data := q[0]
	q = q[1:]

	return data, q
}

func main() {

	s := Stack{
		q1: Queue{},
		q2: Queue{},
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s)

	s.Pop()

	fmt.Println(s)
	s.Pop()

	fmt.Println(s)
	s.Pop()

	fmt.Println(s)
	s.Pop()

	fmt.Println(s)

}
