package gomisc

import (
	"os"
	"time"
)

// TODO play with rand

// golden ratio: sqrt(5) / 2 - .5

// Scrambles bits from multiple low quality random bits
type Hash32 struct {
	result  u32
	hashNum u32
}

// Mix in another number
func (h Hash32) Mix(value u32) Hash32 {
	value ^= h.hashNum
	h.hashNum *= 0x931e8875
	value *= h.hashNum
	value ^= value >> 16
	h.result = h.result*0xca01f9dd - value*0x4973f715
	h.result ^= h.result >> 16
	return h
}

// Scrambles bits from multiple low quality random bits
func Hash32New(value u32) Hash32 {
	return Hash32{0x9e3779b9, 0x43b0d7e5}.Mix(value)
}

// Returns the resulting bits
func (g Hash32) Result() u32 {
	return g.result
}

// Scrambles bits from multiple low quality random bits
func Hash32From(values ...u32) u32 {
	result := Hash32New(values[0])
	for _, value := range values[1:] {
		result = result.Mix(value)
	}
	return result.Result()
}

// Scrambles bits from multiple low quality random bits
type Hash64 struct {
	result  u64
	hashNum u64
}

// Mix in another number
func (g Hash64) Mix(value u64) Hash64 {
	value ^= g.hashNum
	g.hashNum *= 0x931e8875
	value *= g.hashNum
	value ^= value >> 32
	g.result = g.result*0xca01f9dd - value*0x4973f715
	g.result ^= g.result >> 32
	return g
}

// Scrambles bits from multiple low quality random bits
func Hash64New(value u64) Hash64 {
	return Hash64{0x9e3779b97f4a7c15, 0x43b0d7e5}.Mix(value)
}

// Returns the resulting bits
func (g Hash64) Result() u64 {
	return g.result
}

// Scrambles bits from multiple low quality random bits
func Hash64From(values ...u64) u64 {
	result := Hash64New(values[0])
	for _, value := range values[1:] {
		result = result.Mix(value)
	}
	return result.Result()
}

// Generates a seed based on time and pid
func SeedGen64() u64 {
	return Hash64From(u64(time.Now().UnixNano()), u64(os.Getpid()))
}

type MCG32 u64

// Initializes with seed
func MCG32New(seed u64) MCG32 {
	return MCG32(seed*2 + 1)
}

// Initializes with SeedGen64
func MCG32Init() MCG32 {
	return MCG32New(SeedGen64())
}

func (s *MCG32) raw() u64 {
	state := *s
	*s = state * 0xf13283ad
	return u64(state)
}

// Generates a random uint32 number
func (s *MCG32) Next() u32 {
	return u32(s.raw() >> 32)
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *MCG32) Range(n si) si {
	return si(s.Next() % u32(n))
}

// Generates number in range [0,1]
func (s *MCG32) Normal32() f32 {
	return f32(s.raw()) / f32(1<<32-1)
}

// Generates number in range [0,1]
func (s *MCG32) Normal64() f64 {
	return f64(s.raw()) / f64(1<<32-1)
}

type PCG32Fast u64

// Initializes with seed
func PCG32FastNew(seed u64) PCG32Fast {
	return PCG32Fast(seed*2 + 1)
}

// Initializes with SeedGen64
func PCG32FastInit() PCG32Fast {
	return PCG32FastNew(SeedGen64())
}

// Generates a random uint32 number
func (s *PCG32Fast) Next() u32 {
	state := *s
	*s = state * 0xf13283ad
	return u32((state ^ state>>22) >> (22 + state>>61))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32Fast) Range(n si) si {
	return si(s.Next() % u32(n))
}

// Generates number in range [0,1]
func (s *PCG32Fast) Normal32() f32 {
	return f32(s.Next()) / f32(1<<32-1)
}

// Generates number in range [0,1]. Has 32bit resolution
func (s *PCG32Fast) Normal64() f64 {
	return f64(s.Next()) / f64(1<<32-1)
}

type PCG32 u64

// Initializes with seed
func PCG32New(seed u64) PCG32 {
	return PCG32(seed)
}

// Initializes with SeedGen64
func PCG32Init() PCG32 {
	return PCG32New(SeedGen64())
}

// Generates a random uint32 number
func (s *PCG32) Next() u32 {
	state := u64(*s)
	*s = PCG32(state*0xf13283ad + 0x9e3779b97f4a7c15)
	return RotateU32(u32((state^state>>18)>>27), u8(state>>59))
}

// Generates number in range [0,n)
// The larger the n the larger the bias (in general)
// Usualy in practice it is insignificant
func (s *PCG32) Range(n si) si {
	return si(s.Next() % u32(n))
}

// Generates number in range [0,1]
func (s *PCG32) Normal32() f32 {
	return f32(s.Next()) / f32(1<<32-1)
}

// Generates number in range [0,1]. Has 32bit resolution
func (s *PCG32) Normal64() f64 {
	return f64(s.Next()) / f64(1<<32-1)
}
