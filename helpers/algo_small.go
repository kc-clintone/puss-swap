package helpers

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

func sortUpToSix(a, b *Stack) []string {
	var ops []string
	n := a.Len()
	toPush := n - 3

	for k := 0; k < toPush; k++ {
		idx := findMinIndex(a)
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

	ops = append(ops, sortThree(a)...)

	for !b.IsEmpty() {
		v := b.Data[0]
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

		Pa(a, b)
		ops = append(ops, "pa")
	}

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
