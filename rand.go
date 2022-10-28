package gomisc

import (
	"os"
	"time"
)

func rot32(value uint32, n uint8) uint32 {
	return value<<n | value>>(32-n)
}
func rot64(value uint64, n uint8) uint64 {
	return value<<n | value>>(64-n)
}

// Generates a seed from multiple low quality random numbers
type SeedGen64 struct {
	result  uint64
	hashNum uint64
}

// Mix in another number
func (g SeedGen64) Mix(value uint64) SeedGen64 {
	value ^= g.hashNum
	g.hashNum *= 0x931e8875
	value *= g.hashNum
	value ^= rot64(value, 32)
	g.result = g.result*0xca01f9dd - value*0x4973f715
	g.result ^= rot64(g.result, 32)
	return g
}

// Generates a seed from multiple low quality random numbers
func SeedGen64New(value uint64) SeedGen64 {
	return SeedGen64{0x9e3779b97f4a7c15, 0x43b0d7e5}.Mix(value)
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

// Initializes with seed
func MCG32New(seed uint64) MCG32 {
	return MCG32(seed*2 + 1)
}

// Initializes with SeedGen64Auto
func MCG32Init() MCG32 {
	return MCG32New(SeedGen64Auto())
}

func (s *MCG32) raw() uint64 {
	state := *s
	*s = state * 0xf13283ad
	return uint64(state)
}

// Generates a random uint32 number
func (s *MCG32) Next() uint32 {
	return uint32(s.raw() >> 32)
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *MCG32) Range(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1]
func (s *MCG32) Normal32() float32 {
	return float32(s.raw()) / float32(1<<32-1)
}

// Generates number in range [0,1]
func (s *MCG32) Normal64() float64 {
	return float64(s.raw()) / float64(1<<64-1)
}

type PCG32Fast uint64

// Initializes with seed
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
	*s = state * 0xf13283ad
	return uint32((state ^ state>>22) >> (22 + state>>61))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32Fast) Range(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1]
func (s *PCG32Fast) Normal32() float32 {
	return float32(s.Next()) / float32(1<<32-1)
}

// Generates number in range [0,1]. Has 32bit resolution
func (s *PCG32Fast) Normal64() float64 {
	return float64(s.Next()) / float64(1<<64-1)
}

type PCG32 uint64

// Initializes with seed
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
	*s = PCG32(state*0xf13283ad + 0x9e3779b97f4a7c15)
	return rot32(uint32((state^state>>18)>>27), uint8(state>>59))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32) Range(n int) int {
	return int(s.Next() % uint32(n))
}

// Generates number in range [0,1]
func (s *PCG32) Normal32() float32 {
	return float32(s.Next()) / float32(1<<32-1)
}

// Generates number in range [0,1]. Has 32bit resolution
func (s *PCG32) Normal64() float64 {
	return float64(s.Next()) / float64(1<<64-1)
}
