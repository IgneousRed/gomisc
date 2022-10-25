package gomisc

type Vec[T Number] []T

func (v Vec[T]) Copy() Vec[T] {
	result := make(Vec[T], len(v))
	copy(result, v)
	return result
}
func (v Vec[T]) Float32() Vec[float32] {
	return MapF(v, func(v T) float32 { return float32(v) })
}
func VecsToFloat32s[T Number](v []Vec[T]) []Vec[float32] {
	return MapF(v, func(v Vec[T]) Vec[float32] { return v.Float32() })
}
func (v Vec[T]) Int() Vec[int] {
	return MapF(v, func(v T) int { return int(v) })
}
func VecsToInts[T Number](v []Vec[T]) []Vec[int] {
	return MapF(v, func(v Vec[T]) Vec[int] { return v.Int() })
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
