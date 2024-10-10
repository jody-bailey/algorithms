package algorithms

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	t.Run("no repeat numbers", func(t *testing.T) {
		list := []int{5, 2, 4, 6, 1, 3}
		want := []int{1, 2, 3, 4, 5, 6}
		got := InsertionSort(list)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but expected %v", got, want)
		}
	})
	t.Run("contains repeat numbers", func(t *testing.T) {
		list := []int{5, 2, 4, 6, 1, 3, 2, 4, 2, 5}
		want := []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 6}
		got := InsertionSort(list)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v but expected %v", got, want)
		}
	})
}
