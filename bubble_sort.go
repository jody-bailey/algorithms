package algorithms

// BubbleSort returns a sorted Go slice representing list
//
// The sorting operation is done by repeatedly swapping adjacent
// elements that are out of order. This is an inefficient sorting
// algorithm.
func BubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j] < list[i] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
