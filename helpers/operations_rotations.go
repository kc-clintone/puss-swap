package helpers

// Ra rotates a up by 1.
func Ra(a *Stack) {
	if a.Len() < 2 {
		return
	}
	first := a.Data[0]
	a.Data = append(a.Data[1:], first)
}

// Rb rotates b up by 1.
func Rb(b *Stack) {
	if b.Len() < 2 {
		return
	}
	first := b.Data[0]
	b.Data = append(b.Data[1:], first)
}

// Rr executes Ra and Rb.
func Rr(a, b *Stack) {
	Ra(a)
	Rb(b)
}
