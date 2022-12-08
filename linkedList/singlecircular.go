package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
	tail *Node
	len  int
}

func (l *linkedList) addLast(val int) {
	newNode := &Node{value: val}
	if l.head == nil {
		newNode.next = newNode
		l.head = newNode

	} else {
		newNode.next = l.tail.next
		l.tail.next = newNode
		// fmt.Println(c.tail.next)
		// fmt.Println(newNode)
	}
	l.tail = newNode
	l.len++
}

func (l *linkedList) addFirst(val int) {
	newNode := &Node{value: val}
	if l.len == 0 {
		newNode.next = newNode
		l.head = newNode
	} else {
		l.tail.next = newNode
		newNode.next = l.head
		l.head = newNode
	}
	l.len++
}

func (l *linkedList) insertAny(val, pos int) {
	newNode := &Node{value: val}
	ptr := l.head
	for i := 1; i < pos-1; i++ {
		ptr = ptr.next
	}
	newNode.next = ptr.next
	ptr.next = newNode
	l.len++
}

func (l *linkedList) display() {
	p := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Printf("\n")
}

func (l *linkedList) delLast() {
	ptr := l.head
	for i := 0; i < l.len-1; i++ {
		ptr = ptr.next
	}
	ptr.next = l.head
	l.len--
}
func (l *linkedList) delFirst() {
	l.tail.next = l.head.next
	l.head = l.head.next
	l.len--
}

func (l *linkedList) delAny(pos int) {
	ptr := l.head
	for i := 0; i < pos-1; i++ {
		ptr = ptr.next
	}
	ptr.next = ptr.next.next
	l.len--
}

func (l *linkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

func main() {
	c := linkedList{}
	c.addLast(1)
	c.addLast(2)
	c.addLast(3)
	c.addLast(4)
	c.display()
	c.addFirst(50)
	c.addFirst(70)
	c.display()
	c.insertAny(10, 3)
	c.display()
	c.delLast()
	c.display()
	c.delFirst()
	c.display()
	c.delAny(3)
	c.display()
}
