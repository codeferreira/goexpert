package matematica

type Numeric interface {
	int | float64
}

func Soma[Type Numeric](a, b Type) Type {
	return a + b
}
