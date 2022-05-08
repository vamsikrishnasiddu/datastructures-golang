package main

import (
	"fmt"
	"log"
)

type Stack []int

type MyQueue struct {
	input  Stack
	output Stack
}

func Constructor() MyQueue {

	return MyQueue{
		input:  Stack{},
		output: Stack{},
	}

}

func (this *MyQueue) Push(x int) {

	this.input = this.input.Push(x)

}

func (s Stack) Empty() bool {
	if len(s) == 0 {
		return true
	}

	return false
}

func (this *MyQueue) Pop() int {
	var val int
	if !this.output.Empty() {

		this.output, val = this.output.Pop()

		return val

	}

	for this.input.Empty() == false {
		this.output = this.output.Push(this.input.Top())
		this.input, _ = this.input.Pop()
	}

	this.output, val = this.output.Pop()

	return val

}

func (this *MyQueue) Peek() int {

	if !this.output.Empty() {
		return this.output.Top()
	}

	for this.input.Empty() == false {
		this.output = this.output.Push(this.input.Top())
		this.input, _ = this.input.Pop()
	}

	return this.output.Top()

}

func (this *MyQueue) Empty() bool {

	if this.output.Empty() == true {

		if this.input.Empty() == true {
			return true
		}

	}

	return false

}

func (s Stack) Push(data int) Stack {
	s = append(s, data)
	return s
}

func (s Stack) Top() int {

	if len(s) == 0 {
		log.Fatal("stack is empty cannot get the top element.")
	}
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

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	fmt.Println(obj)
	param_2 := obj.Pop()
	fmt.Println(param_2)
	//param_3 := obj.Peek()
	//fmt.Println(param_3)
	param_4 := obj.Empty()
	fmt.Println(param_4)

	//fmt.Println(param_2, param_3, param_4)

}
