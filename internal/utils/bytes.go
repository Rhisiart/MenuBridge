package utils

func Write16(buf []byte, offset int, value int) {
	hi := (value & 0xFF00) >> 8
	lo := value & 0xFF
	buf[offset] = byte(hi)
	buf[offset+1] = byte(lo)
}
