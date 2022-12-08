package main

import "fmt"

func mergeSorted(a []int, left int, right int) {
	if left < right {
		mid := (left+right)/2 
		mergeSorted(a, left, mid)
		mergeSorted(a, mid+1, right)
		merge(a, left, mid, right)
	}
}

func merge(a []int, left, mid, right int)  {
	i := left
	j := mid + 1
	k := left
	b := make([]int,right+1)
	for i <= mid && j <= right {
		if a[i] < a[j] {
			b[k] = a[i]
			i = i + 1
			k = k + 1
		} else {
			b[k] = a[j]
			j = j + 1
			k = k + 1
		}
	}
	for i <= mid {
		b[k] = a[i]
		i += 1
		k += 1
	}
	for j <= right {
		b[k] = a[j]
		j += 1
		k += 1
	}
	for x := left; x <= right; x++ {
		a[x] = b[x]
	}

}

func main() {
	d := []int{5, 2, 1, 4, 3}
	mergeSorted(d, 0, len(d)-1)
	fmt.Println(d)
}

// func Merge(a1 []int, a2 []int) []int {

// 	var r = make([]int, len(a1)+len(a2))
// 	var i = 0
// 	var j = 0

// 	for i < len(a1) && j < len(a2) {

// 		if a1[i] <= a2[j] {
// 			r[i+j] = a1[i]
// 			i++
// 		} else {
// 			r[i+j] = a2[j]
// 			j++
// 		}

// 	}

// 	for i < len(a1) {
// 		r[i+j] = a1[i]
// 		i++
// 	}
// 	for j < len(a2) {
// 		r[i+j] = a2[j]
// 		j++
// 	}

// 	return r

// }

// func Merge_Sort(items []int) []int {

// 	if len(items) < 2 {
// 		return items

// 	}

// 	var middle = len(items) / 2
// 	var a = Merge_Sort(items[:middle])
// 	var b = Merge_Sort(items[middle:])
// 	return Merge(a, b)

// }

// func main() {

// 	fmt.Print(Merge_Sort([]int{10, 92, 83, 43, 45, 64, 744, 33, 22, 14, 2344, 56743}), "\n")

// }
