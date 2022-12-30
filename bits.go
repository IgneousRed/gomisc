package gomisc

// lowest `n` bits from `value`.
func LowestBitsU16(value u16, n u8) u16 {
	return value & u16(1<<n-1)
}

// lowest `n` bits from `value`.
func LowestBitsU32(value u32, n u8) u32 {
	return value & u32(1<<n-1)
}

// lowest `n` bits from `value`.
func LowestBitsU64(value u64, n u8) u64 {
	return value & u64(1<<n-1)
}

// Rotate bits in `value` by `n` places.
func RotateU32(value u32, n u8) u32 {
	return value<<n | value>>(32-n)
}

// Rotate bits in `value` by `n` places.
func RotateU64(value u64, n u8) u64 {
	return value<<n | value>>(64-n)
}
