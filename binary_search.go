package algorithms

// BinarySearch returns the index position of the element x if found in the list.
// If x is not found, the value -1 is returned.
func BinarySearch(list []int, x int) int {
	return binarySearch(list, x, 0, len(list)-1)
}

func binarySearch(list []int, x, left, right int) int {
	m := (left + right) / 2

	if list[m] == x {
		return m
	} else if left == right {
		return -1 // Not found
	} else if x < m {
		return binarySearch(list, x, left, m-1)
	} else {
		return binarySearch(list, x, m+1, right)
	}
}
