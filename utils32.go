// +build 386 amd64p32 arm armbe mips mipsle mips64p32 mips64p32le ppc s390 sparc

package dither

//bitInterleave interleaves an unsigned integer with zero
func bitInterleave(n uint) uint {
	n = (n | n<<16) & 0x0000FFFF
	n = (n | n<<8) & 0x00FF00FF
	n = (n | n<<4) & 0x0F0F0F0F
	n = (n | n<<2) & 0x33333333
	n = (n | n<<1) & 0x55555555
	return n
}

//bitReverse reverses bits in an unsigned integer
func bitReverse(n uint) uint {
	n = (n&0xAAAAAAAA)>>1 | (n&0x55555555)<<1
	n = (n&0xCCCCCCCC)>>2 | (n&0x33333333)<<2
	n = (n&0xF0F0F0F0)>>4 | (n&0x0F0F0F0F)<<4
	n = (n&0xFF00FF00)>>8 | (n&0x00FF00FF)<<8
	n = n>>16 | n<<16
	return n
}

//reverseInterleave interleave the bits from two integers in reverse
func reverseInterleave(a, b, bc uint) uint {
	return bitReverse(bitInterleave(a)|bitInterleave(b)<<1) >> (32 - bc*2)
}
