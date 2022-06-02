package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func preorder(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.data)

	preorder(node.left)
	preorder(node.right)
}

func inorder(node *Node) {
	if node == nil {
		return
	}

	inorder(node.left)
	fmt.Println(node.data)
	inorder(node.right)
}

func postorder(node *Node) {
	if node == nil {
		return
	}

	postorder(node.left)
	postorder(node.right)
	fmt.Println(node.data)
}

func main() {
	root := &Node{data: 1}

	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.right = &Node{data: 5}

	fmt.Println(root)

	//preorder(root)
	//inorder(root)
	preorder(root)
}
