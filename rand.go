package gomisc

import (
	"os"
	"time"
)

// TODO play with rand

// golden ratio: sqrt(5) / 2 - .5

// Scrambles bits from multiple low quality random bits
type Hash32 struct {
	result  uint32
	hashNum uint32
}

// Mix in another number
func (h Hash32) Mix(value uint32) Hash32 {
	value ^= h.hashNum
	h.hashNum *= 0x931e8875
	value *= h.hashNum
	value ^= value >> 16
	h.result = h.result*0xca01f9dd - value*0x4973f715
	h.result ^= h.result >> 16
	return h
}

// Scrambles bits from multiple low quality random bits
func Hash32New(value uint32) Hash32 {
	return Hash32{0x9e3779b9, 0x43b0d7e5}.Mix(value)
}

// Returns the resulting bits
func (g Hash32) Result() uint32 {
	return g.result
}

// Scrambles bits from multiple low quality random bits
func Hash32From(values ...uint32) uint32 {
	result := Hash32New(values[0])
	for _, value := range values[1:] {
		result = result.Mix(value)
	}
	return result.Result()
}

// Scrambles bits from multiple low quality random bits
type Hash64 struct {
	result  uint64
	hashNum uint64
}

// Mix in another number
func (g Hash64) Mix(value uint64) Hash64 {
	value ^= g.hashNum
	g.hashNum *= 0x931e8875
	value *= g.hashNum
	value ^= value >> 32
	g.result = g.result*0xca01f9dd - value*0x4973f715
	g.result ^= g.result >> 32
	return g
}

// Scrambles bits from multiple low quality random bits
func Hash64New(value uint64) Hash64 {
	return Hash64{0x9e3779b97f4a7c15, 0x43b0d7e5}.Mix(value)
}

// Returns the resulting bits
func (g Hash64) Result() uint64 {
	return g.result
}

// Scrambles bits from multiple low quality random bits
func Hash64From(values ...uint64) uint64 {
	result := Hash64New(values[0])
	for _, value := range values[1:] {
		result = result.Mix(value)
	}
	return result.Result()
}

// Generates a seed based on time and pid
func SeedGen64() uint64 {
	return Hash64From(uint64(time.Now().UnixNano()), uint64(os.Getpid()))
}

type MCG32 uint64

// Initializes with seed
func MCG32New(seed uint64) MCG32 {
	return MCG32(seed*2 + 1)
}

// Initializes with SeedGen64
func MCG32Init() MCG32 {
	return MCG32New(SeedGen64())
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
	return float64(s.raw()) / float64(1<<32-1)
}

type PCG32Fast uint64

// Initializes with seed
func PCG32FastNew(seed uint64) PCG32Fast {
	return PCG32Fast(seed*2 + 1)
}

// Initializes with SeedGen64
func PCG32FastInit() PCG32Fast {
	return PCG32FastNew(SeedGen64())
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
	return float64(s.Next()) / float64(1<<32-1)
}

type PCG32 uint64

// Initializes with seed
func PCG32New(seed uint64) PCG32 {
	return PCG32(seed)
}

// Initializes with SeedGen64
func PCG32Init() PCG32 {
	return PCG32New(SeedGen64())
}

// Generates a random uint32 number
func (s *PCG32) Next() uint32 {
	state := uint64(*s)
	*s = PCG32(state*0xf13283ad + 0x9e3779b97f4a7c15)
	return RotateU32(uint32((state^state>>18)>>27), uint8(state>>59))
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
	return float64(s.Next()) / float64(1<<32-1)
}
