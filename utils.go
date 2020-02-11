package dither

const seed uint = 0x7FFFFFFF

func xorshift(n uint) uint {
	n ^= (n << 13) & 0xFFFFFFFF
	n ^= (n >> 17) & 0xFFFFFFFF
	n ^= (n << 5) & 0xFFFFFFFF
	return n
}
