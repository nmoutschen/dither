package dither

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkNewOrdered4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOrdered(4)
	}
}

func BenchmarkNewOrdered32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOrdered(32)
	}
}

func BenchmarkNewOrdered256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewOrdered(256)
	}
}

func BenchmarkNewRandomS4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomS(4, seed)
	}
}

func BenchmarkNewRandomS32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomS(32, seed)
	}
}

func BenchmarkNewRandomS256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomS(256, seed)
	}
}

func BenchmarkThreshold4(b *testing.B) {
	d := NewOrdered(4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Threshold(8)
	}
}

func BenchmarkThreshold32(b *testing.B) {
	d := NewOrdered(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Threshold(999)
	}
}

func BenchmarkThreshold256(b *testing.B) {
	d := NewOrdered(256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Threshold(30000)
	}
}

func ExampleNew() {
	d := New(2)
	fmt.Println(d.Data)
	//Output:
	//[0 0 0 0]
}

func ExampleNewOrdered() {
	d := NewOrdered(2)
	fmt.Println(d.Data)
	fmt.Println(d.Threshold(2))
	//Output:
	//[0 2 3 1]
	//[1 0 0 1]
}

func ExampleNewRandom() {
	d := NewRandom(2)
	fmt.Println(d.Data)
	fmt.Println(d.Threshold(2))
	//Output:
	//[2 3 1 0]
	//[0 0 1 1]
}

func TestNewOrdered(t *testing.T) {
	testCases := []struct {
		S        uint
		Expected []uint
	}{
		{2, []uint{0, 2, 3, 1}},
		{4, []uint{0, 8, 2, 10, 12, 4, 14, 6, 3, 11, 1, 9, 15, 7, 13, 5}},
	}

	for i, tc := range testCases {
		if res := NewOrdered(tc.S); !reflect.DeepEqual(res.Data, tc.Expected) {
			t.Errorf("NewOrdered(%d) = %v; want %v for test case %d", tc.S, res.Data, tc.Expected, i)
		}
	}
}

func TestNewRandom(t *testing.T) {
	testCases := []struct {
		S        uint
		Expected []uint
	}{
		{2, []uint{2, 3, 1, 0}},
		{4, []uint{14, 7, 8, 15, 3, 13, 12, 4, 9, 1, 6, 10, 2, 11, 5, 0}},
		{8, []uint{2, 42, 25, 62, 58, 8, 55, 53, 23, 36, 16, 46, 5, 34, 54, 50, 1, 57, 31, 29, 63, 18, 27, 15, 4, 26, 7, 28, 3, 22, 52, 47, 48, 24, 41, 19, 0, 61, 32, 44, 13, 10, 51, 37, 60, 43, 59, 45, 6, 14, 39, 21, 11, 35, 20, 56, 9, 40, 12, 33, 38, 49, 17, 30}},
	}

	for i, tc := range testCases {
		if res := NewRandom(tc.S); !reflect.DeepEqual(res.Data, tc.Expected) {
			t.Errorf("NewRandom(%d) = %v; want %v for test case %d", tc.S, res.Data, tc.Expected, i)
		}
	}
}

func TestNewRandomS(t *testing.T) {
	testCases := []struct {
		S, NS    uint
		Expected []uint
	}{
		{2, 10, []uint{1, 2, 3, 0}},
		{4, 1, []uint{5, 13, 0, 15, 7, 12, 10, 14, 2, 1, 4, 6, 11, 8, 3, 9}},
	}

	for i, tc := range testCases {
		if res := NewRandomS(tc.S, tc.NS); !reflect.DeepEqual(res.Data, tc.Expected) {
			t.Errorf("NewRandom(%d, %d) = %v; want %v for test case %d", tc.S, tc.NS, res.Data, tc.Expected, i)
		}
	}
}

func TestMap(t *testing.T) {
	lessThanTwo := func(v uint) uint {
		if v < 2 {
			return 1
		}
		return 0
	}
	lessThanFour := func(v uint) uint {
		if v < 4 {
			return 1
		}
		return 0
	}
	lessThanEight := func(v uint) uint {
		if v < 8 {
			return 1
		}
		return 0
	}

	testCases := []struct {
		D        *Dither
		F        func(uint) uint
		Expected []uint
	}{
		{NewOrdered(2), lessThanTwo, []uint{1, 0, 0, 1}},
		{NewOrdered(2), lessThanFour, []uint{1, 1, 1, 1}},
		{NewOrdered(4), lessThanTwo, []uint{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}},
		{NewOrdered(4), lessThanFour, []uint{1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0}},
		{NewOrdered(4), lessThanEight, []uint{1, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 1}},
	}

	for i, tc := range testCases {
		if res := tc.D.Map(tc.F); !reflect.DeepEqual(res, tc.Expected) {
			t.Errorf("D.Map() = %v; want %v for test case %d", res, tc.Expected, i)
		}
	}
}

func TestThreshold(t *testing.T) {
	testCases := []struct {
		S, T uint
	}{
		{4, 1},
		{4, 9},
		{8, 63},
		{256, 30000},
		{16, 7},
	}

	for i, tc := range testCases {
		d := NewOrdered(tc.S)
		count1 := uint(0)
		count0 := uint(0)
		for _, v := range d.Threshold(tc.T) {
			if v == 1 {
				count1++
			} else if v == 0 {
				count0++
			}
		}

		if count1 != tc.T {
			t.Errorf("d.Threshold(%d) contains %d 1s; want %d in test case %d", tc.T, count1, tc.T, i)
			for x := uint(0); x < tc.S; x++ {
				fmt.Println(d.Data[x*tc.S : (x+1)*tc.S])
			}
		}
		if count0 != tc.S*tc.S-tc.T {
			t.Errorf("d.Threshold(%d) contains %d 0s; want %d in test case %d", tc.T, count0, tc.S*tc.S-tc.T, i)
		}
	}
}
