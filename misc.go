package gomisc

import (
	"fmt"
	"log"
	"strconv"
)

type Atom struct{}
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
func Reduce[T any](s []T, f func(T, T) T) T {
	if len(s) == 0 {
		var result T
		return result
	}
	result := s[0]
	for _, v := range s[1:] {
		result = f(result, v)
	}
	return result
}
func Ternary[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}
func BToI(value bool) int {
	if value {
		return 1
	}
	return 0
}
func IToB[T Number](value T) bool {
	return value > 0
}
func IToA[T Int](value T) string {
	return strconv.Itoa(int(value))
}
func AToI[T Int](value string) (T, error) {
	i, err := strconv.Atoi(value)
	return T(i), err
}
