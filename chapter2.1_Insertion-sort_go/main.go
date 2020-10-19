package main

import "fmt"

func main() {

	A := []int{31, 41, 59, 26, 41, 58}
	B := []int{31, 41, 59, 26, 41, 58}
	Insertion_sort(A)
	Insertion_sort_desc(B)
	fmt.Println(A)
	fmt.Println(B)
}

func Insertion_sort(A []int) {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for ; i >= 0 && A[i] > key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
}

func Insertion_sort_desc(A []int) {
	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for ; i >= 0 && A[i] < key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
}