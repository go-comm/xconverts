package xconverts

import (
	"strconv"
	"time"
)

func ParseBytesToInt8(s []byte) (int8, error) {
	v, err := strconv.ParseInt(string(s), 10, 0)
	return int8(v), err
}

func ParseBytesToUint8(s []byte) (uint8, error) {
	v, err := strconv.ParseUint(string(s), 10, 0)
	return uint8(v), err
}

func ParseBytesToInt16(s []byte) (int16, error) {
	v, err := strconv.ParseInt(string(s), 10, 0)
	return int16(v), err
}

func ParseBytesToUint16(s []byte) (uint16, error) {
	v, err := strconv.ParseUint(string(s), 10, 0)
	return uint16(v), err
}

func ParseBytesToInt(s []byte) (int, error) {
	v, err := strconv.ParseInt(string(s), 10, 0)
	return int(v), err
}

func ParseBytesToUint(s []byte) (uint, error) {
	v, err := strconv.ParseUint(string(s), 10, 0)
	return uint(v), err
}

func ParseBytesToInt32(s []byte) (int32, error) {
	v, err := strconv.ParseInt(string(s), 10, 0)
	return int32(v), err
}

func ParseBytesToUint32(s []byte) (uint32, error) {
	v, err := strconv.ParseUint(string(s), 10, 0)
	return uint32(v), err
}

func ParseBytesToInt64(s []byte) (int64, error) {
	return strconv.ParseInt(string(s), 10, 0)
}

func ParseBytesToUint64(s []byte) (uint64, error) {
	return strconv.ParseUint(string(s), 10, 0)
}

func ParseBytesToFloat32(s []byte) (float32, error) {
	v, err := strconv.ParseFloat(string(s), 32)
	return float32(v), err
}

func ParseBytesToFloat64(s []byte) (float64, error) {
	v, err := strconv.ParseFloat(string(s), 64)
	return float64(v), err
}

func ParseBytesToTime(s []byte, layout string) (time.Time, error) {
	return time.Parse(layout, string(s))
}
