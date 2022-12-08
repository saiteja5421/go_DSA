package main

import (
	"fmt"
	"strings"
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

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value

	if l.head == nil {
		l.head = newNode
		l.len++
	} else {
		iteration := l.head
		for ; iteration.next != nil; iteration = iteration.next {
		}
		iteration.next = newNode
		l.len++
	}
}

func (l *linkedList) addFirst(value int) {
	newnode := new(node)
	newnode.value = value

	if l.head == nil {
		l.head = newnode
		l.len++
		//l.tail=newnode
	} else {
		newnode.next = l.head
		l.head = newnode
		l.len++
	}
}

func (l *linkedList) removeFirst() {
	if l.len == 0 {
		fmt.Println("node is empty")
		return
	} else {
		//i := l.head
		l.head = l.head.next
		l.len--
	}
}

func (l *linkedList) remove(value int) {
	sai := new(node)

	for iteration := l.head; iteration != nil; iteration = iteration.next {
		if iteration.value == value {
			sai.next = iteration.next
			l.len--
			return
		}
		sai = iteration
	}

}

func (l *linkedList) addAtPos(value int, pos int) {
	sai := node{value: value}
	if pos < 0 {
		fmt.Println("index cannot be less than 0")
	} else if pos > l.len {
		fmt.Println("pos out of range")
	} else {
		i := l.head
		j := 0
		for j < pos-1 {
			i = i.next
			j = j + 1
		}
		sai.next = i.next
		i.next = &sai
		l.len++
	}
}

func (l *linkedList) removeAtPos(pos int) {
	//newNode :=node{value: value}
	if pos < 0 {
		fmt.Println("index cannot be less than 0")
	} else if pos > l.len {
		fmt.Println("pos out of range")
	} else {
		i := l.head
		j := 0
		for j < pos-1 {
			i = i.next
			j = j + 1
		}
		i.next = i.next.next
		l.len--
	}
}

func (l linkedList) display() []int {
	var s []int
	for iteration := l.head; iteration != nil; iteration = iteration.next {
		s = append(s, iteration.value)
	}
	return s
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

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iteration := l.head; iteration != nil; iteration = iteration.next {
		sb.WriteString(fmt.Sprintf("%d ", iteration.value))
	}
	return sb.String()
}
func main() {
	l := linkedList{}
	l.addFirst(50)
	fmt.Println(l)
	l.add(3)
	l.add(2)
	l.add(4)
	l.add(5)
	fmt.Println(l.len)
	//l.remove(50)
	//fmt.Println(l.len)
	fmt.Println(l)
	l.add(20)
	fmt.Println(l)
	l.addFirst(70)
	fmt.Println(l.len)
	fmt.Println(l)
	l.addAtPos(10, 5)
	fmt.Println(l)
	l.removeAtPos(0)
	fmt.Println(l)
	// l.removeFirst()
	// fmt.Println(l)
	l.remove(20)
	fmt.Println(l)
	a := l.Search(10)
	fmt.Println(a)
}
