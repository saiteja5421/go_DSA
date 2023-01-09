package main

import "fmt"

func max(a []int) int {
	b := a[0]
	for _, i := range a {
		if i > b {
			b = i
		}
	}
	return b
}
func insertionSort(a []int) {
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
	
}

func bucketSort(A []int) {
	n:=len(A)
	m := max(A)
	var buckets [10][]int
	for i:=0;i<n;i++{
        j:= n * A[i]/(m+1)
		if len(buckets[j])==0{
		   buckets[j]=[]int{A[i]}
		}else{
			buckets[j]=append(buckets[j], A[i])
		}
	}
	//fmt.Println(buckets)
	for i:=0;i<10;i++{
		insertionSort(buckets[i])
	}
    //fmt.Println(buckets)
	k:=0
	for i:=0;i<10;i++{
		c:=len(buckets[i])
		if c>0{
		for j:=0;j<c;j++{
			A[k]=buckets[i][0]
			//fmt.Println(A[k])
			//fmt.Println(len(buckets[i]))
			buckets[i]=buckets[i][1:]
			//fmt.Println(len(buckets[i]))
			k=k+1
			
		}
	}
	} 
	//fmt.Println(A)
}

func main() {
	a := []int{63,250,835,947,651,28}
	bucketSort(a)
    fmt.Println(a)
}
