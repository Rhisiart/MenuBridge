package types

type Encoded interface {
	Encode(data []byte, idx int, seq byte) (int, error)
	Type() byte
}
