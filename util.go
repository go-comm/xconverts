package xconverts

import "errors"

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
