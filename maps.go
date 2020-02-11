package dither

//Dither represents a dither map
type Dither struct {
	Data []uint
}

//New creates a new dither map of size s*s
func New(s uint) *Dither {
	return &Dither{
		Data: make([]uint, s*s),
	}
}

//NewOrdered creates a new ordered dither map of size s*s
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

//NewRandom returns a dither map with random placement
func NewRandom(s uint) *Dither {
	return NewRandomS(s, seed)
}

//NewRandomS returns a dither map with random placement with a specific seed
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
