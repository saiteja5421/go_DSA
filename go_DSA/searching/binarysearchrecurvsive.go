package main

import "fmt"

func binaryrecurvise(a []int, key, l, r int) int {

	if l > r {
		return -1
	} else {
		var mid = (l + r) / 2
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			return binaryrecurvise(a, key, l, mid-1)
		} else if key > a[mid] {
			return binaryrecurvise(a, key, mid+1, r)
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 7}
	b := binaryrecurvise(a, 7, 0, len(a)-1)
	fmt.Println(b)
}
