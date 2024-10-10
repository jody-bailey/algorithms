package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{2, 4, 5, 7, 1, 2, 3, 6}
	got := BubbleSort(a)
	want := []int{1, 2, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but expected %v", got, want)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = BubbleSort([]int{2, 4, 5, 7, 1, 2, 3, 6})
	}
}
