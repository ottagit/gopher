package main

func MakeRange(min, max int) []int {
	a := make([]int, max)
	for i := range a {
		a[i] = min + i
	}
	return a
}
