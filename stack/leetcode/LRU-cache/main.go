package main

type Node struct {
	Key int
	Val int

	Next *Node
	Prev *Node
}

type LRUCache struct {
	Capacity int
	Map      map[int]*Node
	Head     *Node
	Tail     *Node
}

func Consructor(capacity int) LRUCache {

	l := LRUCache{Capacity: capacity, Map: make(map[int]*Node), Head: &Node{-1, -1, nil, nil}, Tail: &Node{-1, -1, nil, nil}}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	return l

}

func (l *LRUCache) addnode(newnode *Node) {
	temp := l.Head
	newnode.Next = temp
	newnode.Prev = l.Head
	l.Head.Next = newnode
	temp.Prev = newnode
}

func deletenode(delnode *Node) {
	delprev := delnode.Prev
	delNext := delnode.Next
	delprev.Next = delNext
	delNext.Prev = delprev
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.Map[key]; ok {
		resnode := val
		res := resnode.Val
		delete(this.Map, key)
		deletenode(resnode)
		this.addnode(resnode)
		this.Map[key] = this.Head.Next
		return res

	}

	return -1
}

func (this *LRUCache) Put(key int, val int) {
	if val, ok := this.Map[key]; ok {
		existingnode := val
		delete(this.Map, key)
		deletenode(existingnode)
	}

	if len(this.Map) == this.Capacity {
		k := this.Tail.Prev.Key
		delete(this.Map, k)
		deletenode(this.Tail.Prev)
	}

	this.addnode(&Node{Key: key, Val: val})
	this.Map[key] = this.Head.Next
}

func main() {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
