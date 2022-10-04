package main

import "fmt"

func main() {

	// Array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// Slice

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// Slice with make
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2, len(slice2), cap(slice2))

	// Slice with append
	slice3 := append(slice2, 1, 2, 3, 4, 5)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// Slice with append and make
	slice4 := make([]int, 0, 5)
	slice4 = append(slice4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// Slice with append  list
	new_slice := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	slice5 := append(new_slice, slice4...)
	fmt.Println(slice5, len(slice5), cap(slice5))

}
