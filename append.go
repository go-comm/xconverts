package xconverts

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func AppendInt8(dst []byte, n int8) []byte {
	return AppendInt64(dst, int64(n))
}

func AppendUint8(dst []byte, n uint8) []byte {
	return AppendUint64(dst, uint64(n))
}

func AppendInt16(dst []byte, n int8) []byte {
	return AppendInt64(dst, int64(n))
}

func AppendUint16(dst []byte, n int8) []byte {
	return AppendUint64(dst, uint64(n))
}

func AppendInt(dst []byte, n int) []byte {
	return AppendInt64(dst, int64(n))
}

func AppendUint(dst []byte, n uint) []byte {
	return AppendUint64(dst, uint64(n))
}

func AppendInt64(dst []byte, n int64) []byte {
	return strconv.AppendInt(dst, n, 10)
}

func AppendUint64(dst []byte, n uint64) []byte {
	return strconv.AppendUint(dst, n, 10)
}

func AppendFloat32(dst []byte, n float32, fmt byte, prec int) []byte {
	return strconv.AppendFloat(dst, float64(n), fmt, prec, 32)
}

func AppendFloat64(dst []byte, n float64, fmt byte, prec int) []byte {
	return strconv.AppendFloat(dst, n, fmt, prec, 64)
}

func AppendStr(dst []byte, s ...string) []byte {
	switch len(s) {
	case 0:
		return dst
	case 1:
		return append(dst, s[0]...)
	default:
	}
	dst = grow(dst, sizeOfStrSlice(s))
	for _, v := range s {
		dst = append(dst, v...)
	}
	return dst
}

func AppendByte(dst []byte, s ...byte) []byte {
	return append(dst, s...)
}

func AppendBytes(dst []byte, s ...[]byte) []byte {
	switch len(s) {
	case 0:
		return dst
	case 1:
		return append(dst, s[0]...)
	default:
	}
	dst = grow(dst, sizeOfBytesSlice(s))
	for _, v := range s {
		dst = append(dst, v...)
	}
	return dst
}

func AppendTime(dst []byte, t time.Time, layout string) []byte {
	return t.AppendFormat(dst, layout)
}

func AppendAny(dst []byte, i interface{}) []byte {
	if b, ok := IsInt(i); ok {
		return AppendInt64(dst, b)
	}
	if b, ok := IsUint(i); ok {
		return AppendUint64(dst, b)
	}
	switch b := i.(type) {
	case float32:
		return AppendFloat32(dst, b, 'f', 10)
	case float64:
		return AppendFloat64(dst, b, 'f', 10)
	case string:
		return AppendStr(dst, b)
	case *string:
		return AppendStr(dst, *b)
	case []byte:
		return AppendBytes(dst, b)
	case *[]byte:
		return AppendBytes(dst, *b)
	case sql.RawBytes:
		return AppendBytes(dst, b)
	case *sql.RawBytes:
		return AppendBytes(dst, *b)
	}
	return AppendStr(dst, fmt.Sprintf("%v", i))
}
