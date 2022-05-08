package main

import (
	"fmt"
	"log"
)

type Queue []int

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

	q := Queue{}
	q = q.Enqueue(1)
	q = q.Enqueue(2)
	q = q.Enqueue(3)
	q = q.Enqueue(4)
	q = q.Enqueue(5)
	fmt.Println(q)

	data, q := q.Dequeue()
	fmt.Println(data)
	data, q = q.Dequeue()
	fmt.Println(data)
	data, q = q.Dequeue()
	fmt.Println(data)
	data, q = q.Dequeue()
	fmt.Println(data)
	data, q = q.Dequeue()
	fmt.Println(data)
	data, q = q.Dequeue()
	fmt.Println(data)

}
