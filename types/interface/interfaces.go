package types

type Table interface {
	MarshalBinary() []byte
	UnmarshalBinary(data []byte) error
}
