package gomisc

// lowest `n` bits from `value`
func LowestBitsU16(value uint16, n uint8) uint16 {
	return value & uint16(1<<n-1)
}

// lowest `n` bits from `value`
func LowestBitsU32(value uint32, n uint8) uint32 {
	return value & uint32(1<<n-1)
}

// lowest `n` bits from `value`
func LowestBitsU64(value uint64, n uint8) uint64 {
	return value & uint64(1<<n-1)
}

//	func ClearLowestNBits[T UInt](value T, n uint8) T {
//		return value & ^T(1<<n-1)
//	}
//
//	func HighestNBits[T UInt](value T, n uint8) T {
//		return value >> (uint8(unsafe.Sizeof(value))*8 - n)
//	}

// Rotate bits in `value` by `n` places
func RotateU32(value uint32, n uint8) uint32 {
	return value<<n | value>>(32-n)
}

// Rotate bits in `value` by `n` places
func RotateU64(value uint64, n uint8) uint64 {
	return value<<n | value>>(64-n)
}
