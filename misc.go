package gomisc

import (
	"fmt"
	"log"
	"strconv"
)

type SInt interface {
	int | int8 | int16 | int32 | int64
}
type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}
type Int interface {
	SInt | UInt
}
type Float interface {
	float32 | float64
}
type Number interface {
	Int | Float
}

// Panic if error != nil
func PanicErr(desc string, err error) {
	if err != nil {
		log.Panic(desc, err)
	}
}

// Panic if b
func PanicIf(b bool, desc string) {
	if b {
		log.Panic(desc)
	}
}

// Fatal if error != nil
func FatalErr(desc string, err error) {
	if err != nil {
		log.Fatal(desc, err)
	}
}

// Fatal if b
func FatalIf(b bool, desc string) {
	if b {
		log.Fatal(desc)
	}
}

// Apply f to every s producing a new slice of results
func MapF[T any, U any](s []T, f func(T) U) []U {
	n := make([]U, len(s))
	for i, e := range s {
		n[i] = f(e)
	}
	return n
}

// Glues 2 values
type Pair[T, U any] struct {
	a T
	b U
}

// Merges 2 slices into a slice of pairs
func Zip[T, U any](ts []T, us []U) []Pair[T, U] {
	PanicIf(len(ts) != len(us), fmt.Sprintf("Can't zip slices of different length (%v vs %v)", len(ts), len(us)))
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}

// Folds a slice into a single value using f
func Reduce[T, U any](slice []T, acc U, f func(U, T) U) U {
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// Ternary expression
func Ternary[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}

// Bool to int
func BToI(value bool) int {
	return Ternary(value, 1, 0)
}

// Bool to float64
func BToF(value bool) float64 {
	return Ternary(value, 1., 0.)
}

// Bool to Number
func BToN[T Number](value bool) T {
	return T(Ternary(value, 1, 0))
}

// Number to bool
func NToB[T Number](value T) bool {
	return value > 0
}

// Int to ascii
func IToA[T Int](value T) string {
	return strconv.Itoa(int(value))
}

// Ascii to int
func AToI[T Int](value string) (T, error) {
	i, err := strconv.Atoi(value)
	return T(i), err
}

// Make 2dimentional slice
func Make2[T any](a, b int) [][]T {
	result := make([][]T, a)
	for i := range result {
		result[i] = make([]T, b)
	}
	return result
}

// Number of true in a bool slice
func CountTrue(bools []bool) int {
	return Reduce(bools, 0, func(s int, b bool) int { return Ternary(b, s+1, s) })
}

// Index of the first true inside bool slice
func FirstTrueIndex(bools []bool) (index int, ok bool) {
	for i, b := range bools {
		if b {
			return i, true
		}
	}
	return 0, false
}

// Slice copy with size len
func SliceNewCopy[T any](slice []T, len int) []T {
	new := make([]T, len)
	copy(new, slice)
	return new
}

// Copies slice with double the size
func SliceExpand[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)*2, min))
}

// Copies slice with half the size
func SliceShrink[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)/2, min))
}
