package dither

//seed is the default seed used by NewRandom
const seed uint = 0x7FFFFFFF

//xorshift is a xorshift32 implementation
//
//This performs an AND 2^32-1 operation after each operation to ensures that the function returns the same result on 32- and 64-bits
//systems without relying on uint32 conversions.
//
//Xorshift generators are not cryptographically secure PRNG, but are amongst the fastest generators, allowing fast computation for very
//large dithering maps.
func xorshift(n uint) uint {
	n ^= (n << 13) & 0xFFFFFFFF
	n ^= (n >> 17) & 0xFFFFFFFF
	n ^= (n << 5) & 0xFFFFFFFF
	return n
}
