package dither

const uintSize = 32 << (^uint(0) >> 32 & 1)

//bitInterleave performs a bit interleave with zeroes
func bitInterleave(n uint) uint {
	//64 bits
	if uintSize == 64 {
		n = (n ^ (n << 32)) & 0x00000000ffffffff
	}
	//32 bits
	n = (n ^ (n << 16)) & 0x0000ffff0000ffff
	//16 bits
	n = (n ^ (n << 8)) & 0x00ff00ff00ff00ff
	//8 bits
	n = (n ^ (n << 4)) & 0x0f0f0f0f0f0f0f0f
	//4 bits
	n = (n ^ (n << 2)) & 0x3333333333333333
	//2 bits
	n = (n ^ (n << 1)) & 0x5555555555555555
	return n
}

//bitReverse reverse l bits in the integer n
func bitReverse(n, l uint) (o uint) {
	//Swap two bits per pass, starting with the most outwards bits (least/most significant) bits, then move inwards to the center.
	for i := uint(0); i < (l>>1)+(l&1); i++ {
		//Swap least significant bit with most significant bit
		o = o | (n&(1<<i))<<(l-i<<1-1)
		//Swap most significant bit with least significant bit
		o = o | (n&(1<<(l-i-1)))>>(l-i<<1-1)
	}
	return o
}
