package gomisc

import "math"

type Rad64 float64
type Deg64 float64

const F32Sign = 1
const F32Exponent = 8
const F32Fraction = 23
const F64Sign = 1
const F64Exponent = 11
const F64Fraction = 52

// 1 sign bit, 8 exponent bits and 23 fraction bits
func F32ToParts(value float32) (bool, uint8, uint32) {
	bits := math.Float32bits(value)
	return NToB(bits >> (F32Exponent + F32Fraction)),
		uint8(bits >> F32Fraction),
		LowestBitsU32(bits, F32Fraction)
}

// 1 sign bit, 11 exponent bits and 52 fraction bits
func F64ToParts(value float64) (bool, uint16, uint64) {
	bits := math.Float64bits(value)
	return NToB(bits >> (F64Exponent + F64Fraction)),
		uint16(bits >> F64Fraction),
		LowestBitsU64(bits, F64Fraction)
}

// 1 sign bit, 8 exponent bits and 23 fraction bits
func PartsToF32(sign bool, exponent uint8, fraction uint32) float32 {
	return math.Float32frombits(
		BToN[uint32](sign)<<(F32Exponent+F32Fraction) |
			uint32(exponent)<<F32Fraction |
			LowestBitsU32(fraction, F32Fraction))
}

// 1 sign bit, 11 exponent bits and 52 fraction bits
func PartsToF64(sign bool, exponent uint16, fraction uint64) float64 {
	return math.Float64frombits(
		BToN[uint64](sign)<<(F64Exponent+F64Fraction) |
			uint64(LowestBitsU16(exponent, F64Exponent))<<F64Fraction |
			LowestBitsU64(fraction, F64Fraction))
}

// Convert to degrees
func (r Rad64) Deg() Deg64 {
	return Deg64(r * 180 / math.Pi)
}

// Convert to radians
func (r Deg64) Rad() Rad64 {
	return Rad64(r * math.Pi / 180)
}

// Radian cosine
func (r Rad64) Cos() float64 {
	return math.Cos(float64(r))
}

// Degree cosine
func (r Deg64) Cos() float64 {
	return r.Rad().Cos()
}

// Radian sine
func (r Rad64) Sin() float64 {
	return math.Sin(float64(r))
}

// Degree sine
func (r Deg64) Sin() float64 {
	return r.Rad().Sin()
}

// Origin to point angle
func Atan2(y, x float64) Rad64 {
	return Rad64(math.Atan2(y, x))
}

// Discards fraction value
func Floor(value float64) float64 { // TODO VS math version
	if value < 0 {
		return float64(int64(value) - 1)
	}
	return float64(int64(value))
}

// Discards fraction value
func FloorI(value float64) int { // TODO VS math version
	if value < 0 {
		return int(value) - 1
	}
	return int(value)
}

// Discards fraction value
func Round(value float64) float64 { // TODO VS math version
	if value < 0 {
		return float64(int64(value - .5))
	}
	return float64(int64(value + .5))
}

// Discards fraction value
func RoundI(value float64) int { // TODO VS math version
	if value < 0 {
		return int(value - .5)
	}
	return int(value + .5)
}

// Square root
func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

// Linear interpolation
func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

// Cubic transition
func FadeCubic(t float64) float64 {
	return t * t * (3 - 2*t)
}

// Quintic transition
func FadeQuintic(t float64) float64 {
	return t * t * t * (t*(t*6-15) + 10)
}
