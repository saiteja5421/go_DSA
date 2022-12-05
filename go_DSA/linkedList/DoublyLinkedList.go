package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
	prev  *node
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
		l.tail = newNode
		l.len++
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
		l.len++
	}
}

func (l *linkedList) addFirst(value int) {
	newNode := new(node)
	newNode.value = value
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.len++
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
		l.len++
	}
}

func (l *linkedList) remove(value int) {
	newNode := new(node)
	for iteration := l.head; iteration != nil; iteration = iteration.next {
		if iteration.value == value {
			fmt.Println(iteration.value)
			newNode.next = iteration.next
			newNode.next.prev = iteration.prev
			l.len--
		}
		newNode = iteration

	}
}
func (l *linkedList) addany(value int, pos int) {
	p := new(node)
	i := l.head
	p.value = value
	if pos < 0 {
		fmt.Println("index cannot be less than 0")
	} else if pos > l.len {
		fmt.Println(l.len)
		fmt.Println("pos out of range")
	} else {
		j := 0
		for j < pos-1 {
			i = i.next
			j = j + 1
		}
		p.next = i.next
		i.next = p
		p.next.prev = p
		p.prev = i

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
		i.next.prev = i
		l.len--
	}
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

func (l linkedList) display() {
	p := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Printf("\n")
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
	p := l.tail
	for p != nil {
		fmt.Printf("--> %d", p.value)
		p = p.prev
	}
	// p=p.prev

}

func main() {
	l := linkedList{}
	l.add(3)
	l.add(5)
	l.add(15)
	l.add(25)
	l.add(35)
	l.display()
	fmt.Println(l.len)
	//fmt.Println(l)
	l.addFirst(50)
	l.display()
	//fmt.Println(l)
	l.removeAtPos(2)
	l.display()
	//fmt.Println(l)
	l.addany(60, 2)
	l.display()
	//fmt.Println(l)
	l.displayrev()
	//l.display()
}
