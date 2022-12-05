package main

import "fmt"

// func (l *s)len() int {
// 	return len(l.a)
// }

type s struct {
	a   []int
	len int
}

func (l *s) push(e int) {
	l.a = append(l.a, e)
	l.len++

}

func (l *s) pop() {
	if l.isempty() {
		fmt.Println("slice is empty")
	} else {
		l.a = l.a[:len(l.a)-1]
		l.len--
	}
}
func (l *s) top() int {
	if l.isempty() {
		fmt.Println("slice is empty")
		return 0
	} else {
		return l.a[l.len-1]
	}
}
func (l *s) isempty() bool {
	return l.len == 0
}
func main() {
	// var s = []int{1, 5, 7}
	// a := len(s)
	// fmt.Println(a)

	var sai s
	sai.push(5)
	sai.push(10)
	sai.push(20)
	fmt.Println(sai.a)
	sai.pop()
	fmt.Println(sai.a)
	//b := sai.top()
	//fmt.Println(b)
	sai.pop()
	sai.pop()
	//sai.top()
	a := sai.isempty()
	fmt.Println(a)
	sai.push(70)
	b := sai.isempty()
	fmt.Println(b)
}
