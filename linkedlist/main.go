package main

import (
	"fmt"
	"math"
)

type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
}

var temp **SinglyLinkedListNode

func AddTwoNumbers(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	var head *SinglyLinkedListNode
	var sum1, sum2 uint64 = 0, 0
	current1 := head1
	current2 := head2
	i := 0
	for current1 != nil {
		sum1 = sum1 + uint64(current1.data)*uint64(math.Pow10(i))
		current1 = current1.next
		i++
	}

	j := 0

	for current2 != nil {
		sum1 = sum2 + uint64(current2.data)*uint64(math.Pow10(j))
		current2 = current2.next
		j++
	}

	sum3 := sum1 + sum2

	//var num int

	for sum3 > 0 {
		rem := sum3 % 10
		head = insertNodeAtTheHead(head, int(rem))
		//num = num*10 + rem

		sum3 = sum3 / 10
	}

	head = reverse(head)
	return head

}

func AddTwoNumbersOrdern(head1 *SinglyLinkedListNode, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	var start *SinglyLinkedListNode
	temp := start

	var carry int

	current1, current2 := head1, head2

	for current1 != nil || current2 != nil || carry != 0 {

		sum := 0

		if current1 != nil {
			sum += current1.data
			current1 = current1.next
		}

		if current2 != nil {
			sum += current2.data
			current2 = current2.next
		}

		sum = sum + carry

		carry = sum / 10

		node := &SinglyLinkedListNode{
			data: sum % 10,
		}

		temp.next = node
		temp = temp.next

	}

	return start.next
}

func printSinglyLinkedList(head *SinglyLinkedListNode) {

	current := head

	for current != nil {
		fmt.Printf("%v  ", current.data)
		current = current.next
	}

	fmt.Println()
}

func DeleteNode(node *SinglyLinkedListNode) {

	node.data = node.next.data
	node.next = node.next.next

}

func HasCycle(head *SinglyLinkedListNode) bool {

	m := make(map[*SinglyLinkedListNode]bool)

	current1 := head

	for current1 != nil {

		if _, ok := m[current1]; !ok {
			m[current1] = true
		} else {
			return true
		}

		current1 = current1.next

	}

	return false

}

func HasCycleOrderN(head *SinglyLinkedListNode) bool {

	slow := head
	fast := head

	for fast != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false

}

func getIntersectionNode(headA, headB *SinglyLinkedListNode) *SinglyLinkedListNode {
	current1 := headA
	current2 := headB

	for current1 != nil {

		for current2 != nil {

			if current1 == current2 {
				return current1
			}

			current2 = current2.next

		}
		current2 = headB
		current1 = current1.next
	}

	return nil

}

func getIntersectionNodeOrderN(headA, headB *SinglyLinkedListNode) *SinglyLinkedListNode {

	if headA == nil || headB == nil {
		return nil
	}

	d1 := headA
	d2 := headB

	for d1 != d2 {

		if d1 == nil {
			d1 = headB
		} else {
			d1 = d1.next
		}

		if d2 == nil {
			d2 = headA
		} else {
			d2 = d2.next
		}

	}

	return d1

}

func insertNodeAtTheHead(head *SinglyLinkedListNode, data int) *SinglyLinkedListNode {

	n := &SinglyLinkedListNode{data: data}
	// we point the nodes next to the head
	// and we will make the node as head
	n.next = head
	head = n
	return head
}

func insertNodeAtTheTail(head *SinglyLinkedListNode, data int) *SinglyLinkedListNode {
	current := head

	node := &SinglyLinkedListNode{data: data}
	// we iterate until we reach the
	// last node and point the node to last nodes next
	// we return the head.
	for current.next != nil {
		current = current.next
	}

	current.next = node

	if node.data == 6 {
		temp = &node
		fmt.Println(temp)

	}

	return head
}
func reversePrint(head *SinglyLinkedListNode) {

	if head == nil {
		return

	} else {
		reversePrint(head.next)
		fmt.Println(head.data)
	}
}

func reverse(head *SinglyLinkedListNode) *SinglyLinkedListNode {

	// we will take two nodes one is temp and temp2
	var temp, temp2 *SinglyLinkedListNode

	// we iterate until we reach the end of the list
	for head != nil {

		// we store the head next in temp2. then we point the head to temp
		// we move the temp to head and head to temp2
		temp2 = head.next
		head.next = temp
		temp = head
		head = temp2

	}

	// since the last node is temp we make it as head and we return the head
	head = temp

	return head

}

func reverseKGroup(head *SinglyLinkedListNode, k int) *SinglyLinkedListNode {

	if head == nil || k == 1 {
		return head
	}

	var dummy *SinglyLinkedListNode = &SinglyLinkedListNode{data: 0}

	dummy.next = head

	var cur, next, prev *SinglyLinkedListNode

	cur = dummy
	next = dummy
	prev = dummy

	var count int = 0
	cur = head
	for cur != nil {
		count++
		cur = cur.next

	}

	cur = dummy

	for count >= k {
		cur = prev.next
		next = cur.next

		for i := 1; i < k; i++ {
			cur.next = next.next
			next.next = prev.next
			prev.next = next
			next = cur.next
		}

		prev = cur
		count = count - k
	}

	return dummy.next

}

func isPalindrome(head *SinglyLinkedListNode) bool {

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	slow = reverse(slow)

	fast = head

	for slow != nil {
		if slow.data != fast.data {
			return false
		}

		slow = slow.next
		fast = fast.next
	}

	return true

}

func mergeTwo(l1 *SinglyLinkedListNode, l2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	// if l1 is not there l2 is the merged list.
	if l1 == nil {
		return l2
	}

	// if l2 is not there l1 is the merged list.
	if l2 == nil {
		return l1
	}
	// we are assuming l1 element is smaller than l2 element
	// if we found l1 is greater than swap l1 l2 . so that l1 is now smaller.
	if l1.data > l2.data {
		l1, l2 = l2, l1

	}

	// result points to  the head of the combined lists.
	result := l1

	// iterate unitl we reach the end of the lists.
	for l1 != nil && l2 != nil {
		// temp variable to hold the node previous to l1.
		var temp *SinglyLinkedListNode

		// iterate until  l1 is less than or equal to l2.
		for l1 != nil && (l1.data <= l2.data) {
			temp = l1
			l1 = l1.next
		}

		// if l2 is greater than l1 point the temps next to l2
		// swap the l1  and l2
		temp.next = l2
		l1, l2 = l2, l1
	}
	// return the result which points to head of the combined lists.

	return result

}

func insertNodeAtTheIndex(head *SinglyLinkedListNode, data int, position int) *SinglyLinkedListNode {

	node := &SinglyLinkedListNode{data: data}
	if position == 0 {

		node.next = head
		head = node

	} else {
		current := head
		for i := 0; i < position-1; i++ {
			current = current.next

		}

		temp := current.next

		current.next = node
		node.next = temp
	}

	return head

}

func deleteNodeAtTheIndexFromStart(head *SinglyLinkedListNode, position int) *SinglyLinkedListNode {
	// we takes two pointers which is current and prevs both points to the head.
	current := head
	prev := head

	// if head is nil then nothing is there to delete we return from here.
	if head == nil {
		fmt.Println("List is empty")
	} else {
		// the deletion is handled differently for different scenarios.
		// when the position is 1 we just need to make the head points to
		// its next and we return the head.
		if position == 1 {

			head = current.next
			current.next = nil

		} else {
			// if the position is not equal to one. we move upto that position.
			// current node points to the node which we want to delete.
			// prev node points to the node just before current.
			for position != 1 {
				prev = current
				current = current.next
				position--
			}

			// we make the previous link to point to the next of the current.

			prev.next = current.next

			current.next = nil
		}

	}

	return head

}

func length(head *SinglyLinkedListNode) int {
	var length int
	for head != nil {
		length++
		head = head.next
	}

	return length
}

func removeNthFromEnd(head *SinglyLinkedListNode, n int) *SinglyLinkedListNode {

	if length(head) == 1 {
		return nil
	}

	current := head

	if n == length(head) {
		node := head
		head = head.next
		node.next = nil
	} else {

		for i := 1; i < length(head)-n-1; i++ {

			current = current.next
		}

		temp := current.next
		current.next = temp.next

	}

	return head

}

func removeTheNodeFromEndOrdern(head *SinglyLinkedListNode, n int) *SinglyLinkedListNode {
	// we create a dummy node start and make its next point to the head
	start := &SinglyLinkedListNode{}

	start.next = head
	//we create two nodes one is fast and slow  will initially points to dummy node.
	slow := start
	fast := start
	// fast pointer moves fastly. we move it upto n postions.
	for i := 0; i < n; i++ {
		fast = fast.next
	}

	// now we have fast pointer which is already at n
	// which will require length -n steps to reach the end of the linked list
	// we will move slow and fast one at a time.
	// By the time fast reached to end. Slow covers the length-n steps. So then we
	//remove the node which is pointed by slows next and point it to slow.next.next

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	// we are returning start.next because if we consider the edge
	// case where we need to remove last node from end of the linked list
	// which is itself the first node. then the existing head will be cut.
	// we make slow to point it to the next node. since the slow and start
	// are same we have the head starting from start.next
	return start.next
}

func findTheMiddleOfLinkedList(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	n := length(head)

	var mid int

	mid = n/2 + 1

	start := &SinglyLinkedListNode{}
	start.next = head

	for i := 0; i < mid; i++ {
		start = start.next
	}

	return start

}

func detectCycle(head *SinglyLinkedListNode) *SinglyLinkedListNode {

	//T.C Order(n)
	//S.C order(n)

	m := make(map[*SinglyLinkedListNode]int)

	current := head

	for current != nil {

		if _, ok := m[current]; !ok {

			m[current] = current.data

		} else {
			return current
		}

		current = current.next
	}

	return nil

}

func detectCycleOrdern(head *SinglyLinkedListNode) *SinglyLinkedListNode {

	if head == nil || head.next == nil {
		return nil
	}

	slow := head
	fast := head
	entry := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {

			for slow != entry {
				slow = slow.next
				entry = entry.next
			}

			return entry

		}
	}

	return nil

}

func findTheMiddleOfLinkedListOrdern(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// we solve it using the tortoise method.
	// we initially make fast and slow point to the head.
	fast := head
	slow := head
	// we move the fast pointer two times and slow pointer one time.
	// unitl fast pointer reaches the nil incase of even
	// or fast pointers next reaches nil incase of odd.
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// by the time fast pointer reaches the end. Slow will point to the mid of the linked list.

	return slow

}
func main() {

	/*
			[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
		[5,6,4]

	*/

	/*
		     [-1,-7,7,-4,19,6,-9,-5,-2,-5]
		6
	*/
	list1 := SinglyLinkedList{
		head: &SinglyLinkedListNode{
			data: -1,
		},
	}

	// for i := 0; i < 30; i++ {
	// 	list1.head = insertNodeAtTheTail(list1.head, 0)
	// }

	list1.head = insertNodeAtTheTail(list1.head, -7)
	list1.head = insertNodeAtTheTail(list1.head, 7)
	list1.head = insertNodeAtTheTail(list1.head, -4)
	list1.head = insertNodeAtTheTail(list1.head, 19)
	list1.head = insertNodeAtTheTail(list1.head, 6)
	fmt.Println("=====================================")
	list1.head = insertNodeAtTheTail(list1.head, -9)
	list1.head = insertNodeAtTheTail(list1.head, -5)
	list1.head = insertNodeAtTheTail(list1.head, -2)
	list1.head = insertNodeAtTheTail(list1.head, -5)
	fmt.Println("=====================================")

	curr := list1.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = *temp

	node := detectCycle(list1.head)

	fmt.Println(node)

	//list1.head = insertNodeAtTheTail(list1.head, 3)

	//printSinglyLinkedList(list1.head)

	//list1.head = findTheMiddleOfLinkedListOrdern(list1.head)

	// printSinglyLinkedList(list1.head)
	// fmt.Println("==============================")

	// list2 := SinglyLinkedList{
	// 	head: &SinglyLinkedListNode{
	// 		data: 5,
	// 	},
	// }

	// list2.head = insertNodeAtTheTail(list2.head, 6)

	// list2.head = insertNodeAtTheTail(list2.head, 4)

	// printSinglyLinkedList(list2.head)
	// fmt.Println("==============================")

	// head3 := AddTwoNumbers(list1.head, list2.head)

	printSinglyLinkedList(list1.head)

	// list1.head = reverseKGroup(list1.head, 3)

	// printSinglyLinkedList(list1.head)

	// list.head = findTheMiddleOfLinkedListOrdern(list.head)

	// printSinglyLinkedList(list.head)

	// fmt.Println("length of the list", length(list.head))

	// fmt.Println("===============================")

	// list.head = deleteNodeAtTheIndexFromStart(list.head, 1)

	// printSinglyLinkedList(list.head)

	// fmt.Println("length of the list", length(list.head))

	//deleteNodeAtTheIndexFromStart(list.head, 0)
	//removeNthFromEnd(list.head, 3)

	//printSinglyLinkedList(list.head)

	//fmt.Println("length of the list", length(list.head))

	//printSinglyLinkedList(list.head)

	//reversePrint(list.head)

	//printSinglyLinkedList(list.head)

	//reversePrint(list.head)

	//list.head = deleteNodeAtTheIndex(list.head, 1)

	//reversePrint(list.head)
	// fmt.Println("================================")
	// printSinglyLinkedList(list.head)
	//list.head = deleteNodeAtTheIndex(list.head, 2)

	//reversePrint(list.head)
	// fmt.Println("================================")
	// printSinglyLinkedList(list.head)

	// list.head = deleteNodeAtTheIndex(list.head, 9)
	// fmt.Println("================================")
	// printSinglyLinkedList(list.head)

	// printSinglyLinkedList(list.head)

	// list.head = deleteNodeAtTheIndex(list.head, 1)
	// printSinglyLinkedList(list.head)

	// list.head = deleteNodeAtTheIndex(list.head, 2)
	// printSinglyLinkedList(list.head)

	// list.head = deleteNodeAtTheIndex(list.head, 5)
	// printSinglyLinkedList(list.head)

	//list.head = reverse(list.head)

	//printSinglyLinkedList(list.head)

	// list2 := SinglyLinkedList{
	// 	head: &SinglyLinkedListNode{
	// 		data: 1,
	// 	},
	// }

	//list2.head = insertNodeAtTheTail(list2.head, 3)

	//list2.head = insertNodeAtTheTail(list2.head, 4)

	//list2.head = insertNodeAtTheTail(list2.head, 4)

	//list2.head = insertNodeAtTheHead(list2.head, 10)

	//printSinglyLinkedList(list2.head)

	//fmt.Println("======================================")

	// list2.head = insertNodeAtTheIndex(list2.head, 15, 5)
	// list2.head = insertNodeAtTheIndex(list2.head, 9, 0)
	// list2.head = insertNodeAtTheIndex(list2.head, 8, 1)

	//head3 := mergeTwo(list.head, list2.head)

	//printSinglyLinkedList(head3)

}
