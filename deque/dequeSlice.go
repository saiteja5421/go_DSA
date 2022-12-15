package main

import "fmt"

type dequeue struct {
	data []int
	len  int
}

func (l *dequeue) addLast(e int) {
	l.data = append(l.data, e)
	l.len++
}
func (l *dequeue) addFirst(e int) {
	d := []int{e}
	d = append(d, l.data...)
	l.data = d
	l.len++
}
func (l *dequeue) removeLast() {
	if l.isempty() {
		fmt.Println("queue is empty")
		return
	} else {
		l.data = l.data[:len(l.data)-1]
		l.len--
	}
}
func (l *dequeue) removeFirst() {
	if l.isempty() {
		fmt.Println("queue is empty")
		return
	} else {
		l.data = l.data[1:]
		l.len--
	}
}
func (l *dequeue) isempty() bool {
	return l.len == 0
}

func (l *dequeue) firstElement() int {
	if l.isempty() {
		fmt.Println("slice is empty")
		return 0
	} else {
		return l.data[0]
	}
}
func (l *dequeue) lastElement() int {
	if l.isempty() {
		fmt.Println("slice is empty")
		return 0
	} else {
		return l.data[len(l.data)-1]
	}
}
func main() {
	var s dequeue
	s.addLast(20)
	s.addLast(30)
	s.addFirst(50)
	s.addFirst(70)
	fmt.Println(s.data)
	s.removeLast()
	s.removeFirst()
	fmt.Println(s.data)

}
