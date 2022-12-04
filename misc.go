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

func PanicErr(desc string, err error) {
	if err != nil {
		log.Panic(desc, err)
	}
}
func PanicIf(b bool, desc string) {
	if b {
		log.Panic(desc)
	}
}
func FatalErr(desc string, err error) {
	if err != nil {
		log.Fatal(desc, err)
	}
}
func FatalIf(b bool, desc string) {
	if b {
		log.Fatal(desc)
	}
}

func MapF[T any, U any](a []T, f func(T) U) []U {
	n := make([]U, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

type Pair[T, U any] struct {
	a T
	b U
}

func Zip[T, U any](ts []T, us []U) []Pair[T, U] {
	PanicIf(len(ts) != len(us), fmt.Sprintf("Can't zip slices of different length (%v vs %v)", len(ts), len(us)))
	pairs := make([]Pair[T, U], len(ts))
	for i := 0; i < len(ts); i++ {
		pairs[i] = Pair[T, U]{ts[i], us[i]}
	}
	return pairs
}
func Reduce[T, U any](slice []T, acc U, f func(U, T) U) U {
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}
func Ternary[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}
func BToI(value bool) int {
	return Ternary(value, 1, 0)
}
func BToF(value bool) float64 {
	return Ternary(value, 1., 0.)
}
func BToN[T Number](value bool) T {
	return T(Ternary(value, 1, 0))
}
func NToB[T Number](value T) bool {
	return value > 0
}
func IToA[T Int](value T) string {
	return strconv.Itoa(int(value))
}
func AToI[T Int](value string) (T, error) {
	i, err := strconv.Atoi(value)
	return T(i), err
}
func Make2[T any](a, b int) [][]T {
	result := make([][]T, a)
	for i := range result {
		result[i] = make([]T, b)
	}
	return result
}
func CountTrue(bools []bool) int {
	return Reduce(bools, 0, func(s int, b bool) int { return Ternary(b, s+1, s) })
}
func FirstTrueIndex(bools []bool) int {
	for i, b := range bools {
		if b {
			return i
		}
	}
	return -1
}
func SliceNewCopy[T any](slice []T, newLen int) []T {
	new := make([]T, newLen)
	copy(new, slice)
	return new
}
func SliceExpand[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)*2, min))
}
func SliceShrink[T any](slice []T, min int) []T {
	return SliceNewCopy(slice, Max(len(slice)/2, min))
}
