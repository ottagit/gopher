package main

import (
	"reflect"
	"testing"

	sum "github.com/ottagit/goer/array-slices"
)

func TestMakeRange(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	checkSum := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("get the sum of a collection of first 8 non-zero numbers", func(t *testing.T) {
		got := sum.Sum(numbers)
		want := 10
		checkSum(t, got, want)
	})

	t.Run("collection of 8 numbers", func(t *testing.T) {

		got := MakeRange(1, 5)
		want := numbers

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
