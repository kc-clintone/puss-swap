package helpers

func radixSort(a, b *Stack) []string {
	var ops []string

	a.Data = Normalize(a.Data)
	n := a.Len()
	if n <= 1 {
		return ops
	}

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

			Pb(a, b)
			ops = append(ops, "pb")

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
