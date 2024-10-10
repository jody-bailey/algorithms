package algorithms

// MergeSort returns a sorted Go slice representing list
//
// The sorting operation is done by following a divide-and-conquer
// approach. It begins by recursively breaking
// the slice in half on each iteration until there is only
// a single element remaining. It then performs a merge
// operation, returning a new sorted slice to the parent,
// resulting in a fully sorted slice by the end of the
// recursion. This algorithm is faster than insertion sort
// but comes at the cost of a larger memory footprint.
func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	midpoint := len(list) / 2
	left := MergeSort(list[:midpoint])
	right := MergeSort(list[midpoint:])

	merged := merge(left, right)
	return merged
}

func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
