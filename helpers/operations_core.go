package helpers

// Sa swaps first two elements of a.
func Sa(a *Stack) {
	if a.Len() < 2 {
		return
	}
	a.Data[0], a.Data[1] = a.Data[1], a.Data[0]
}

func Sb(b *Stack) {
	if b.Len() < 2 {
		return
	}
	b.Data[0], b.Data[1] = b.Data[1], b.Data[0]
}

func Ss(a, b *Stack) {
	Sa(a)
	Sb(b)
}

func Pa(a, b *Stack) {
	if b.IsEmpty() {
		return
	}
	val := b.Data[0]
	b.Data = b.Data[1:]
	a.Data = append([]int{val}, a.Data...)
}

func Pb(a, b *Stack) {
	if a.IsEmpty() {
		return
	}
	val := a.Data[0]
	a.Data = a.Data[1:]
	b.Data = append([]int{val}, b.Data...)
}
