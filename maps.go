package dither

//Dither is a representation of a dithering threshold map
//
//This is used to calculate threshold values, to place a certain number of points distributed evenly (see NewOrdered) or in a random
//looking fashion (see NewRandom and NewRandomS). Dithering threshold map contains a series of values with each values bound between
//zero and the size of the map (s*s) and represented once. This means that all numbers will be present only once throughout the map.
//
//This can then be used to look at values below or above a certain number. If one needs to have 64 points in a map containing 1032
//points, one can look at the position of all the points below 64. Depending on the value placement algorithm (e.g. ordered, random),
//the values below that number could look evenly distributed across the map, or randomly picked.
//
////The dither map is internally represented as a slice of size s*s in row-major order. This means that the values correspond to the
//following (x, y) coordinates: (0, 0), (1, 0), (2, 0) ... (s-1, 0), (1, 1), (2, 1) ... (s-2, s-1), (s-1, s-1).
type Dither struct {
	Data []uint
}

//New creates a new dithering threshold map of size s*s
//
//This initializes an empty map that does not contain any value and therefore should seldomly be used directly.
func New(s uint) *Dither {
	return &Dither{
		Data: make([]uint, s*s),
	}
}

//NewOrdered creates a new ordered dithering threshold map of size s*s
//
//See https://en.wikipedia.org/wiki/Ordered_dithering
//
//This assumes that s is a power of two.
func NewOrdered(s uint) *Dither {
	d := New(s)
	bc := uint(0)
	for v := s; v != 1; v = v >> 1 {
		bc++
	}

	for y := uint(0); y < s; y++ {
		for x := uint(0); x < s; x++ {
			val := reverseInterleave(x^y, y, bc)
			d.Data[y*s+x] = val
		}
	}

	return d
}

//NewRandom returns a random dithering threshold map
//
//This uses 0x7FFFFFFF (the 9th Mersenne Prime) as seed
func NewRandom(s uint) *Dither {
	return NewRandomS(s, seed)
}

//NewRandomS returns a random dithering threshold map using a user-provided seed
//
//This uses xorshift as a pseudo-random number generator. It is amongst the fastest PRNG available, but does not provide any cryptographic
//guarantee, which are not needed for this use-case. Using the same seed will yield the same dithering threshold map.
func NewRandomS(s, ns uint) *Dither {
	d := New(s)
	for i := uint(0); i < s*s; i++ {
		d.Data[i] = i
	}
	for i := s*s - 1; i > 0; i-- {
		ns = xorshift(ns)
		d.Data[i], d.Data[ns%i] = d.Data[ns%i], d.Data[i]
	}
	return d
}

//Map maps a function over the dithering map and returns a slice
//
//The mapping function should take an unsigned integer and return an unsigned integer.
func (d *Dither) Map(f func(uint) uint) []uint {
	o := make([]uint, len(d.Data))
	for i := 0; i < len(d.Data); i++ {
		o[i] = f(d.Data[i])
	}

	return o
}

//Threshold returns a x*y slice containing 1 for all values below the threshold and 0 for all other values.
//
//This can be used to get t out of x*y points from the dithering map. The output slice will always have `t` indices with a value of one,
//as long as t is less than x*y. Otherwise, it will contain x*y indices with a value of one (the entire slice).
func (d *Dither) Threshold(t uint) []uint {
	return d.Map(func(v uint) uint {
		if v < t {
			return 1
		}
		return 0
	})
}
