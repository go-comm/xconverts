package xconverts

import "testing"

func TestGrow(t *testing.T) {
	var p []byte

	p = []byte("h")

	t.Log(len(p), cap(p))
	p = grow(p, 2)
	t.Log(len(p), cap(p))

	p = append(p, []byte("jj")...)

	t.Log(len(p), cap(p))
}
