package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	items []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.items = append(s.items, node)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() *TreeNode {

	return s.items[len(s.items)-1]

}

func (s *Stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}

func preorderTraversal(root *TreeNode) []int {

	var result []int = []int{}

	var st Stack

	st.Push(root)

	for !st.isEmpty() {
		node := st.Top()

		result = append(result, node.Val)

		st.Pop()

		if node.Right != nil {
			st.Push(node.Right)
		}

		if node.Left != nil {
			st.Push(node.Left)
		}

	}

	return result

}

func inorderTraversal(root *TreeNode) []int {

	var result []int = []int{}

	var st Stack

	node := root

	for true {

		if node != nil {
			st.Push(node)
			node = node.Left
		} else {

			if st.isEmpty() {
				break
			}
			node = st.Top()
			st.Pop()
			result = append(result, node.Val)
			node = node.Right
		}

	}

	return result

}

func postorderTraversal(root *TreeNode) []int {

	var result = []int{}

	var st1, st2 Stack

	st1.Push(root)

	for !st1.isEmpty() {
		root = st1.Top()
		st2.Push(root)
		st1.Pop()

		if root.Left != nil {
			st1.Push(root.Left)
		}

		if root.Right != nil {
			st1.Push(root.Right)
		}
	}

	for !st2.isEmpty() {
		result = append(result, st2.Top().Val)
		st2.Pop()
	}

	return result

}

func postorderTraversalSingleStack(root *TreeNode) []int {

	var result = []int{}

	var st Stack

	cur := root

	for cur != nil || !st.isEmpty() {

		if cur != nil {
			st.Push(cur)
			cur = cur.Left
		} else {
			temp := st.Top().Right

			if temp == nil {

				temp = st.Top()
				st.Pop()
				result = append(result, temp.Val)

				for !st.isEmpty() && temp == st.Top().Right {
					temp = st.Top()
					st.Pop()
					result = append(result, temp.Val)
				}

			} else {
				cur = temp
			}
		}
	}

	return result

}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Left.Right = &TreeNode{Val: 7}
	root.Right.Left.Right.Right = &TreeNode{Val: 8}

	//result := preorderTraversal(root)

	result := postorderTraversalSingleStack(root)

	fmt.Println(result)

}
