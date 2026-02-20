package helpers

// radixSort implements the radix sort algorithm replacement: chunk-based algorithm.
func radixSort(a, b *Stack) []string {
	var ops []string

	// normalize values to ranks 0..n-1
	a.Data = Normalize(a.Data)
	n := a.Len()
	if n <= 1 {
		return ops
	}

	// chunk size tuned for 100 elements; adjust as needed for larger inputs
	chunkSize := 20
	if n > 100 && n <= 500 {
		chunkSize = 45
	} else if n > 500 {
		chunkSize = 80
	}

	curr := 0
	for curr < n {
		limit := curr + chunkSize - 1
		if limit >= n {
			limit = n - 1
		}

		// push all elements in [curr..limit] from a to b
		for {
			idx := -1
			for i, v := range a.Data {
				if v >= curr && v <= limit {
					idx = i
					break
				}
			}
			if idx == -1 {
				break
			}

			// bring the element to top with minimal rotations
			if idx <= a.Len()/2 {
				for i := 0; i < idx; i++ {
					Ra(a)
					ops = append(ops, "ra")
				}
			} else {
				for i := 0; i < a.Len()-idx; i++ {
					Rra(a)
					ops = append(ops, "rra")
				}
			}

			// push to b
			Pb(a, b)
			ops = append(ops, "pb")

			// keep larger elements near top of b to reduce rotations when pushing back
			if b.Len() > 1 {
				threshold := curr + (chunkSize / 2)
				if b.Data[0] < threshold {
					Rb(b)
					ops = append(ops, "rb")
				}
			}
		}

		curr = limit + 1
	}

	// push back from b to a in descending order (largest first)
	for !b.IsEmpty() {
		maxIdx := findMaxIndex(b)
		if maxIdx <= b.Len()/2 {
			for i := 0; i < maxIdx; i++ {
				Rb(b)
				ops = append(ops, "rb")
			}
		} else {
			for i := 0; i < b.Len()-maxIdx; i++ {
				Rrb(b)
				ops = append(ops, "rrb")
			}
		}
		Pa(a, b)
		ops = append(ops, "pa")
	}

	return ops
}

// findMaxIndex returns the index of the largest element in the stack.
func findMaxIndex(s *Stack) int {
	if s.Len() == 0 {
		return -1
	}
	max := s.Data[0]
	idx := 0
	for i, v := range s.Data {
		if v > max {
			max = v
			idx = i
		}
	}
	return idx
}
