package types

type Table interface {
	MarshalBinary() []byte
	UnmarshalBinary(data []byte) error
}

type Encoded interface {
	Encode(data []byte, idx int, seq byte) (int, error)
	Type() byte
}
