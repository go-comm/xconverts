package xconverts

func JoinStr(dst []byte, s []string, sep string) []byte {
	switch len(s) {
	case 0:
		return dst
	case 1:
		return append(dst, s[0]...)
	default:
	}
	dst = grow(dst, sizeOfStrSlice(s))
	dst = append(dst, s[0]...)
	for i := 1; i < len(s); i++ {
		dst = append(dst, sep...)
		dst = append(dst, s[i]...)
	}
	return dst
}

func JoinBytes(dst []byte, s [][]byte, sep []byte) []byte {
	switch len(s) {
	case 0:
		return dst
	case 1:
		return append(dst, s[0]...)
	default:
	}
	dst = grow(dst, sizeOfBytesSlice(s))
	dst = append(dst, s[0]...)
	for i := 1; i < len(s); i++ {
		dst = append(dst, sep...)
		dst = append(dst, s[i]...)
	}
	return dst
}

type int8Slice []int8

func (x int8Slice) Len() int           { return len(x) }
func (x int8Slice) Index(i int) []byte { return FormatInt8ToBytes(x[i]) }

func JoinInt8s(dst []byte, x []int8, sep []byte) []byte {
	return Join(dst, (int8Slice)(x), sep)
}

type uint8Slice []uint8

func (x uint8Slice) Len() int           { return len(x) }
func (x uint8Slice) Index(i int) []byte { return FormatUint8ToBytes(x[i]) }

func JoinUint8s(dst []byte, x []uint8, sep []byte) []byte {
	return Join(dst, (uint8Slice)(x), sep)
}

type int16Slice []int16

func (x int16Slice) Len() int           { return len(x) }
func (x int16Slice) Index(i int) []byte { return FormatInt16ToBytes(x[i]) }

func JoinInt16s(dst []byte, x []int16, sep []byte) []byte {
	return Join(dst, (int16Slice)(x), sep)
}

type uint16Slice []uint16

func (x uint16Slice) Len() int           { return len(x) }
func (x uint16Slice) Index(i int) []byte { return FormatUint16ToBytes(x[i]) }

func JoinUint16s(dst []byte, x []uint16, sep []byte) []byte {
	return Join(dst, (uint16Slice)(x), sep)
}

type intSlice []int

func (x intSlice) Len() int           { return len(x) }
func (x intSlice) Index(i int) []byte { return FormatIntToBytes(x[i]) }

func JoinInts(dst []byte, x []int, sep []byte) []byte {
	return Join(dst, (intSlice)(x), sep)
}

type uintSlice []uint

func (x uintSlice) Len() int           { return len(x) }
func (x uintSlice) Index(i int) []byte { return FormatUintToBytes(x[i]) }

func JoinUints(dst []byte, x []uint, sep []byte) []byte {
	return Join(dst, (uintSlice)(x), sep)
}

type int32Slice []int32

func (x int32Slice) Len() int           { return len(x) }
func (x int32Slice) Index(i int) []byte { return FormatInt32ToBytes(x[i]) }

func JoinInt32s(dst []byte, x []int32, sep []byte) []byte {
	return Join(dst, (int32Slice)(x), sep)
}

type uint32Slice []uint32

func (x uint32Slice) Len() int           { return len(x) }
func (x uint32Slice) Index(i int) []byte { return FormatUint32ToBytes(x[i]) }

func JoinUint32s(dst []byte, x []uint32, sep []byte) []byte {
	return Join(dst, (uint32Slice)(x), sep)
}

type int64Slice []int64

func (x int64Slice) Len() int           { return len(x) }
func (x int64Slice) Index(i int) []byte { return FormatInt64ToBytes(x[i]) }

func JoinInt64s(dst []byte, x []int64, sep []byte) []byte {
	return Join(dst, (int64Slice)(x), sep)
}

type uint64Slice []uint64

func (x uint64Slice) Len() int           { return len(x) }
func (x uint64Slice) Index(i int) []byte { return FormatUint64ToBytes(x[i]) }

func JoinUint64s(dst []byte, x []uint64, sep []byte) []byte {
	return Join(dst, (uint64Slice)(x), sep)
}

type float32Slice []float32

func (x float32Slice) Len() int           { return len(x) }
func (x float32Slice) Index(i int) []byte { return formatFloat32ToBytes(x[i]) }

func JoinFloat32s(dst []byte, x []float32, sep []byte) []byte {
	return Join(dst, (float32Slice)(x), sep)
}

type float64Slice []float64

func (x float64Slice) Len() int           { return len(x) }
func (x float64Slice) Index(i int) []byte { return formatFloat64ToBytes(x[i]) }

func JoinFloat64s(dst []byte, x []float64, sep []byte) []byte {
	return Join(dst, (float64Slice)(x), sep)
}

type JoinInterface interface {
	Index(i int) []byte
	Len() int
}

func Join(dst []byte, s JoinInterface, sep []byte) []byte {
	n := s.Len()
	switch n {
	case 0:
		return dst
	case 1:
		return append(dst, s.Index(0)...)
	default:
	}
	dst = append(dst, s.Index(0)...)
	for i := 1; i < n; i++ {
		dst = append(dst, sep...)
		dst = append(dst, s.Index(i)...)
	}
	return dst
}
