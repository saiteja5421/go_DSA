package main

import (
	"fmt"
	"math"
	"strconv"
)

func max(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
func sort(a []int) []int {
	n := len(a)
	maxnumber := max(a)
	//fmt.Println(maxnumber)
	maxnum := strconv.Itoa(maxnumber)
	digits := len(maxnum)
	//fmt.Println(digits)
	//l:=[]int{}
	bins := make([][]int, 10)
	for i := 0; i < digits; i++ {
		for j := 0; j < n; j++ {
			e := int((float64(a[j]) / math.Pow(10, float64(i)))) % 10
			bins[e] = append(bins[e], a[j])
			//fmt.Println(e)
			// if len(bins[e]) > 0 {
			// 	bins[e] = append(bins[e], a[j])
			// } else {
			// 	bins[e] = []int{a[j]}
			// }
		}
		//fmt.Println(bins)

		k := 0
		for x := 0; x < 10; x++ {
			c := len(bins[x])
			//fmt.Println(c)
			if c > 0 {
				for y := 0; y < c; y++ {
					a[k] = bins[x][0]
					bins[x] = bins[x][1:]
					//fmt.Println(len(bins[x]))
					k = k + 1
				}
			}

		}

	}
	return a
}

func main() {
	b :=[]int{64, 250, 775, 947, 771, 28}
	a := sort(b)
	fmt.Println(a)
}
