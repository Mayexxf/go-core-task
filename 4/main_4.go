package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(equalRows(slice1, slice2))

}

func equalRows(a, b []string) []string {
	equal := make(map[string]struct{})
	result := make([]string, 0)

	for _, v := range b {
		equal[v] = struct{}{}
	}

	for _, v := range a {
		if _, ok := equal[v]; !ok {
			result = append(result, v)
		}
	}

	return result
}
