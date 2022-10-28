package gomisc

import "unsafe"

//	func BitSlice[T UInt](value T, index, len uint8) T {
//		return
//	}
func LowestNBits[T UInt](value T, n uint8) T {
	return value & T(1<<n-1)
}
func ClearLowestNBits[T UInt](value T, n uint8) T {
	return value & ^T(1<<n-1)
}
func HighestNBits[T UInt](value T, n uint8) T {
	return value >> (uint8(unsafe.Sizeof(value))*8 - n)
}

// type UInt interface {
// 	uint | uint8 | uint16 | uint32 | uint64 | uintptr
// }
// func Foo[T UInt](value T) {
// 	bitCount := // ?
// }
