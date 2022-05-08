package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	items []Pair
}

type Pair struct {
	Node *TreeNode
	Num  int
}

func (s *Stack) Push(p Pair) {
	s.items = append(s.items, p)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() Pair {

	return s.items[len(s.items)-1]

}

func (s *Stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}

	return false
}

func preInPostOrderTraversal(root *TreeNode) {

	var pre, in, post = []int{}, []int{}, []int{}

	var st Stack

	p := Pair{root, 1}

	st.Push(p)

	for !st.isEmpty() {

		it := st.Top()

		st.Pop()

		if it.Num == 1 {
			pre = append(pre, it.Node.Val)
			it.Num++
			st.Push(it)

			if it.Node.Left != nil {
				st.Push(Pair{it.Node.Left, 1})
			}
		} else if it.Num == 2 {
			in = append(in, it.Node.Val)
			it.Num++
			st.Push(it)

			if it.Node.Right != nil {
				st.Push(Pair{it.Node.Right, 1})
			}
		} else {
			post = append(post, it.Node.Val)
		}

	}

	fmt.Println(pre, in, post)

}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	// //result := preorderTraversal(root)

	preInPostOrderTraversal(root)

	// fmt.Println(result)

}
