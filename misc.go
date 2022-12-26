package gomisc

import (
	"fmt"
	"log"
	"strconv"
)

type s8 = int8
type s16 = int16
type s32 = int32
type s64 = int64
type u8 = uint8
type u16 = uint16
type u32 = uint32
type u64 = uint64
type f32 = float32
type f64 = float64

type SInt interface {
	int | s8 | s16 | s32 | s64
}
type UInt interface {
	uint | u8 | u16 | u32 | u64 | up
}
type Int interface {
	SInt | UInt
}
type Float interface {
	f32 | f64
}
type Number interface {
	Int | Float
}

// Panic if `err` != nil.
func PanicErr(desc string, err error) {
	if err != nil {
		log.Panic(desc, err)
	}
}

// Panic if `b`.
func PanicIf(b bool, desc string) {
	if b {
		log.Panic(desc)
	}
}

// Fatal if `err` != nil.
func FatalErr(desc string, err error) {
	if err != nil {
		log.Fatal(desc, err)
	}
}

// Fatal if `b`.
func FatalIf(b bool, desc string) {
	if b {
		log.Fatal(desc)
	}
}

// Apply `f` to every `s` producing a slice of results.
func MapF[T any, U any](s []T, f func(T) U) []U {
	n := make([]U, len(s))
	for i, e := range s {
		n[i] = f(e)
	}
	return n
}

// Glues 2 values.
type Pair[T, U any] struct {
	a T
	b U
}

// Merges `ts` and `us` into a slice of pairs.
func Zip[T, U any](ts []T, us []U) []Pair[T, U] {
	PanicIf(len(ts) != len(us), fmt.Sprintf("Can't zip slices of different length (%v vs %v)", len(ts), len(us)))
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}

// Folds `slice` into a single value using `f` starting from `acc`.
func Reduce[T, U any](slice []T, acc U, f func(U, T) U) U {
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// `t` if `b` otherwise `f`.
func Ternary[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}

// Bool to sign.
func BToS(value bool) f64 {
	if value {
		return 1
	}
	return -1
}

// Bool to int.
func BToI(value bool) int {
	if value {
		return 1
	}
	return 0
}

// Bool to f64.
func BToF(value bool) f64 {
	if value {
		return 1
	}
	return 0
}

// Bool to Number.
func BToN[T Number](value bool) T {
	if value {
		return 1
	}
	return 0
}

// Number to bool.
func NToB[T Number](value T) bool {
	return value > 0
}

// Int to ascii.
func IToA[T Int](value T) string {
	return strconv.Itoa(int(value))
}

// Ascii to int.
func AToI[T Int](value string) (T, error) {
	i, err := strconv.Atoi(value)
	return T(i), err
}

// Make 2dimentional slice.
func Make2[T any](a, b int) [][]T {
	result := make([][]T, a)
	for i := range result {
		result[i] = make([]T, b)
	}
	return result
}

// Number of true in `bools`.
func CountTrue(bools []bool) int {
	result := 0
	for _, b := range bools {
		if b {
			result++
		}
	}
	return result
}

// Index of the first true inside `bools`.
func FirstTrueIndex(bools []bool) (index int, ok bool) {
	for i, b := range bools {
		if b {
			return i, true
		}
	}
	return 0, false
}

// Copy `slice` copy with size `len`.
func SliceNewCopy[T any](slice []T, len int) []T {
	new := make([]T, len)
	copy(new, slice)
	return new
}

// Copy `slice` with double the size.
func SliceExpand[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)*2, min))
}

// Copiy `slice` with half the size.
func SliceShrink[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)/2, min))
}
