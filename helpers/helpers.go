package helpers

func PointerOf[A any](a A) *A {
	return &a
}
