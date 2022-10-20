package gomisc

import "time"

type PCG32Fast uint64

func PCG32FastNew(seed uint64) PCG32Fast {
	return PCG32Fast(seed*2 + 1)
}

// Initializes with time.UnixNano
func PCG32FastInit() PCG32Fast {
	return PCG32Fast(uint64(time.Now().UnixNano())*2 + 1)
}

// Generates a uint32 number
func (s PCG32Fast) Raw() uint32 {
	s *= 4046619565
	return uint32((s ^ s>>22) >> (22 + s>>61))
}

// Generates number in range [0,n)
func (s PCG32Fast) Range(n uint32) uint32 {
	return s.Raw() % n
}

// Generates number in range [0,1)
func (s PCG32Fast) Exclusive01F32() float32 {
	return float32(s.Raw()) / float32(1<<32)
}

// Generates number in range [0,1). Has 32bit resolution
func (s PCG32Fast) Exclusive01F64() float64 {
	return float64(s.Raw()) / float64(1<<32)
}
