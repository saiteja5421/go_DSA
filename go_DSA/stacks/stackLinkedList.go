package main

import "fmt"

type node struct {
	value int
	next  *node
}
type stackLinedlist struct {
	head *node
	len  int
}

func (l *stackLinedlist) push(e int) {
	newNode := node{value: e}
	if l.len == 0 {
		l.head = &newNode
	} else {
		newNode.next = l.head
		l.head = &newNode
	}
	l.len++
}

func (l *stackLinedlist) pop() int {
	if l.head != nil {
		p := l.head.value
		l.head = l.head.next
		l.len--
		return p
	} else {
		fmt.Println("stack is empty")
		return 0
	}
}
func (l *stackLinedlist) top() int {
	if l.head == nil {
		fmt.Println("stack is empty")
		return 0
	}
	e := l.head
	return e.value
}
func (l *stackLinedlist) display() {
	p := l.head
	for p != nil {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Printf("\n")
}

func main() {
	l := stackLinedlist{}
	l.push(4)
	l.push(50)
	l.push(30)
	l.display()
	a := l.pop()
	// l.pop()
	// l.pop()
	fmt.Println(a)
	l.display()
	b := l.top()
	fmt.Println(b)

}
