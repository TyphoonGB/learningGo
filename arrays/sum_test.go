package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 3, 3, 7})
		want := []int{5, 13}
		checkSums(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 3, 3, 7})
		want := []int{0, 13}
		checkSums(t, got, want)
	})

}
