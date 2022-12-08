package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}
type linkedList struct {
	head *node
	tail *node
	len  int
}

// func (l *linkedList) addLast(value int) {
// 	// Create new node
// 	tmp := new(node)
// 	tmp.value = value

// 	//fmt.Println("\nInserting", tmp, "at the start ")
// 	if l.head == nil {
// 		l.head = tmp
// 		tmp.next=l.head
// 		l.len++
// 	}
// 		p := l.head
// 		for ; p.next != l.head; p = p.next {
// 		}
// 	    p.next = tmp
// 		tmp.next = l.head
// 		l.head = tmp
// 		l.len++

// }

func (l *linkedList) addLast(value int) {
	newNode := &node{value: value}
	if l.head == nil {
		newNode.next = newNode
		l.head = newNode
	} else {
		newNode.next = l.tail.next
		l.tail.next = newNode
		// newNode.prev = l.tail
	}

	l.tail = newNode
	l.len++

}
func (l linkedList) displayrev() {
	//p := l.tail
	// sb := strings.Builder{}
	// for p != nil {
	// 	//fmt.Println(p.value)
	// 	sb.WriteString(fmt.Sprintf("%d ", p.value))
	// 	p = p.prev
	// }
	// return sb.String()
	p := l.head
	for p != l.head {
		fmt.Printf("-->%d", p.value)
		p = p.next
	}

}

func (c *linkedList) display() {
	p := c.head
	for i := 0; i < c.len; i++ {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Printf("\n")
}
func main() {
	l := linkedList{}
	l.addLast(2)
	l.addLast(4)
	l.addLast(6)
	l.addLast(8)
	l.addLast(10)
	l.addLast(20)
	//fmt.Println(l)
	l.display()
}
