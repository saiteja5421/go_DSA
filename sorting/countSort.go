package main

import "fmt"

func max(a []int) int {
	max_num := a[0]
	for _, v := range a {
		if v > max_num {
			max_num = v
		}
	}
	return max_num
}

func sort(a []int) []int {
	n := len(a)
	maxsize := max(a)
	b := make([]int, maxsize+1)
	for i := 0; i < n; i++ {
		b[a[i]] = b[a[i]] + 1
	}
	i := 0 
	j := 0 
	for i <maxsize+1{
		if b[i] > 0{
			a[j]=i 
			j=j+1
			b[i]=b[i]-1
		}else{
			i=i+1
		}
	}
	return a
}

func main() {
	d := sort([]int{5, 2, 0, 3, 1})
	fmt.Println(d)

}
