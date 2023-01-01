package gomisc

import "time"

type Seconds int64
type Millis int64
type Micros int64
type Nanos int64

func SecondsGet() Seconds {
	return Seconds(time.Now().Unix())
}
func MillisGet() Millis {
	return Millis(time.Now().UnixMilli())
}
func MicrosGet() Micros {
	return Micros(time.Now().UnixMicro())
}
func NanosGet() Nanos {
	return Nanos(time.Now().UnixNano())
}
