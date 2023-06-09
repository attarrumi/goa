package goa

import "unsafe"

//go:linkname unsafe_one unsafe_one
func unsafe_one(fn unsafe.Pointer, x unsafe.Pointer) unsafe.Pointer

//go:linkname unsafe_two unsafe_two
func unsafe_two(fn unsafe.Pointer, x, y unsafe.Pointer) unsafe.Pointer

//go:linkname unsafe_three unsafe_three
func unsafe_three(fn unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer) unsafe.Pointer

func UnsafeOne(fn unsafe.Pointer, x unsafe.Pointer) unsafe.Pointer {
	return unsafe_one(fn, x)
}

func UnsafeTwo(fn unsafe.Pointer, x, y unsafe.Pointer) unsafe.Pointer {
	return unsafe_two(fn, x, y)
}

func UnsafeThree(fn unsafe.Pointer, x, y, z unsafe.Pointer) unsafe.Pointer {
	return unsafe_three(fn, x, y, z)
}
