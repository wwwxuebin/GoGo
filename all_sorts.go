package main

import (
	"fmt"
)

//all the code are base on the java code in algorithm(4 version)

//base merge sort
func merge_sort(a []int, lo, mid, hi int) {
	var aux [15]int

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func sort(a []int, lo, hi int) {
	if hi > lo {
		mid := lo + (hi-lo)/2

		sort(a, lo, mid)
		sort(a, mid+1, hi)

		merge_sort(a, lo, mid, hi)
	}
}

//follow method is improved merge sort

func merge_sort_insertion(src []int, dst []int, lo, mid, hi int) {
	i := lo
	j := mid + 1

	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src[j] < src[i] {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func exch(src []int, i, j int) {
	temp := src[i]
	src[i] = src[j]
	src[j] = temp
}

func insertionSort(src []int, lo, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && src[j] < src[j-1]; j-- {
			exch(src, j, j-1)
		}
	}
}

func arrayCopy(src []int, src_offset int, dst []int, dst_offset, length int) {
	for i := 0; i < length; i++ {
		dst[dst_offset+i] = src[src_offset+i]
	}
}

func sort_insertion(src []int, dst []int, lo, hi int) {
	if hi <= lo+7 {
		insertionSort(dst, lo, hi)
	} else {
		mid := lo + (hi-lo)/2

		sort_insertion(dst, src, lo, mid)
		sort_insertion(dst, src, mid+1, hi)

		if src[mid] <= src[mid+1] {
			arrayCopy(src, lo, dst, lo, hi-lo+1)
		} else {
			merge_sort_insertion(src, dst, lo, mid, hi)
		}
	}
}

func clone(src []int, dst []int) {
	array_size := len(src)

	for i := 0; i < array_size; i++ {
		dst[i] = src[i]
	}
}

func sort_insertion_base(array []int, lo, hi int) {
	var b = make([]int, 15)
	clone(array, b)

	sort_insertion(array, b, lo, hi)

	clone(b, array)
}

func main() {
	a := []int{8, 5, 9, 6, 18, 55, 47, 28, 36, 14, 89, 2, 1, 68, 12}

	fmt.Println(a)
	sort_insertion_base(a, 0, len(a)-1)

	fmt.Println(a)
}
