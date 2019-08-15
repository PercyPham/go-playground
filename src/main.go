package main

import "fmt"

func main() {
	slice1 := []byte{0, 1, 2, 3}
	slice2 := []byte{4, 5, 6}

	bigSlice := append(slice1, slice2...)
	bigSlice[2] = 123

	printSlice(slice1)
	printSlice(slice2)
	printSlice(bigSlice)
}

func printSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
