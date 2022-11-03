package gomisc

import "unsafe"

func LowestNBits[T UInt](value T, n uint8) T {
	return value & T(1<<n-1)
}
func ClearLowestNBits[T UInt](value T, n uint8) T {
	return value & ^T(1<<n-1)
}
func HighestNBits[T UInt](value T, n uint8) T {
	return value >> (uint8(unsafe.Sizeof(value))*8 - n)
}
func Rotate32(value uint32, n uint8) uint32 {
	return value<<n | value>>(32-n)
}
func Rotate64(value uint64, n uint8) uint64 {
	return value<<n | value>>(64-n)
}
