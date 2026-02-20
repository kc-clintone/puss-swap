package helpers

// Rra reverse-rotates a down by 1.
func Rra(a *Stack) {
	if a.Len() < 2 {
		return
	}
	last := a.Data[a.Len()-1]
	a.Data = append([]int{last}, a.Data[:a.Len()-1]...)
}

// Rrb reverse-rotates b down by 1.
func Rrb(b *Stack) {
	if b.Len() < 2 {
		return
	}
	last := b.Data[b.Len()-1]
	b.Data = append([]int{last}, b.Data[:b.Len()-1]...)
}

// Rrr executes Rra and Rrb.
func Rrr(a, b *Stack) {
	Rra(a)
	Rrb(b)
}
