package main

import "fmt"

type heap struct {
	data  []int
	csize int
	hi    int
}

func (t *heap) length() int {
	return len(t.data)
}

func (t *heap) isempty() bool {
	return len(t.data) == 0
}

func (t *heap) insert(e int) {
	if t.csize == 10 {
		fmt.Println("No space")
		return
	}
	t.csize = t.csize + 1
	t.hi = t.csize
	for t.hi > 1 && e > t.data[t.hi/2] {
		t.data[t.hi] = t.data[t.hi/2]
		t.hi = t.hi / 2
	}
	t.data[t.hi] = e
}

func (t *heap) deleteMax() int {
	if t.csize == 0 {
		fmt.Println("Heap is empty")
		return 0
	}
	e := t.data[1]
	t.data[1] = t.data[t.csize]
	t.data[t.csize] = 0
	t.csize = t.csize - 1
	i := 1
	j := i * 2
	for j <= t.csize {
		if t.data[j] < t.data[j+1] {
			j++
		}
		if t.data[i] < t.data[j] {
			temp := t.data[i]
			t.data[i] = t.data[j]
			t.data[j] = temp
			i = j
			j = i * 2
		} else {
			break
		}

	}
	return e
}

func  heapSort(A []int) {
	H:= heap{}
	n := len(A)
	for i := 0; i < n; i++ {
		H.insert(A[i])
	}
	k := n - 1
	for j := 0; j < n; j++ {
		A[k] = H.deleteMax()
		k = k - 1
	}
}

func main() {
	h := heap{}
	h.insert(25)
	h.insert(14)
	h.insert(2)
	h.insert(20)
	h.insert(10)
	h.insert(40)
	fmt.Println(h.data)
	fmt.Println(h.csize)
	h.deleteMax()
	fmt.Println(h.data)
	fmt.Println(h.csize)
	h.deleteMax()
	fmt.Println(h.data)
	fmt.Println(h.csize)
	a:=[]int{345,45,86,654,28}
	heapSort(a)
	fmt.Println(a)
}
