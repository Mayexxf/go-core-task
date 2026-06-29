package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(equalSlices(a, b))
}

func equalSlices[T comparable](a, b []T) (bool, []T) {
	set := make(map[T]struct{}, len(b))
	result := make([]T, 0)
	for _, v := range b {
		set[v] = struct{}{}
	}
	for _, v := range a {
		if _, ok := set[v]; ok {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
