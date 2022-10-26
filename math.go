package gomisc

import "math"

func Abs[T Number](value T) T {
	if value < 0 {
		return -value
	}
	return value
}
func Repeat0[T Int](value, max T) T {
	return value - T(math.Floor(float64(value)/float64(max)))*max
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
