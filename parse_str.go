package xconverts

import (
	"strconv"
	"time"
)

func ParseInt8(s string) (int8, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	return int8(v), err
}

func ParseUint8(s string) (uint8, error) {
	v, err := strconv.ParseUint(s, 10, 0)
	return uint8(v), err
}

func ParseInt16(s string) (int16, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	return int16(v), err
}

func ParseUint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 10, 0)
	return uint16(v), err
}

func ParseInt(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	return int(v), err
}

func ParseUint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 0)
	return uint(v), err
}

func ParseInt32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	return int32(v), err
}

func ParseUint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 10, 0)
	return uint32(v), err
}

func ParseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 0)
}

func ParseUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 0)
}

func ParseFloat32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 32)
	return float32(v), err
}

func ParseFloat64(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	return float64(v), err
}

func ParseTime(s string, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}
