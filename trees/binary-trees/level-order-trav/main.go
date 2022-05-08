package main

import (
	"fmt"
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	items []*TreeNode
}

func (q *Queue) Enqueuee(val *TreeNode) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeuee() {
	l := len(q.items)

	if l == 0 {
		log.Fatal("queuee is empty.")
	}

	q.items = q.items[1:]
}

func (q *Queue) isEmpty() bool {

	if len(q.items) == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Front() *TreeNode {
	return q.items[0]
}

func levelOrder(root *TreeNode) [][]int {

	var ans [][]int

	if root == nil {
		return ans
	}

	var q Queue
	q.Enqueuee(root)

	for !q.isEmpty() {

		size := q.Size()

		var level []int = []int{}

		for i := 0; i < size; i++ {

			treenode := q.Front()
			q.Dequeuee()

			if treenode.Left != nil {
				q.Enqueuee(treenode.Left)
			}

			if treenode.Right != nil {
				q.Enqueuee(treenode.Right)
			}

			level = append(level, treenode.Val)

		}
		ans = append(ans, level)

	}

	return ans

}

func main() {

	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	a := levelOrder(root)

	fmt.Println(a)

}
