package main

import (
	"fmt"
)

type Node struct {
	element int
	next    *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type hashChain struct {
	hashTableSize int
	hashTable     []*LinkedList
}

func (l *LinkedList) insertSorted(e int) {
	new := &Node{e, nil}
	if l.size == 0 {
		l.head = new
	} else {
		p := l.head
		q := l.head

		for p != nil && p.element < e {
			q = p
			p = p.next
		}

		if p == l.head {
			new.next = l.head
			l.head = new
		} else {
			new.next = q.next
			q.next = new
		}
	}
	l.size++
}

func (l *LinkedList) search(key int) bool {
	if l.size != 0 {
		p := l.head
		for p != nil {
			if p.element == key {
				return true
			}
			p = p.next
		}
		return false
	}
	return false
}

func (l *LinkedList) display() {
	p := l.head
	for p != nil {
		fmt.Print(p.element, "-->")
		p = p.next
	}
	fmt.Println()
}

func constructor() *hashChain {
	H := &hashChain{}
	H.hashTableSize = 10
	H.hashTable = make([]*LinkedList, H.hashTableSize)
	for i := 0; i < H.hashTableSize; i++ {
		H.hashTable[i] = &LinkedList{}
	}
	return H
}

func (h *hashChain) hashCode(key int) int {
	return key % h.hashTableSize
}

func (h *hashChain) insert(e int) {
	i := h.hashCode(e)
	h.hashTable[i].insertSorted(e)
}

func (h *hashChain) search(key int) bool {
	i := h.hashCode(key)
	return h.hashTable[i].search(key)
}

func (h *hashChain) display() {
	for i := 0; i < h.hashTableSize; i++ {
		fmt.Print("[", i, "] ")
		h.hashTable[i].display()
	}
	fmt.Println()
}

func main() {
	H := constructor()
	H.insert(54)
	H.insert(78)
	H.insert(64)
	H.insert(92)
	H.insert(34)
	H.insert(86)
	H.insert(28)
	H.display()
	fmt.Println(H.search(54))
}
