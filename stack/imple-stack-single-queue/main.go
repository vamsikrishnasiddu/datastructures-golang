package main

import (
	"fmt"
	"log"
)

type Queue []int

type Stack struct {
	q Queue
}

func (s *Stack) Push(data int) {
	s.q = s.q.Enqueue(data)

	for i := 0; i < len(s.q)-1; i++ {

		s.q = s.q.Enqueue(s.q.Front())
		_, s.q = s.q.Dequeue()

	}

}

func (s *Stack) Pop() {

	if len(s.q) == 0 {
		log.Fatalf("Stack is empty.")
	}
	_, s.q = s.q.Dequeue()
}

func (q Queue) Front() int {
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

	s := Stack{q: Queue{}}

	s.Push(1)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)

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
