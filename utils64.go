// +build amd64 arm64 arm64be ppc64 ppc64le mips64 mips64le s390x sparc64

package dither

const uintSize = 32 << (^uint(0) >> 32 & 1)

//bitInterleave interleaves an unsigned integer with zero
func bitInterleave(n uint) uint {
	if uintSize == 64 {
		n = (n | n<<32) & 0x00000000_FFFFFFFF
		n = (n | n<<16) & 0x0000FFFF_0000FFFF
		n = (n | n<<8) & 0x00FF00FF_00FF00FF
		n = (n | n<<4) & 0x0F0F0F0F_0F0F0F0F
		n = (n | n<<2) & 0x33333333_33333333
		n = (n | n<<1) & 0x55555555_55555555
	} else {
		n = (n | n<<16) & 0x0000FFFF
		n = (n | n<<8) & 0x00FF00FF
		n = (n | n<<4) & 0x0F0F0F0F
		n = (n | n<<2) & 0x33333333
		n = (n | n<<1) & 0x55555555
	}
	return n
}

//bitReverse reverses bits in an unsigned integer
func bitReverse(n uint) uint {
	if uintSize == 64 {
		n = (n&0xAAAAAAAA_AAAAAAAA)>>1 | (n&0x55555555_55555555)<<1
		n = (n&0xCCCCCCCC_CCCCCCCC)>>2 | (n&0x33333333_33333333)<<2
		n = (n&0xF0F0F0F0_F0F0F0F0)>>4 | (n&0x0F0F0F0F_0F0F0F0F)<<4
		n = (n&0xFF00FF00_FF00FF00)>>8 | (n&0x00FF00FF_00FF00FF)<<8
		n = (n&0xFFFF0000_FFFF0000)>>16 | (n&0x0000FFFF_0000FFFF)<<16
		n = n>>32 | n<<32
	} else {
		n = (n&0xAAAAAAAA)>>1 | (n&0x55555555)<<1
		n = (n&0xCCCCCCCC)>>2 | (n&0x33333333)<<2
		n = (n&0xF0F0F0F0)>>4 | (n&0x0F0F0F0F)<<4
		n = (n&0xFF00FF00)>>8 | (n&0x00FF00FF)<<8
		n = n>>16 | n<<16
	}
	return n
}

//reverseInterleave interleave the bits from two integers in reverse
func reverseInterleave(a, b, bc uint) uint {
	return bitReverse(bitInterleave(a)|bitInterleave(b)<<1) >> (64 - bc*2)
}
