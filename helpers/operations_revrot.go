package helpers

func Rra(a *Stack) {
	if a.Len() < 2 {
		return
	}
	last := a.Data[a.Len()-1]
	a.Data = append([]int{last}, a.Data[:a.Len()-1]...)
}

func Rrb(b *Stack) {
	if b.Len() < 2 {
		return
	}
	last := b.Data[b.Len()-1]
	b.Data = append([]int{last}, b.Data[:b.Len()-1]...)
}

func Rrr(a, b *Stack) {
	Rra(a)
	Rrb(b)
}
