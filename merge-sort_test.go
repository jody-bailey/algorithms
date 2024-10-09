package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{2, 4, 5, 7, 1, 2, 3, 6}
	got := MergeSort(a)
	want := []int{1, 2, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MergeSort([]int{2, 4, 5, 7, 1, 2, 3, 6})
	}
}
