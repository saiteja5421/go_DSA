package main

import "fmt"

type queue struct {
	data []int
	len  int
}

func (l *queue) push(e int) {
	l.data = append(l.data, e)
	l.len++
}

func (l *queue) pop() {
	if l.isempty() {
		fmt.Println("queue is empty")
		return
	} else {
		l.data = l.data[1:]
		l.len--
	}
}
func (l *queue) isempty() bool {
	return l.len == 0
}

func (l *queue) top() int {
	if l.isempty() {
		fmt.Println("slice is empty")
		return 0
	} else {
		return l.data[0]
	}
}
func main() {
	var s queue
	s.push(20)
	s.push(30)
	fmt.Println(s.data)
	s.pop()
	fmt.Println(s.data)

}
