package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

var ErrIndexOutOfRange = errors.New("index out of range")

func main() {
	var originalSlice []int
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, 0+rand.Intn(11-0+1))
	}
	fmt.Printf("originalSlice n= %d\n", originalSlice)

	fmt.Println("\n===Test func sliceExample===")
	example := sliceExample(originalSlice, func(v int) bool { return v%2 != 0 })
	fmt.Printf("sliceExample return n= %d\n", example)

	fmt.Println("\n===Test func copySlice===")
	copyS := copySlice(originalSlice)
	fmt.Printf("Add element %d in original slice", 100)
	originalSlice = append(originalSlice, 100)
	fmt.Printf("Copy slice n=%d\n", copyS)
	fmt.Printf("Slice original: %v\n", originalSlice)
	fmt.Printf("Slice copy: %v\n", copyS)
	if len(originalSlice) != len(copyS) {
		fmt.Println("Great: Slices are not equal")
	} else {
		fmt.Println("Error: Slices are equal")
	}

	fmt.Println("\n===Test func addElements===")
	fmt.Printf("Origin slice = %d\n", originalSlice)
	element := 0 + rand.Intn(101-0+1)
	fmt.Printf("Add element = %d\n", element)
	fmt.Printf("New slice = %d\n", addElements(originalSlice, element))

	fmt.Println("\n===Test func removeElement===")
	fmt.Printf("Origin slice n = %d\n", originalSlice)
	index := 0 + rand.Intn(11-0+1)
	fmt.Printf("Deleted index = %d\n", index)
	removeElementSlice, err := removeElement(originalSlice, index)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("New slice = %d\n", removeElementSlice)
}

// sliceExample filters elements from a slice based on a predicate function and returns a new slice without modifying the original.
func sliceExample[T any](slice []T, predict func(T) bool) []T {
	s := make([]T, len(slice))
	copy(s, slice)
	n := 0

	for _, v := range s {
		if !predict(v) {
			s[n] = v
			n++
		}
	}

	var zero T
	for i := n; i < len(s); i++ {
		s[i] = zero
	}

	return s[:n]
}

// addElements appends an element of any type to a copy of the given slice and returns the new slice.
func addElements[T any](slice []T, n T) []T {
	s := make([]T, len(slice))
	copy(s, slice)

	return append(s, n)
}

// copySlice creates and returns a new slice, copying all elements from the provided slice without modifying the original.
func copySlice[T any](slice []T) []T {
	s := make([]T, len(slice))
	copy(s, slice)
	return s
}

// removeElement removes an element at the specified index n from a slice of integers and returns a new modified slice.
func removeElement(slice []int, n int) ([]int, error) {
	s := make([]int, len(slice))
	copy(s, slice)
	if n > len(s) {
		return slice, ErrIndexOutOfRange
	}
	if n == len(s) {
		return slices.Delete(s, n-1, n), nil
	}
	return slices.Delete(s, n, n+1), nil
}
