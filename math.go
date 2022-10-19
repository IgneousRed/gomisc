package gomisc

import "math"

func Float32Parts(value float32) (bool, uint8, uint32) {
	asd := math.Float32bits(value)
	return IToB(asd >> 31), uint8(asd >> 23), asd & (math.MaxUint32 >> 9)
}
func Float64Parts(value float64) (bool, uint16, uint64) {
	asd := math.Float64bits(value)
	return IToB(asd >> 63), uint16((asd << 1) >> 52), asd & (math.MaxUint32 >> 12)
}
func RoundI[T Float](value T) int {
	if value < 0 {
		return int(value - .5)
	}
	return int(value + .5)
}
func Abs[T Number](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
func Clamp[T Number](value, min, max T) T {
	if value > max {
		return max
	} else if value < min {
		return min
	}
	return value
}
func ClampReport[T Number](value, min, max T) (T, bool) {
	if value > max {
		return max, true
	} else if value < min {
		return min, true
	}
	return value, false
}
func Min[T Number](values ...T) T {
	return Reduce(values, func(a, b T) T {
		if a > b {
			return b
		}
		return a
	})
}
func Max[T Number](values ...T) T {
	return Reduce(values, func(a, b T) T {
		if a < b {
			return b
		}
		return a
	})
}
func Sqrt[T Float](value T) T {
	return T(math.Sqrt(float64(value)))
}
func Pow[T Float](base, exp T) T {
	return T(math.Pow(float64(base), float64(exp)))
}

func WithSign[T Number](signFrom, magFrom T) T {
	if signFrom < 0 {
		return -Abs(magFrom)
	}
	return Abs(magFrom)
}

func SignBitAndMag[T Number](value T) (signBit int, magnitude T) {
	if value < 0 {
		return 1, -value
	}
	return 0, value
}

type Vec[T Number] []T

func (v Vec[T]) Copy() Vec[T] {
	result := make(Vec[T], len(v))
	copy(result, v)
	return result
}
func (v Vec[T]) Add1(other T) Vec[T] {
	return MapF(v, func(v T) T { return v + other })
}
func (v Vec[T]) Add(other Vec[T]) Vec[T] {
	return MapF(Zip(v, other), func(v Pair[T, T]) T { return v.a + v.b })
}
func (v Vec[T]) Sub1(other T) Vec[T] {
	return MapF(v, func(v T) T { return v - other })
}
func (v Vec[T]) Sub(other Vec[T]) Vec[T] {
	return MapF(Zip(v, other), func(v Pair[T, T]) T { return v.a - v.b })
}
func (v Vec[T]) Mul1(other T) Vec[T] {
	return MapF(v, func(v T) T { return v * other })
}
func (v Vec[T]) Mul(other Vec[T]) Vec[T] {
	return MapF(Zip(v, other), func(v Pair[T, T]) T { return v.a * v.b })
}
func (v Vec[T]) Div1(other T) Vec[T] {
	return MapF(v, func(v T) T { return v / other })
}
func (v Vec[T]) Div(other Vec[T]) Vec[T] {
	return MapF(Zip(v, other), func(v Pair[T, T]) T { return v.a / v.b })
}
func (v Vec[T]) Abs() Vec[T] {
	return MapF(v, func(v T) T { return Abs(v) })
}
func (v Vec[T]) Min() T {
	return Min(v...)
}
func (v Vec[T]) Max() T {
	return Max(v...)
}
func (v Vec[T]) RoundI() Vec[int] {
	switch v := any(v).(type) {
	case Vec[float32]:
		return MapF(v, func(v float32) int { return RoundI(v) })
	case Vec[float64]:
		return MapF(v, func(v float64) int { return RoundI(v) })
	default:
		panic("Can only round a float")
	}
}
func (v Vec[T]) Magnitude() float32 { // expand type
	PanicIf(len(v) == 0, "Trying to find magnitude of empty vec")
	result := float32(v[0] * v[0])
	for _, e := range v[1:] {
		result += float32(e * e)
	}
	return Sqrt(result)
}
func (v Vec[T]) Normalize() Vec[float32] { // expand type
	magRec := 1. / v.Magnitude()
	result := make(Vec[float32], len(v)) // with 0 len?
	for i, e := range v {
		result[i] = float32(e) * magRec
	}
	return result
}
func (v Vec[T]) Rotate90() Vec[T] {
	PanicIf(len(v) != 2, "Rotate90 requires exactly len 2")
	return Vec[T]{-v[1], v[0]}
}
