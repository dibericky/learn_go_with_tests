package main

import "testing"

func TestSum(t *testing.T) {

	assertSum := func(t *testing.T, want, got int, numbers [5]int) {
		t.Helper()
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertSum(t, want, got, numbers)
	})
}
