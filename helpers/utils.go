package helpers

import "sort"

func IsSorted(a *Stack) bool {
	for i := 0; i < a.Len()-1; i++ {
		if a.Data[i] > a.Data[i+1] {
			return false
		}
	}
	return true
}

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
