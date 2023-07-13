package bit

type Bit int

const (
	O Bit = iota
	I
)

func (b Bit) String() string {
	switch b {
	case O:
		return "O"
	case I:
		return "I"
	default:
		panic("invalid bit")
	}
}

func (b Bit) Not() Bit {
	switch b {
	case O:
		return I
	case I:
		return O
	default:
		panic("invalid Bit")
	}
}

func (b Bit) And(b2 Bit) Bit {
	switch {
	case b == O || b2 == O:
		return O
	case b == I && b2 == I:
		return I
	default:
		panic("invalid bits")
	}
}

func (b Bit) Or(b2 Bit) Bit {
	switch {
	case b == I || b2 == I:
		return I
	case b == O && b2 == O:
		return O
	default:
		panic("invalid bits")
	}
}
