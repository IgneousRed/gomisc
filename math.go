package gomisc

import "math"

// Wraps `value` to `len`, even if `value` < 0.
func Wrap[T Number](value, len T) T {
	return value - T(Floor(float64(value)/float64(len)))*len
}

// Absolute value (non-negative).
func Abs[T Number](value T) T { // TODO VS math version
	if value < 0 {
		return -value
	}
	return value
}

// Confines `value` between `min`-`max` range.
func Clamp[T Number](value, min, max T) T {
	if value >= max {
		return max
	} else if value <= min {
		return min
	}
	return value
}

// True if `value` is outside `min`-`max` range.
func IsOutside[T Number](value, min, max T) (T, bool) {
	if value > max {
		return max, true
	} else if value < min {
		return min, true
	}
	return value, false
}

// Power `base` to `exponent`.
func Pow[T Number](base, exp T) T {
	return T(math.Pow(float64(base), float64(exp)))
}

// `magFrom` with `signFrom` sign.
func WithSign[T Number](signFrom, magFrom T) T {
	if signFrom*magFrom < 0 {
		return -magFrom
	}
	return magFrom
}

// Sign bit and magnitude.
func SignBitAndMag[T Number](value T) (signBit si, magnitude T) {
	if value < 0 {
		return 1, -value
	}
	return 0, value
}

// The lowest value.
func Min[T Number](values ...T) T {
	result := values[0]
	for _, v := range values[1:] {
		if v < result {
			result = v
		}
	}
	return result
}

// The highest value.
func Max[T Number](values ...T) T {
	result := values[0]
	for _, v := range values[1:] {
		if v > result {
			result = v
		}
	}
	return result
}

// Sum all the values.
func Sum[T Number](values ...T) T {
	result := values[0]
	for _, v := range values[1:] {
		result += v
	}
	return result
}
