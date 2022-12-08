package main

import "fmt"

func binarysearch(a []int, key int) int {
	var l = 0
	var r = len(a) - 1

	for l <= r {
		var mid = (l + r) / 2
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			r = mid - 1
		} else if key > a[mid] {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	b := binarysearch([]int{1, 2, 3, 4, 5, 6}, 4)
	fmt.Println(b)
}