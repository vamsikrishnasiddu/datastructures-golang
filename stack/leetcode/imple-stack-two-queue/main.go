package main

import (
	"fmt"
	"log"
)

type Queue []int

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

type MyStack struct {
	q1 Queue
	q2 Queue
}

func Constructor() MyStack {

	return MyStack{q1: Queue{}, q2: Queue{}}

}

func (this *MyStack) Push(x int) {

	this.q2 = this.q2.Enqueue(x)

	for len(this.q1) != 0 {
		this.q2 = this.q2.Enqueue(this.q1.Top())
		_, this.q1 = this.q1.Dequeue()
	}

	this.q1, this.q2 = this.q2, this.q1

}

func (this *MyStack) Pop() int {
	var val int
	val, this.q1 = this.q1.Dequeue()

	return val

}

func (this *MyStack) Top() int {

	return this.q1.Top()

}

func (this *MyStack) Empty() bool {

	if len(this.q1) == 0 {
		return true
	}

	return false

}

func main() {
	mystack := Constructor()

	mystack.Push(1)
	//mystack.Push(2)
	//param_2 := mystack.Top()
	param_3 := mystack.Pop()
	param_4 := mystack.Empty()
	fmt.Println(param_3, param_4)

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
