package helpers

// sortThree sorts a stack of three elements and returns the operations performed.
func sortThree(a *Stack) []string {
	var ops []string

	x := a.Data[0]
	y := a.Data[1]
	z := a.Data[2]

	if x > y && y < z && x < z {
		Sa(a)
		ops = append(ops, "sa")
	} else if x > y && y > z {
		Sa(a)
		Rra(a)
		ops = append(ops, "sa", "rra")
	} else if x > y && y < z && x > z {
		Ra(a)
		ops = append(ops, "ra")
	} else if x < y && y > z && x < z {
		Sa(a)
		Ra(a)
		ops = append(ops, "sa", "ra")
	} else if x < y && y > z && x > z {
		Rra(a)
		ops = append(ops, "rra")
	}

	return ops
}

// findMinIndex returns the index of the smallest element in the stack.
func findMinIndex(a *Stack) int {
	if a.Len() == 0 {
		return -1
	}
	min := a.Data[0]
	idx := 0
	for i, v := range a.Data {
		if v < min {
			min = v
			idx = i
		}
	}
	return idx
}

// sortUpToSix handles stacks of size 4..6 by pushing the smallest n-3 elements to b,
// sorting the remaining 3 in a, then inserting back each element from b into the
// correct place in a with minimal rotations.
func sortUpToSix(a, b *Stack) []string {
	var ops []string
	n := a.Len()
	toPush := n - 3

	// push the smallest `toPush` elements to b
	for k := 0; k < toPush; k++ {
		idx := findMinIndex(a)
		// bring min to top
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
	}

	// now a has 3 elements -> sort them with the minimal routine
	ops = append(ops, sortThree(a)...)

	// insert each element from b into correct place in sorted a
	for !b.IsEmpty() {
		v := b.Data[0] // top of b
		// find insertion index in a (first element > v), if none -> idx = 0 (insert after last)
		idx := -1
		for i, val := range a.Data {
			if val > v {
				idx = i
				break
			}
		}
		if idx == -1 {
			idx = 0
		}

		// rotate a so that a[idx] becomes top (minimal rotations)
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

		// push from b to a
		Pa(a, b)
		ops = append(ops, "pa")
	}

	// finally rotate a so the smallest element is at the top
	minIdx := findMinIndex(a)
	if minIdx <= a.Len()/2 {
		for i := 0; i < minIdx; i++ {
			Ra(a)
			ops = append(ops, "ra")
		}
	} else {
		for i := 0; i < a.Len()-minIdx; i++ {
			Rra(a)
			ops = append(ops, "rra")
		}
	}

	return ops
}
