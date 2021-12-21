package xconverts

import (
	"errors"
)

var ErrConvertFail = errors.New("xconverts: convert fail")

func grow(p []byte, n int) []byte {
	b := p
	if n > cap(p)-len(p) {
		b = make([]byte, len(p)+n)
		copy(b, p)
		b = b[:len(p)]
	}
	return b
}

func sizeOfStrSlice(p []string) int {
	n := 0
	for _, v := range p {
		n += len(v)
	}
	return n
}

func sizeOfBytesSlice(p [][]byte) int {
	n := 0
	for _, v := range p {
		n += len(v)
	}
	return n
}

func IsInt(i interface{}) (int64, bool) {
	switch b := i.(type) {
	case int8:
		return int64(b), true
	case int16:
		return int64(b), true
	case int:
		return int64(b), true
	case int32:
		return int64(b), true
	case int64:
		return b, true
	}
	return 0, false
}

func IsUint(i interface{}) (uint64, bool) {
	switch b := i.(type) {
	case uint8:
		return uint64(b), true
	case uint16:
		return uint64(b), true
	case uint:
		return uint64(b), true
	case uint32:
		return uint64(b), true
	case uint64:
		return b, true
	}
	return 0, false
}

func IsFloat32(i interface{}) (float32, bool) {
	switch b := i.(type) {
	case float32:
		return b, true
	}
	return 0, false
}

func IsFloat64(i interface{}) (float64, bool) {
	switch b := i.(type) {
	case float64:
		return b, true
	}
	return 0, false
}
