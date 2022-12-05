package gomisc

import "math"

type Rad32 float32
type Rad64 float64
type Rad interface {
	Rad32 | Rad64
}

const Float32Sign = 1
const Float32Exponent = 8
const Float32Fraction = 23
const Float64Sign = 1
const Float64Exponent = 11
const Float64Fraction = 52

// 1 sign bit, 8 exponent bits and 23 fraction bits
func Float32ToParts(value float32) (bool, uint8, uint32) {
	bits := math.Float32bits(value)
	return NToB(bits >> (Float32Exponent + Float32Fraction)),
		uint8(bits >> Float32Fraction),
		LowestNBits(bits, Float32Fraction)
}

// 1 sign bit, 11 exponent bits and 52 fraction bits
func Float64ToParts(value float64) (bool, uint16, uint64) {
	bits := math.Float64bits(value)
	return NToB(bits >> (Float64Exponent + Float64Fraction)),
		uint16(bits >> Float64Fraction),
		LowestNBits(bits, Float64Fraction)
}

// 1 sign bit, 8 exponent bits and 23 fraction bits
func Float32FromParts(sign bool, exponent uint8, fraction uint32) float32 {
	return math.Float32frombits(
		BToN[uint32](sign)<<(Float32Exponent+Float32Fraction) |
			uint32(exponent)<<Float32Fraction |
			LowestNBits(fraction, Float32Fraction))
}

// 1 sign bit, 11 exponent bits and 52 fraction bits
func Float64FromParts(sign bool, exponent uint16, fraction uint64) float64 {
	return math.Float64frombits(
		BToN[uint64](sign)<<(Float64Exponent+Float64Fraction) |
			uint64(LowestNBits(exponent, Float64Exponent))<<Float64Fraction |
			LowestNBits(fraction, Float64Fraction))
}

// radian cosine
func Cos[T Rad](value T) float64 {
	return math.Cos(float64(value))
}

// radian sine
func Sin[T Rad](value T) float64 {
	return math.Sin(float64(value))
}

// origin to point angle
func Atan2[T Float](y, x T) Rad64 {
	return Rad64(math.Atan2(float64(y), float64(x)))
}

// discards fraction value
func Floor[T Float](value T) T {
	return T(Ternary(value < 0, int64(value)-1, int64(value)))
}

// discards fraction value
func FloorI[T Float](value T) int {
	return int(Ternary(value < 0, int64(value)-1, int64(value)))
}

// discards fraction value
func Round[T Float](value T) T {
	return value + T(Ternary(value < 0, -.5, .5))
}

// discards fraction value
func RoundI[T Float](value T) int {
	return int(value + T(Ternary(value < 0, -.5, .5)))
}

// square root
func Sqrt[T Float](value T) T {
	return T(math.Sqrt(float64(value)))
}

// Linear interpolation
func Lerp[T Float](a, b T, t T) T {
	return T(a + (b-a)*t)
}

// Quintic transition
func Fade[T Float](t T) T {
	return t * t * t * (t*(t*6-15) + 10)
}
