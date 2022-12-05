package main

import "fmt"

func Sort(a []int) []int {
	n := len(a)
	for i := n-1; i > 0; i = i - 1 {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	return a
}
func main() {
	b := Sort([]int{5, 3, 2, 4, 1})
	fmt.Println(b)
}
