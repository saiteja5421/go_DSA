package main 


import (
	"fmt"
)

type hashLinearProbe struct {
	hashTableSize int
	hashTable     []int
}

func constructor() *hashLinearProbe {
	H := &hashLinearProbe{}
	H.hashTableSize = 10
	H.hashTable = make([]int, H.hashTableSize)
	return H
}

func (h *hashLinearProbe) hashCode(key int) int {
	return key % h.hashTableSize
}

func (h *hashLinearProbe) lprobe(key int) int {
	i := h.hashCode(key)
	j := 0

	for h.hashTable[(i+j)%h.hashTableSize] != 0 {
		j++
	}
	return (i + j) % h.hashTableSize
}

func (h *hashLinearProbe) insert(e int) {
	i := h.hashCode(e)
	if h.hashTable[i] == 0 {
		h.hashTable[i] = e
	} else {
		i = h.lprobe(e)
		h.hashTable[i] = e
	}
}

func (h *hashLinearProbe) search(key int) bool {
	i := h.hashCode(key)
	j := 0

	for h.hashTable[(i+j)%h.hashTableSize] != key {
		if h.hashTable[(i+j)%h.hashTableSize] == 0 {
			return false
		}
		j++
	}
	if i+j > h.hashTableSize {
		return false
	}
	return true
}

func (h *hashLinearProbe) display() {
	fmt.Printf("%#v\n", h.hashTable)
}

func main() {
	H := constructor()
	H.insert(54)
	H.insert(78)
	H.insert(64)
	H.insert(92)
	H.insert(34)
	H.insert(28)
	H.insert(86)
	H.display()
	fmt.Printf("%t\n", H.search(54))

}