package xconverts

import (
	"strconv"
	"time"
)

func FormatInt8ToBytes(n int8) []byte {
	return strconv.AppendInt([]byte(nil), int64(n), 10)
}

func FormatUint8ToBytes(n uint8) []byte {
	return strconv.AppendUint([]byte(nil), uint64(n), 10)
}

func FormatInt16ToBytes(n int16) []byte {
	return strconv.AppendInt([]byte(nil), int64(n), 10)
}

func FormatUint16ToBytes(n uint16) []byte {
	return strconv.AppendUint([]byte(nil), uint64(n), 10)
}

func FormatIntToBytes(n int) []byte {
	return strconv.AppendInt([]byte(nil), int64(n), 10)
}

func FormatUintToBytes(n uint) []byte {
	return strconv.AppendUint([]byte(nil), uint64(n), 10)
}

func FormatInt32ToBytes(n int32) []byte {
	return strconv.AppendInt([]byte(nil), int64(n), 10)
}

func FormatUint32ToBytes(n uint32) []byte {
	return strconv.AppendUint([]byte(nil), uint64(n), 10)
}

func FormatInt64ToBytes(n int64) []byte {
	return strconv.AppendInt([]byte(nil), n, 10)
}

func FormatUint64ToBytes(n uint64) []byte {
	return strconv.AppendUint([]byte(nil), n, 10)
}

func FormatFloat32ToBytes(n float32, fmt byte, prec int) []byte {
	return strconv.AppendFloat([]byte(nil), float64(n), fmt, prec, 32)
}

func formatFloat32ToBytes(n float32) []byte {
	return FormatFloat32ToBytes(n, 'f', 10)
}

func FormatFloat64ToBytes(n float64, fmt byte, prec int) []byte {
	return strconv.AppendFloat([]byte(nil), n, fmt, prec, 64)
}

func formatFloat64ToBytes(n float64) []byte {
	return FormatFloat64ToBytes(n, 'f', 10)
}

func FormatTimeToBytes(t time.Time, layout string) []byte {
	return t.AppendFormat([]byte(nil), layout)
}

func formatTimeToBytes(t time.Time) []byte {
	return FormatTimeToBytes(t, time.RFC3339)
}
