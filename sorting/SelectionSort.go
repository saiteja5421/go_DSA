package main

import "fmt"

func Sort(a []int) []int {

	n := len(a)
	for i := 0; i < n-1; i++ {
		position := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[position] {
				position = j
			}

		}
		temp := a[i]
		a[i] = a[position]
		a[position] = temp

	}

	return a
}

func main() {
	b := Sort([]int{2, 5, 4, 1, 3})
	fmt.Println(b)
}
