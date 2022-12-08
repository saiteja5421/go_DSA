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
		l.head.prev = newNode
	} else {
		newNode.next = l.tail.next
		l.tail.next = newNode
		newNode.prev = l.tail
		// newNode.prev = l.tail
	}

	l.tail = newNode
	l.len++

}

func (l *linkedList) addfirst(value int) {
	newNode := new(node)
	newNode.value = value
	if l.len == 0 {
		newNode.next = newNode
		l.head = newNode
		l.head.prev = newNode
	} else {
		//l.head.prev = newNode
		l.tail.next = newNode
		newNode.next = l.head
		newNode.prev = l.tail
		l.head = newNode
	}
	l.tail = newNode
	l.len++
}

func (l *linkedList) remove(value int) {
	cur := l.head
	for cur != l.head {
		if cur.value == value {
			cur.next = cur.next.next
			cur.prev = cur.next.prev
			l.len--
		}
		cur = cur.next
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
		p.prev = p
		l.len++
	}
}

func (l *linkedList) removeFirst() {
	if l.len == 0 {
		fmt.Println("list is empty")
	}

	//e = l.head.value
	l.tail.next = l.head.next
	l.head = l.head.next
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}

	//return e
}

func (l *linkedList) removeany(pos int) {
	p := l.head
	i := 0
	for i < pos-1 {
		p = p.next
		i++
		//e = p.next.value
	}
	p.next = p.next.next
	p.next.prev = p
	l.len--

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
	for i := 0; i < l.len; i++ {
		fmt.Printf("-->%d", p.prev.value)
		p = p.prev
	}

}

func (l linkedList) display() {
	p := l.head
	for i := 0; i < l.len; i++ {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Printf("\n")
}

// func (l linkedList) String() string {
// 	sb := strings.Builder{}

// 	for iteration := l.head; iteration != nil; iteration = iteration.next {
// 		sb.WriteString(fmt.Sprintf("%d ", iteration.value))
// 	}
// 	return sb.String()
// }

func main() {
	l := linkedList{}
	l.addLast(34)
	l.addLast(15)
	l.addLast(20)
	l.display()
	//fmt.Println(l)
	//l.displayrev()
	l.addfirst(50)
	l.display()
	l.addany(70, 2)
	l.display()
	l.removeany(3)
	l.display()
	l.removeFirst()
	fmt.Println(l)
	l.remove(15)
	fmt.Println(l)
	l.removeany(1)
	fmt.Println(l)
	l.addany(70, 2)
	fmt.Println(l)
}
