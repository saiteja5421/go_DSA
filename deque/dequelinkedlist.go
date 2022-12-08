package main

import "fmt"

type node struct {
	data int
	next *node
}
type queuelinkedlist struct {
	front *node
	rear  *node
	len   int
}

func (l *queuelinkedlist) isempty() bool {
	return l.len == 0
}
func (l *queuelinkedlist) addLast(e int) {
	newNode := new(node)
	newNode.data = e
	if l.isempty() {
		l.front = newNode
	} else {
		l.rear.next = newNode
	}
	l.rear = newNode
	l.len++
}
func (l *queuelinkedlist) addFirst(e int) {
	newNode := new(node)
	newNode.data = e
	if l.isempty() {
		l.front = newNode
	} else {
		newNode.next = l.front
		l.front = newNode
	}
	l.rear = newNode
	l.len++
}
func (l *queuelinkedlist) removeFirst() int {
	if l.isempty() {
		fmt.Println("dequeue is empty")
		return 0
	} else {
		e := l.front.data
		l.front = l.front.next
		l.len--
		return e
	}

}

func (l *queuelinkedlist) removeLast() int {
	if l.isempty() {
		fmt.Println("deque is empty")
		return 0
	} else {
		p := l.front
		for i := 0; i < l.len-2; i++ {
			p = p.next
		}
		l.rear = p
		e := p.next.data
		l.rear.next = nil
		return e
	}
}
func (l *queuelinkedlist) display() {
	p := l.front
	for p != nil {
		fmt.Printf("%d ", p.data)
		p = p.next
	}
	fmt.Printf("\n")
}

func main() {
	l := queuelinkedlist{}
	l.addLast(20)
	l.addLast(30)
	l.addFirst(50)
	l.display()
	a := l.removeFirst()
	b := l.removeLast()
	fmt.Println("first element removed is", a)
	fmt.Println("first element removed Last is", b)
	l.display()
}
