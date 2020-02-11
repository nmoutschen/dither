package dither

const uintSize = 32 << (^uint(0) >> 32 & 1)

//reverseInterleave interleave the bits from two integers in reverse
func reverseInterleave(a, b, bc uint) (o uint) {
	for i := uint(0); i < bc; i++ {
		o = o | ((a>>i)&1)<<(bc*2-i*2-1)
		o = o | ((b>>i)&1)<<(bc*2-i*2-2)
	}
	return o
}
