package gomisc

import "math"

// Mathematical constants.
const (
	E       = 2.7182818284590452353602874713527
	Phi     = 1.6180339887498948482045868343656
	Pi      = 3.1415926535897932384626433832795
	Tau     = 6.2831853071795864769252867665590
	Rad2Deg = 180 / math.Pi
	Deg2Rad = math.Pi / 180

	Sqrt2   = 1.4142135623730950488016887242097
	SqrtE   = 1.6487212707001281468486507878142
	SqrtPi  = 1.7724538509055160272981674833411
	SqrtPhi = 1.2720196495140689642524224617375

	Ln2    = 0.6931471805599453094172321214582
	Log2E  = 1 / Ln2
	Ln10   = 2.3025850929940456840179914546844
	Log10E = 1 / Ln10
)

const F32Sign = 1
const F32Exponent = 8
const F32Fraction = 23
const F64Sign = 1
const F64Exponent = 11
const F64Fraction = 52

type Rad f64

// Angle in degrees.
func Deg(value f64) Rad {
	return Rad(value * Deg2Rad)
}

// Radian cosine.
func (r Rad) Cos() f64 {
	return Cos(f64(r))
}

// Radian sine.
func (r Rad) Sin() f64 {
	return Sin(f64(r))
}

// Float32 bits
var F32ToU32 = math.Float32bits

// Float64 bits
var F64ToU64 = math.Float64bits

// Float32 from bits
var U32ToF32 = math.Float32frombits

// Float64 from bits
var U64ToF64 = math.Float64frombits

// 1 sign bit, 8 exponent bits and 23 fraction bits.
func F32ToParts(value f32) (bool, u8, u32) {
	bits := math.Float32bits(value)
	return NToB(bits >> (F32Exponent + F32Fraction)),
		u8(bits >> F32Fraction),
		LowestBitsU32(bits, F32Fraction)
}

// 1 sign bit, 11 exponent bits and 52 fraction bits.
func F64ToParts(value f64) (bool, u16, u64) {
	bits := math.Float64bits(value)
	return NToB(bits >> (F64Exponent + F64Fraction)),
		u16(bits >> F64Fraction),
		LowestBitsU64(bits, F64Fraction)
}

// 1 sign bit, 8 exponent bits and 23 fraction bits.
func PartsToF32(sign bool, exponent u8, fraction u32) f32 {
	return math.Float32frombits(
		BToN[u32](sign)<<(F32Exponent+F32Fraction) |
			u32(exponent)<<F32Fraction |
			LowestBitsU32(fraction, F32Fraction))
}

// 1 sign bit, 11 exponent bits and 52 fraction bits.
func PartsToF64(sign bool, exponent u16, fraction u64) f64 {
	return math.Float64frombits(
		BToN[u64](sign)<<(F64Exponent+F64Fraction) |
			u64(LowestBitsU16(exponent, F64Exponent))<<F64Fraction |
			LowestBitsU64(fraction, F64Fraction))
}

// Radian cosine.
var Cos = math.Cos

// Radian cosine.
var Sin = math.Sin

// Origin to point angle.
var Atan2 = math.Atan2

// Sign of `num`.
func Sign[T Float](num T) T {
	if num < 0 {
		return -1
	}
	return 1
}

// Discards fraction value.
// Casts to int64
func Floor(value f64) f64 {
	if value < 0 {
		return f64(s64(value) - 1)
	}
	return f64(s64(value))
}

// Discards fraction value.
func FloorI(value f64) int {
	if value < 0 {
		return int(value) - 1
	}
	return int(value)
}

// Discards fraction value.
// Casts to int64
func Round(value f64) f64 {
	if value < 0 {
		return f64(s64(value - .5))
	}
	return f64(s64(value + .5))
}

// Discards fraction value.
func RoundI(value f64) int {
	if value < 0 {
		return int(value - .5)
	}
	return int(value + .5)
}

// Square root.
func Sqrt(value f64) f64 {
	return math.Sqrt(value)
}

// Linear interpolation.
func Lerp(a, b, t f64) f64 {
	return a + (b-a)*t
}

// Cubic transition.
func FadeCubic(t f64) f64 {
	return t * t * (3 - 2*t)
}

// Quintic transition.
func FadeQuintic(t f64) f64 {
	return t * t * t * (t*(t*6-15) + 10)
}
