package gomisc

import (
	"os"
	"time"
)

func rot32(value uint32, n uint8) uint32 {
	return value<<n | value>>(-n&31)
}
func rot64(value uint64, n uint8) uint64 {
	return value<<n | value>>(-n&63)
}

// Generates a seed from multiple low quality random numbers
type SeedGen64 struct {
	result  uint64
	hashNum uint64
	rotNum  uint8
}

// Mix in another number
func (g SeedGen64) Mix(value uint64) SeedGen64 {
	value ^= g.hashNum
	g.hashNum *= 0x931e8875
	value *= g.hashNum
	value ^= rot64(value, g.rotNum)
	g.rotNum = (g.rotNum + 39) % 64
	g.result = g.result*0xca01f9dd - value*0x4973f715
	g.result ^= rot64(g.result, g.rotNum)
	return g
}

// Generates a seed from multiple low quality random numbers
func SeedGen64New(value uint64) SeedGen64 {
	return SeedGen64{0xfd258f8f3210c68, 0x43b0d7e5, 38}.Mix(value)
}

// Returns the resulting seed
func (g SeedGen64) Result() uint64 {
	return g.result
}

// Generates a seed based on time and pid
func SeedGen64Auto() uint64 {
	result := SeedGen64New(uint64(time.Now().UnixNano()))
	return result.Mix(uint64(os.Getpid())).Result()
}

type MCG32 uint64

func MCG32New(seed uint64) MCG32 {
	return MCG32(seed*2 + 1)
}

// Initializes with SeedGen64Auto
func MCG32Init() MCG32 {
	return MCG32New(SeedGen64Auto())
}

// Generates a random uint32 number
func (s *MCG32) Next() uint32 {
	state := *s
	*s = state * 0xa343836d
	return uint32(state)
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *MCG32) Range(n uint32) uint32 {
	return s.Next() % n
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *MCG32) RangeI(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1)
func (s *MCG32) Exclusive01F32() float32 {
	return float32(s.Next()) / float32(1<<32)
}

// Generates number in range [0,1). Has 32bit resolution
func (s *MCG32) Exclusive01F64() float64 {
	return float64(s.Next()) / float64(1<<32)
}

type PCG32Fast uint64

func PCG32FastNew(seed uint64) PCG32Fast {
	return PCG32Fast(seed*2 + 1)
}

// Initializes with SeedGen64Auto
func PCG32FastInit() PCG32Fast {
	return PCG32FastNew(SeedGen64Auto())
}

// Generates a random uint32 number
func (s *PCG32Fast) Next() uint32 {
	state := *s
	*s = state * 0xa343836d
	return uint32((state ^ state>>22) >> (22 + state>>61))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32Fast) Range(n uint32) uint32 {
	return s.Next() % n
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32Fast) RangeI(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1)
func (s *PCG32Fast) Exclusive01F32() float32 {
	return float32(s.Next()) / float32(1<<32)
}

// Generates number in range [0,1). Has 32bit resolution
func (s *PCG32Fast) Exclusive01F64() float64 {
	return float64(s.Next()) / float64(1<<32)
}

type PCG32 uint64

func PCG32New(seed uint64) PCG32 {
	return PCG32(seed)
}

// Initializes with SeedGen64Auto
func PCG32Init() PCG32 {
	return PCG32New(SeedGen64Auto())
}

// Generates a random uint32 number
func (s *PCG32) Next() uint32 {
	state := uint64(*s)
	*s = PCG32(state*0xa343836d + 0xfd258f8f3210c68)
	return rot32(uint32((state^state>>18)>>27), uint8(state>>59))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32) Range(n uint32) uint32 {
	return s.Next() % n
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32) RangeI(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1)
func (s *PCG32) Exclusive01F32() float32 {
	return float32(s.Next()) / float32(1<<32)
}

// Generates number in range [0,1). Has 32bit resolution
func (s *PCG32) Exclusive01F64() float64 {
	return float64(s.Next()) / float64(1<<32)
}
