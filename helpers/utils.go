package helpers

import "sort"

// IsSorted checks if the stack is sorted in ascending order. It iterates through the stack's data and compares each element with the next one. If it finds any element that is greater than the following element, it returns false, indicating that the stack is not sorted. If it successfully iterates through the entire stack without finding any such cases, it returns true, confirming that the stack is sorted.

func IsSorted(a *Stack) bool {
	for i := 0; i < a.Len()-1; i++ {
		if a.Data[i] > a.Data[i+1] {
			return false
		}
	}
	return true
}

// Normalize takes a slice of integers and returns a new slice where each integer is replaced by its index in the sorted order of the original slice. This is useful for the radix sort algorithm, as it allows us to work with a smaller range of integers. The function first creates a copy of the input slice and sorts it. Then, it builds a map that associates each integer with its index in the sorted slice. Finally, it constructs the result slice by replacing each integer in the original slice with its corresponding index from the map.

func Normalize(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	indexMap := make(map[int]int)
	for i, v := range sorted {
		indexMap[v] = i
	}

	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = indexMap[v]
	}

	return result
}
