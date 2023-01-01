package gomisc

// u16 into 2 bytes
func U16ToU8s(value u16) []u8 {
	return []u8{
		u8(value),
		u8(value >> 8),
	}
}

// u32 into 4 bytes
func U32ToU8s(value u32) []u8 {
	return []u8{
		u8(value),
		u8(value >> 8),
		u8(value >> 16),
		u8(value >> 24),
	}
}

// u64 into 8 bytes
func U64ToU8s(value u64) []u8 {
	return []u8{
		u8(value),
		u8(value >> 8),
		u8(value >> 16),
		u8(value >> 24),
		u8(value >> 32),
		u8(value >> 40),
		u8(value >> 48),
		u8(value >> 56),
	}
}

// s16 into 2 bytes
func S16ToU8s(value s16) []u8 {
	return U16ToU8s(u16(value))
}

// s32 into 4 bytes
func S32ToU8s(value s32) []u8 {
	return U32ToU8s(u32(value))
}

// s64 into 8 bytes
func S64ToU8s(value s64) []u8 {
	return U64ToU8s(u64(value))
}

// f32 into 4 bytes
func F32ToU8s(value f32) []u8 {
	return U32ToU8s(F32ToU32(value))
}

// f64 into 8 bytes
func F64ToU8s(value f64) []u8 {
	return U64ToU8s(F64ToU64(value))
}

// u16 from 2 bytes
func U8sToU16(bytes []u8) u16 {
	return u16(bytes[0]) |
		u16(bytes[1])<<8
}

// u32 from 4 bytes
func U8sToU32(bytes []u8) u32 {
	return u32(bytes[0]) |
		u32(bytes[1])<<8 |
		u32(bytes[2])<<16 |
		u32(bytes[3])<<24
}

// u64 from 8 bytes
func U8sToU64(bytes []u8) u64 {
	return u64(bytes[0]) |
		u64(bytes[1])<<8 |
		u64(bytes[2])<<16 |
		u64(bytes[3])<<24 |
		u64(bytes[4])<<32 |
		u64(bytes[5])<<40 |
		u64(bytes[6])<<48 |
		u64(bytes[7])<<56
}

// s16 from 2 bytes
func U8sToS16(bytes []u8) s16 {
	return s16(U8sToU16(bytes))
}

// s32 from 4 bytes
func U8sToS32(bytes []u8) s32 {
	return s32(U8sToU32(bytes))
}

// s64 from 8 bytes
func U8sToS64(bytes []u8) s64 {
	return s64(U8sToU64(bytes))
}

// f32 from 4 bytes
func U8sToF32(bytes []u8) f32 {
	return U32ToF32(U8sToU32(bytes))
}

// f64 from 8 bytes
func U8sToF64(bytes []u8) f64 {
	return U64ToF64(U8sToU64(bytes))
}
