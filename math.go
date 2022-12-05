package gomisc

import "math"

// Absolute value (non-negative)
func Abs[T Number](value T) T {
	return Ternary(value < 0, -value, value)
}

// Wraps both negative and len or more
func Wrap[T Number](value, len T) T {
	return value - T(math.Floor(float64(value)/float64(len)))*len
}

// Confines value between min-max range
func Clamp[T Number](value, min, max T) T {
	return Ternary(value >= max, max, Ternary(value <= min, min, value))
}

// True if value is outside min-max range
func IsOutside[T Number](value, min, max T) (T, bool) {
	if value > max {
		return max, true
	} else if value < min {
		return min, true
	}
	return value, false
}

// The lowest value
func Min[T Number](values ...T) T {
	return Reduce(values[1:], values[0], func(a, b T) T {
		return Ternary(a > b, b, a)
	})
}

// The highest value
func Max[T Number](values ...T) T {
	return Reduce(values[1:], values[0], func(a, b T) T {
		return Ternary(a < b, b, a)
	})
}

// Base to exponent power
func Pow[T Number](base, exp T) T {
	return T(math.Pow(float64(base), float64(exp)))
}

// Sign from one with magnitude from another
func WithSign[T Number](signFrom, magFrom T) T {
	return Ternary(signFrom < 0, -Abs(magFrom), Abs(magFrom))
}

// Sign bit and magnitude
func SignBitAndMag[T Number](value T) (signBit int, magnitude T) {
	if value < 0 {
		return 1, -value
	}
	return 0, value
}

// Sum all the values
func Sum[T Number](values ...T) T {
	return Reduce(values[1:], values[0], func(a, b T) T { return a + b })
}
