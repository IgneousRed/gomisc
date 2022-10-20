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
func Cos[T Float](value T) T {
	return T(math.Cos(float64(value)))
}
func Sin[T Float](value T) T {
	return T(math.Sin(float64(value)))
}
