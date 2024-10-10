package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("short list", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		got := BinarySearch(a, 3)
		want := 2

		if got != want {
			t.Errorf("got %v but expected %v", got, want)
		}
	})
	t.Run("longer list", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
		got := BinarySearch(a, 20)
		want := 19

		if got != want {
			t.Errorf("got %v but expected %v", got, want)
		}
	})
}
