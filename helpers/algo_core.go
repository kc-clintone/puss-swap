package helpers

// Sort takes two stacks and returns a slice of operations to sort stack a using stack b as auxiliary.
func Sort(a, b *Stack) []string {
	var ops []string

	if IsSorted(a) || a.Len() == 0 {
		return ops
	}

	if a.Len() == 2 {
		if a.Data[0] > a.Data[1] {
			Sa(a)
			ops = append(ops, "sa")
		}
		return ops
	}

	if a.Len() == 3 {
		return sortThree(a)
	}

	// optimize for small sizes (4,5,6) to reduce number of operations
	if a.Len() >= 4 && a.Len() <= 6 {
		return sortUpToSix(a, b)
	}

	return radixSort(a, b)
}
