package main

import "fmt"

func linearsearch(a []int, key int) int {
	for i, v := range a {
		if key == v {
			return i
		}

	}
	return -1
}

func main() {
	b := linearsearch([]int{45, 2, 76, 345, 23}, 345)
	fmt.Println(b)
}
