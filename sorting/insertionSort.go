package main

import "fmt"

func Sort(a []int) []int {
	n := len(a)
	for i := 0; i < n; i++ {
		cvalue := a[i]
		position := i
		for position > 0 && a[position-1] > cvalue {
			a[position] = a[position-1]
			position = position - 1
		}
		a[position] = cvalue
	}
	return a
}

func main() {
	b := Sort([]int{5, 3, 2, 4, 1})
	fmt.Println(b)
	// var t , i uint
    // t , i = 1 , 1

    // for i = 1 ; i < 10 ; i++ {
    //     fmt.Printf("%d << %d = %d \n", t , i , t<<i)
    // }


    // fmt.Println()

    // t = 512
    // for i = 1 ; i < 10 ; i++ {
    //     fmt.Printf("%d >> %d = %d \n", t , i , t>>i)
	// }
}
