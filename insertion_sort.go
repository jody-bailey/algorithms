// Package algorithms implements various algorithms from the
// book "Introduction to Algorithms, 3rd Edition".
//
// The algorithms package should NOT be used in place of
// the Go standard library as its main purpose is for learning.
package algorithms

// InsertionSort returns a sorted Go slice representing list
//
// The sorting operation is done in-place, thus not requiring
// any additional memory usage. This algorithm is most efficient
// for smaller datasets.
func InsertionSort(list []int) []int {
	for j := 1; j < len(list); j++ {
		key := list[j]
		i := j - 1
		for i >= 0 && list[i] > key {
			list[i+1] = list[i]
			i = i - 1
		}
		list[i+1] = key
	}
	return list
}
