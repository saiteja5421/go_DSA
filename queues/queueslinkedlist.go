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
func (l *queuelinkedlist) enqueue(e int) {
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
func (l *queuelinkedlist)dequeue() int{
	if l.isempty(){
		fmt.Println("queue is empty")
        return 0
	}else{
		e := l.front.data
		l.front=l.front.next
		return e
	}
}
func (l *queuelinkedlist) display() {
	p := l.front
	for p != nil {
		fmt.Printf("%d ",p.data)
		p=p.next
	}
	fmt.Printf("\n")
}

func main() {
	l := queuelinkedlist{}
	l.enqueue(20)
	l.enqueue(30)
   l.display()
   l.dequeue()
   l.display()
}
